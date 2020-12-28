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
	id roomId

	// Players are initialized to nil
	// It's assumed that non-nil players are at the start of the array
	// And that all uninitialized (nil) players come after the initialized
	// players
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
	disconnectedClientId := r.clientIdFromPlayerId(receive.Player.id)

	if disconnectedClientId == -1 {
		log.Printf("room handleDisconnect error - no player found")
		return
	}

	// move all players to fill in disconnected players
	r.players[disconnectedClientId] = nil
	for i := disconnectedClientId + 1; i < maxNumPlayersInRoom; i++ {
		r.players[i-1] = r.players[i]
	}

	r.pushRoomStateToPlayers()
}

func (r *room) handleJoinRoom(receive Message) {
	log.Printf("Got join room from player %v\n", receive.Player)
	playerFound := false
	for i, player := range r.players {
		if player == nil {
			r.players[i] = receive.Player
			send, err := NewMessage(receive.Player, "room_joined", EmptyData{})
			if err != nil {
				break
			}
			receive.Player.toPlayer <- send
			log.Printf("Room state %v\n", r)
			playerFound = true
			break
		}
	}

	if !playerFound {
		log.Printf("room handleJoinRoom error - no player found")
		return
	}

	// forward room state changed to all clients
	r.pushRoomStateToPlayers()
}

func (r *room) handleRequestRoomState(receive Message) {
	log.Printf("Got request room state from player %v\n", receive.Player)
	data := r.roomStateData()
	data.ClientId = r.clientIdFromPlayerId(receive.Player.id)
	msg, err := NewMessage(receive.Player, "room_state", data)
	if err != nil {
		log.Printf("handleRequestRoomState err %v\n", err)
		return
	}
	receive.Player.toPlayer <- msg
}

func (r *room) pushRoomStateToPlayers() {
	data := r.roomStateData()
	for clientId, player := range r.players {
		if player == nil {
			break
		}
		data.ClientId = clientId
		msg, err := NewMessage(player, "room_state", data)
		if err != nil {
			log.Printf("pushRoomStateToPlayers err %v\n", err)
			return
		}
		player.toPlayer <- msg
	}
}

func (r *room) clientIdFromPlayerId(id playerId) int {
	clientId := -1
	for i, player := range r.players {
		if player == nil {
			break
		}

		if player.id == id {
			clientId = i
			break
		}
	}

	if clientId == -1 {
		log.Printf("Error: no player with id %d found in room %s\n", id, r.id)
	}

	return clientId
}

func (r *room) roomStateData() RoomStateData {
	ret := RoomStateData{PlayerNames: []string{}}
	for _, player := range r.players {
		if player == nil {
			break
		}

		ret.PlayerNames = append(ret.PlayerNames, player.displayName)
	}

	return ret
}

func (r *room) numPlayers() int {
	n := 0
	for _, player := range r.players {
		if player != nil {
			n += 1
		}
	}

	return n
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
	r.messageHandlers["request_room_state"] = r.handleRequestRoomState

	go r.serve()
	return &r
}
