package app

import (
	//	"encoding/json"
	"log"
	//    "sync"
	//	"github.com/gorilla/websocket"
)

const (
	maxNumPlayersInRoom = 4
)

type roomId string

// Represents a single room.
// A room can have 0 to maxNumPlayersinRoom players.
type room struct {
	id              roomId
	players         [maxNumPlayersInRoom]*player
	receive         chan Message
	messageHandlers map[string]func(Message)

	// TODO waiting area channel to signal when room is done with the game?
}

func (r *room) serve() {
	// TODO this is basically the same code as WaitingArea.Serve()
	// generalize code?
	log.Printf("room %s serving...\n", r.id)

	for {
		receive := <-r.receive
		handler, ok := r.messageHandlers[receive.Type]

		if !ok {
			log.Printf("Unhandled message type %s\n", receive.Type)
			continue
		}

		handler(receive)
	}
}

func (r *room) handleDisconnect(receive Message) {
	// TODO delete (or invalidate?) player from array
	for i, player := range r.players {
		log.Printf("its a player %v\n", player)
		if player == nil {
			continue
			if player == nil {
				continue
			}
		}
		if player.id == receive.Player.id {
			r.players[i] = nil
			return
		}
	}

	log.Printf("room handleDisconnect error - no player found")
}

func (r *room) handleJoinRoom(receive Message) {
	log.Printf("Got join room from player %v\n", receive.Player)
	for i, player := range r.players {
		if player == nil {
			r.players[i] = receive.Player
			send, err := NewMessage(receive.Player, "room_joined", EmptyData{})
			if err != nil {
				break
			}
			receive.Player.toPlayer <- send
			log.Printf("Room state %v\n", r)
			return
		}
	}

	log.Printf("room handleJoinRoom error - no player found")
}

func newRoom(id roomId) *room {
	r := room{
		// TODO generate real random-ish string
		id:              "ABCD",
		players:         [maxNumPlayersInRoom]*player{nil, nil, nil, nil},
		receive:         make(chan Message),
		messageHandlers: make(map[string]func(Message)),
	}

	r.messageHandlers["disconnect"] = r.handleDisconnect
	r.messageHandlers["join_room"] = r.handleJoinRoom

	go r.serve()
	return &r
}
