package app

import (
	"encoding/json"
	"log"
	//    "sync"
	"github.com/gorilla/websocket"
)

type WaitingArea struct {
	WaitingForPlayers map[roomId]*room
	InGame            map[roomId]*room
	// TODO lock needed?
	ConnectedPlayersNotInRoom map[playerId]player
	nextId                    playerId
	receive                   chan Message

	messageHandlers map[string]func(*WaitingArea, Message)
}

// Serves connected players not in a room.
// Handles displayName assignment and players creating/joining rooms
func (w *WaitingArea) Serve() {
	log.Println("WaitingArea serving...")
	for {
		receive := <-w.receive
		handler, ok := w.messageHandlers[receive.Type]

		if !ok {
			log.Printf("Unhandled message type %s\n", receive.Type)
			continue
		}

		handler(w, receive)
	}
}

func (w *WaitingArea) CreateNewRoom() roomId {
	r := newRoom("ABCD")
	w.WaitingForPlayers[r.id] = newRoom("ABCD")
	return r.id
}

func (w *WaitingArea) AddNewConnectedPlayer(conn *websocket.Conn) {
	if _, ok := w.ConnectedPlayersNotInRoom[w.nextId]; ok {
		log.Fatal("Error - player ids are not unique")
	}

	p := newPlayer(w.nextId, conn, w.receive)
	log.Printf("Got new player id %v\n", p)
	w.nextId += 1
	w.ConnectedPlayersNotInRoom[p.id] = p
}

func handleCreateRoom(w *WaitingArea, receive Message) {
	newRoomId := w.CreateNewRoom()
	log.Printf("created new room %s\n", newRoomId)
	p, ok := w.ConnectedPlayersNotInRoom[receive.PlayerId]

	if !ok {
		log.Printf("Could not find player %s\n", receive.PlayerId)
		return
	}

	send, err := NewMessage(receive.PlayerId, "room_created", RoomCreatedData{RoomId: newRoomId})
	if err != nil {
		log.Printf("Error creating message %v\n", err)
		return
	}
	p.toPlayer <- send
}

func handleJoinRoom(w *WaitingArea, receive Message) {
	// Assumes room exists, sends error if no room exists
	var nested JoinRoomData
	err := json.Unmarshal(receive.Data, &nested)
	if err != nil {
		log.Printf("Could not unmarshal nested packet %v", err)
		return
	}
	log.Printf("got join room %s from player %s\n", nested.RoomId, nested.Name)
	player, ok := w.ConnectedPlayersNotInRoom[receive.PlayerId]

	if !ok {
		// TODO send error messages or something
		log.Printf("failed to join room because player %d not found\n", receive.PlayerId)
		return
	}

	room, ok := w.WaitingForPlayers[nested.RoomId]

	if !ok {
		log.Printf("failed to join room because room %d not found", nested.RoomId)
		return
	}

	delete(w.ConnectedPlayersNotInRoom, receive.PlayerId)
	room.players = append(room.players, player)

	send, err := NewMessage(receive.PlayerId, "room_joined", EmptyData{})

	// TODO swap out player channels
	room = w.WaitingForPlayers[nested.RoomId]
	log.Printf("Room state: %v\n", room)
	player.toPlayer <- send
}

func handleDisconnect(w *WaitingArea, receive Message) {
	log.Printf("Player %d disconnected\n", receive.PlayerId)

	// Forward disconnect message to send thread
	w.ConnectedPlayersNotInRoom[receive.PlayerId].toPlayer <- receive
	delete(w.ConnectedPlayersNotInRoom, receive.PlayerId)

	log.Printf("connected players %v", w.ConnectedPlayersNotInRoom)
}

func NewWaitingArea() WaitingArea {
	w := WaitingArea{
		WaitingForPlayers:         make(map[roomId]*room),
		InGame:                    make(map[roomId]*room),
		ConnectedPlayersNotInRoom: make(map[playerId]player),
		nextId:                    0,
		receive:                   make(chan Message),
		messageHandlers:           make(map[string]func(*WaitingArea, Message)),
	}

	w.messageHandlers["create_room"] = handleCreateRoom
	w.messageHandlers["join_room"] = handleJoinRoom
	w.messageHandlers["disconnect"] = handleDisconnect

	return w
}
