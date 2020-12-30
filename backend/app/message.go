package app

import (
	"encoding/json"
	"github.com/MGC3/bigtwo/backend/app/game"
	"log"
)

// This type defines the outer message type, which contains
// a string that identifies what kind of message is stored
// as bytes in data
type Message struct {
	// TODO abstract concept of a 'sender'?
	Player *player `json:"-"`

	Type string `json:"type"`

	// TODO add 'internal' bool to be able to tell if the message was received
	// from a client or if the message was created internally on the backend
	// for validation

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
	RoomId roomId `json:"room"`

	// player's name?
	Name string `json:"name"`
}

// Type == "play_move"
type PlayMoveData struct {
	Cards []game.JsonCard `json:"cards"`
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

// Type == "room_state"
type RoomStateData struct {
	PlayerNames []string `json:"players"`
	ClientId    int      `json:"client_id"`
}

// Type == "game_state"
type GameStateData struct {
	UserHand        []game.JsonCard   `json:"user_hand"`
	AllPlayerHands  []OtherPlayerHand `json:"all_player_hands"`
	LastPlayedHand  []game.JsonCard   `json:"last_played_hand"`
	LastAction      string            `json:"last_action"`
	CurrentUserTurn string            `json:"current_user_turn"`
	ClientId        int               `json:"client_id"`
	GameOver        bool              `json:"game_over"`
}

type OtherPlayerHand struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

//
// Internal messages sent between server threads
type EmptyData struct {
}

func NewMessage(player *player, messageType string, data interface{}) (Message, error) {
	packedData, err := json.Marshal(data)
	if err != nil {
		return Message{}, nil
	}

	m := Message{
		Player: player,
		Type:   messageType,
		Data:   json.RawMessage(packedData),
	}

	return m, nil
}

func sendErrorToPlayer(toPlayer chan Message, errorString string) {
	log.Printf("error: %s\n", errorString)
	msg, _ := NewMessage(nil, "error", ErrorData{Reason: errorString})
	toPlayer <- msg
}
