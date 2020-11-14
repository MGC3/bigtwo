package app

import (
    "encoding/json"
    "log"
//    "sync"
    "github.com/gorilla/websocket"
)

const (
    maxNumPlayersInRoom = 4
)

type roomId string

// Represents a single room.
// A room can have 0 to maxNumPlayersinRoom players.
type room struct {
    id roomId
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
            newRoomId := w.CreateNewRoom()
            log.Printf("created new room %s\n", newRoomId)
            p := w.ConnectedPlayersNotInRoom[msg.PlayerId]

            // TODO make this nicer
            msgData, _ := json.Marshal(RoomCreatedData{
                RoomId: newRoomId,
            });
            msg := Message {{}
                PlayerId: msg.PlayerId,
                Type: "room_created",
                Data: json.RawMessage(msgData),
            };
            p.toPlayer <- msg
        case "join_room":
            // Assumes room exists, sends error if no room exists
            var nested JoinRoomData
            err := json.Unmarshal(msg.Data, &nested)
            if err != nil {
                log.Printf("Could not unmarshal nested packet %v", err)
                continue
            }
            log.Printf("got join room %s from player %s\n", nested.Room, nested.Name)
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

func (w *WaitingArea) CreateNewRoom() roomId {
    r := room {
        id: "ABCD",
        players: []player{},
        receive: make(chan Message),
    }

    w.WaitingForPlayers[r.id] = r
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

func NewWaitingArea() WaitingArea {
    return WaitingArea {
        WaitingForPlayers: make(map[roomId]room),
        InGame: make(map[roomId]room),
        ConnectedPlayersNotInRoom: make(map[playerId]player),
        nextId: 0,
        receive: make(chan Message),
    }
}
