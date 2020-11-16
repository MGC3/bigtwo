/*

Used https://github.com/gorilla/websocket/blob/master/examples/chat/client.go for reference.

*/

package app

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type playerId int

type player struct {
	id          playerId
	displayName string
	conn        *websocket.Conn
	toPlayer    chan Message

	toServerLock sync.Mutex
	toServer     chan Message
}

// Factory func for creating a new player.
// Players are not initialized with a display name -- defaults to empty string.
func newPlayer(id playerId, conn *websocket.Conn, toServer chan Message) player {
	p := player{
		id:           id,
		displayName:  "",
		conn:         conn,
		toPlayer:     make(chan Message),
		toServerLock: sync.Mutex{},
		toServer:     toServer,
	}

	go p.receiveThread()
	go p.sendThread()

	return p
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
			p.toServer <- Message{PlayerId: p.id, Type: "disconnect", Data: []byte{}}
			return
		}
		log.Printf("ReceiveThread got %v from p %d\n", bytes, p.id)
		// TODO format message to have ID of sending player

		if err := json.Unmarshal(bytes, &msg); err != nil {
			log.Printf("receiveThread failed to unmarshal bc %v", err)
			continue
		}

		msg.PlayerId = p.id
		p.toServerLock.Lock()
		// TODO unmarshal msg into Message type
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
		} else {
			log.Printf("sendThread wrote message %v bytes\n", msg)
		}
	}
}
