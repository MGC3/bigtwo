package app

import (
    "fmt"
    "sync"
    "github.com/gorilla/websocket"
)

const (
    maxNumPlayersInRoom = 4
)
// Represents a single room.
// A room can have 0 to maxNumPlayersinRoom players.
// The room's state (waiting, in progress, finished) is used
// to determine if someone should be allowed to join a room.
type roomId string

type room struct {
    id roomId
    players []player
}

type WaitingArea struct {
    WaitingForPlayers map[roomId]room
    InGame map[roomId]room
    ConnectedPlayersNotInRoom map[playerId]player 
    nextId int
}

// pass in waitgroup by channel to goroutines that are listening to the connection
// so that I can swap send channels 
func (w *WaitingArea) AddNewConnectedPlayer(conn *websocket.Conn) {

}

// Creates a new room witha unique Id
func (f *FrontDesk) CreateRoom() (int, error) {
    currentRoomId := f.nextId
    f.nextId += 1
    f.activeRooms[currentRoomId] = &room{players: []player{}, state: roomWaitingForPlayers}
    return currentRoomId, nil
}

func (f *FrontDesk) JoinRoom(roomId int, conn *websocket.Conn) error {
    if _, ok := f.activeRooms[roomId]; !ok {
        return fmt.Errorf("No active room with ID %d", roomId)
    }

    if len(f.activeRooms[roomId].players) >= maxNumPlayersInRoom {
        return fmt.Errorf("Max number of players in room %d", roomId)
    }

    newPlayer := player {
        id: 0,
        displayName: "test",
        conn: conn,
    }
    f.activeRooms[roomId].players = append(f.activeRooms[roomId].players, newPlayer)
    return nil
}

func NewWaitingArea() WaitingArea {
    return WaitingArea {
        WaitingForPlayers: make map[]
    }
}
