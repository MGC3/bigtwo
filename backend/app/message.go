package app

import (
	//"log"
	"encoding/json"
	"github.com/MGC3/bigtwo/backend/app/game"
)

// This type defines the outer message type, which contains
// a string that identifies what kind of message is stored
// as bytes in data
type Message struct {
	PlayerId playerId `json:"-"`
	Type     string   `json:"type"`

	// TODO omit empty?
	// TODO can this be an empty interface? It might be nice to have nested
	// messages passed around internally that don't need to be marshalled/unmarshalled
	Data json.RawMessage `json:"data"`
	//Data interface{} `json:"data"`
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
	Room roomId `json:"room"`

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
	RoomId roomId `json:"room_id"`
}

type EmptyData struct {
}

func NewMessage(id playerId, messageType string, data interface{}) (Message, error) {
	packedData, err := json.Marshal(data)
	if err != nil {
		return Message{}, nil
	}

	m := Message{
		PlayerId: id,
		Type:     messageType,
		Data:     json.RawMessage(packedData),
	}

	return m, nil
}
