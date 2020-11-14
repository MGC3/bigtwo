package app

import (
    "log"
//    "sync"
    "github.com/gorilla/websocket"
)

const (
    maxNumPlayersInRoom = 4
)

type roomState int
const (
    waitingForPlayers roomState = iota
    inGame
    gameOver
)

type roomId string

// Represents a single room.
// A room can have 0 to maxNumPlayersinRoom players.
type room struct {
    id roomId
    state roomState
    players []player
    receive chan Message 
}

type WaitingArea struct {
    WaitingForPlayers map[roomId]room
    InGame map[roomId]room
    // TODO lock needed?
    ConnectedPlayersNotInRoom map[playerId]player 
    nextId playerId 
    receive chan Message 
}

// Serves connected players not in a room.
// Handles displayName assignment and players creating/joining rooms
func (w *WaitingArea) Serve() {
    log.Println("WaitingArea serving...")
    for {
        msg := <-w.receive
        switch msg.Type {
        case "send_display_name":
            log.Println("got display name")
        case "create_room":
            log.Println("got create room")
        case "join_room":
            // Assumes room exists, sends error if no room exists
            log.Println("got join room")
        case "disconnect":
            log.Printf("Player %d disconnected\n", msg.PlayerId)
            p := w.ConnectedPlayersNotInRoom[msg.PlayerId]
            p.toPlayer <- msg
            delete(w.ConnectedPlayersNotInRoom, p.id)
            log.Printf("connected players %v", w.ConnectedPlayersNotInRoom)
        default:
            log.Printf("error -- unhandled message type %s\n", msg.Type)
            return
        }
    }
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

func NewWaitingArea() WaitingArea {
    return WaitingArea {
        WaitingForPlayers: make(map[roomId]room),
        InGame: make(map[roomId]room),
        ConnectedPlayersNotInRoom: make(map[playerId]player),
        nextId: 0,
        receive: make(chan Message),
    }
}
