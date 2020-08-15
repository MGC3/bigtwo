package app


const (
    maxNumPlayersInRoom = 4
)

type roomState int

const (
    roomWaitingForPlayers roomState = iota
    roomGameInProgress
    roomGameFinished
)

// Represents a single room.
// A room can have 0 to maxNumPlayersinRoom players.
// The room's state (waiting, in progress, finished) is used
// to determine if someone should be allowed to join a room.
type room struct {
    players []player
    state roomState
}

// Manages all of the active rooms.
// Determines whether incoming clients can create/join a room.
// 
type FrontDesk struct {
    activeRooms map[int]room
    nextId int
}

// Creates a new room witha unique Id
func (f *FrontDesk) CreateRoom() (int, error) {
    currentRoomId := f.nextId
    f.nextId += 1
    f.activeRooms[currentRoomId] = room{players: []player{}, state: roomWaitingForPlayers}
    return currentRoomId, nil
}
