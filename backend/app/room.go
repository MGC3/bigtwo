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
	players         []player
	receive         chan Message
	messageHandlers map[string]func(*room, Message)
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

		handler(r, receive)
	}
}

func roomHandleDisconnect(r *room, receive Message) {
	// TODO delete (or invalidate?) player from array
	for _, player := range r.players {
		log.Printf("its a player %v\n", player)
	}
}

func newRoom(id roomId) *room {
	r := room{
		// TODO generate real random-ish string
		id:              "ABCD",
		players:         []player{},
		receive:         make(chan Message),
		messageHandlers: make(map[string]func(*room, Message)),
	}

	r.messageHandlers["disconnect"] = roomHandleDisconnect
	return &r
}
