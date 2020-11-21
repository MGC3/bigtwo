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
	ConnectedPlayersNotInRoom map[playerId]*player
	nextId                    playerId
	receive                   chan Message

	messageHandlers map[string]func(Message)
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

		handler(receive)
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

func (w *WaitingArea) handleCreateRoom(receive Message) {
	newRoomId := w.CreateNewRoom()
	log.Printf("created new room %s\n", newRoomId)
	p, ok := w.ConnectedPlayersNotInRoom[receive.Player.id]

	if !ok {
		log.Printf("Could not find player %v\n", receive.Player)
		return
	}

	send, err := NewMessage(receive.Player, "room_created", RoomCreatedData{RoomId: newRoomId})
	if err != nil {
		log.Printf("Error creating message %v\n", err)
		return
	}
	p.toPlayer <- send
}

func (w *WaitingArea) handleJoinRoom(receive Message) {
	// Assumes room exists, sends error if no room exists
	var nested JoinRoomData
	err := json.Unmarshal(receive.Data, &nested)
	if err != nil {
		log.Printf("Could not unmarshal nested packet %v", err)
		return
	}
	log.Printf("got join room %s from player %s\n", nested.RoomId, nested.Name)
	if _, ok := w.ConnectedPlayersNotInRoom[receive.Player.id]; !ok {
		// TODO send error messages or something
		log.Printf("failed to join room because player %d not found\n", receive.Player.id)
		return
	}

	room, ok := w.WaitingForPlayers[nested.RoomId]

	if !ok {
		log.Printf("failed to join room because room %d not found", nested.RoomId)
		return
	}

	// After deleting, the player is passed off to the thread running room.serve
	delete(w.ConnectedPlayersNotInRoom, receive.Player.id)

	receive.Player.swapToServerChannel(room.receive)

	// Forward the message to the room itself
	room.receive <- receive

	log.Printf("Room state: %v\n", room)
}

func (w *WaitingArea) handleDisconnect(receive Message) {
	log.Printf("Player %d disconnected\n", receive.Player.id)

	// Forward disconnect message to send thread
	if _, ok := w.ConnectedPlayersNotInRoom[receive.Player.id]; !ok {
		log.Printf("Error: no player with ID %d found in connected player map\n", receive.Player.id)
		return
	}
	receive.Player.toPlayer <- receive
	delete(w.ConnectedPlayersNotInRoom, receive.Player.id)

	log.Printf("connected players %v", w.ConnectedPlayersNotInRoom)
}

func NewWaitingArea() WaitingArea {
	w := WaitingArea{
		WaitingForPlayers:         make(map[roomId]*room),
		InGame:                    make(map[roomId]*room),
		ConnectedPlayersNotInRoom: make(map[playerId]*player),
		nextId:                    0,
		receive:                   make(chan Message),
		messageHandlers:           make(map[string]func(Message)),
	}

	w.messageHandlers["create_room"] = w.handleCreateRoom
	w.messageHandlers["join_room"] = w.handleJoinRoom
	w.messageHandlers["disconnect"] = w.handleDisconnect

	return w
}
