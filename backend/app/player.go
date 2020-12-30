/*

Used https://github.com/gorilla/websocket/blob/master/examples/chat/client.go for reference.

*/

package app

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"

	"github.com/MGC3/bigtwo/backend/app/game"
)

type playerId int

const invalidPlayerId = -1

type player struct {
	id          playerId
	displayName string
	conn        *websocket.Conn
	toPlayer    chan Message

	toServerLock sync.Mutex
	toServer     chan Message

	// TODO is there a better place to put this?
	// This won't get used until a game is started
	currentHand []game.Card
}

// TODO hacky way of using channels in main?
func UninitializedPlayer(conn *websocket.Conn) *player {
	p := player{
		id:           invalidPlayerId,
		displayName:  "",
		conn:         conn,
		toPlayer:     make(chan Message),
		toServerLock: sync.Mutex{},
		toServer:     nil,
	}

	return &p
}

func (p *player) initialize(id playerId, toServer chan Message) error {
	if p.id != invalidPlayerId || p.toServer != nil {
		return errors.New("Can't initialize player that already has an ID")
	}

	p.id = id
	p.toServer = toServer
	go p.receiveThread()
	go p.sendThread()
	return nil
}

// Thread for reading message from the websocket connection and sending them to the room
// Just loops forever and forwards messages from conn to room.receiveChannel
func (p *player) receiveThread() {
	// TODO set connection parameters
	log.Printf("receiveThread running for player %d", p.id)
	var msg Message
	for {
		_, bytes, err := p.conn.ReadMessage()
		if err != nil {
			/*
			   if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			       log.Printf("receiveThread err %v", err)
			   }
			*/

			// TODO signal to server that this user has disconnected
			log.Printf("receiveThread stopping for player %d\n", p.id)
			p.toServer <- Message{Player: p, Type: "disconnect", Data: []byte{}}
			return
		}

		if err := json.Unmarshal(bytes, &msg); err != nil {
			log.Printf("receiveThread failed to unmarshal bc %v", err)
			continue
		}

		// TODO is this weird to pass around a pointer to the receiving player like this?
		msg.Player = p

		// TODO could definitely have race conditions if e.g., connection receives a message meant for WaitingArea
		// but swaps channels to room handler after message is received but before sending
		p.toServerLock.Lock()
		p.toServer <- msg
		p.toServerLock.Unlock()
	}
}

// Each player is going to be talking to different parts of the backend at different times
// ie, player messages need to be handled differently when the player is not in a room vs
// in a room waiting for players vs in game
// This function lets the backend safely swap out the receive channel
func (p *player) swapToServerChannel(newChannel chan Message) {
	p.toServerLock.Lock()
	p.toServer = newChannel
	p.toServerLock.Unlock()
}

// Thread for sending message from room to websocket connection
func (p *player) sendThread() {
	log.Printf("sendThread running for player %d\n", p.id)
	for {
		msg, ok := <-p.toPlayer
		// TODO signal that connection is over by closing channel? Better way to do this?
		if !ok {
			p.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		if msg.Type == "disconnect" {
			log.Printf("SendThread for player %d got dc message. Exiting", p.id)
			return
		}

		err := p.conn.WriteJSON(msg)
		if err != nil {
			log.Printf("sendThread write failed %v\n", err)
			// Don't let thread die here -- need to wait for disconnect message
			// to avoid deadlock maybe not 100% sure
		}
	}
}
