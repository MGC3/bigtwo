package app

import (
    //"log"

    "github.com/MGC3/bigtwo/backend/app/game"
)

// This type defines the outer message type, which contains
// a string that identifies what kind of message is stored
// as bytes in data
type Message struct {
    // -1 for server?
    PlayerId playerId 
    Type string `json:"type"`

    // or maybe this should be a string?
    Data []byte `json:"data"`
}

//
// Client to server nested messages
//

// Type == "send_data_name"
type SendDisplayNameData struct {
    Name string `json:"name"`
}

// Type == "join_room"
type JoinRoomData struct {
    Room string `json:"room"`

    // player's name?
    Name string `json:"name"`
}

// Type == "play_move"
type PlayMoveData struct {
    Cards []game.Card `json:"cards"`
}

//
// Server to client nested messages
//

// Type == "error"
type ErrorData struct {
    Reason string `json:"reason"`
}

// Type == "room_created"
type RoomCreatedData struct {
    RoomId string `json:"room_id"`
}
