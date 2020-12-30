package app

import (
	"encoding/json"
	"github.com/MGC3/bigtwo/backend/app/game"
	"log"
	"math/rand"
)

const (
	maxNumPlayersInRoom = 4
	roomCodeLetters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
	players       [maxNumPlayersInRoom]*player
	receive       chan Message
	toWaitingArea chan roomId
	lobbyHandlers map[string]func(Message)

	// TODO waiting area channel to signal when room is done with the game?
	gameHandlers map[string]func(Message)
	inGame       bool

	//
	// TODO can/should this be refactored to be more "modular"?
	clientIdTurn   int
	numTurnsPassed int

	// Indicates that anything can be played
	lastHandCleared bool
	lastPlayedHand  game.PlayedHand
}

func (r *room) serve() {
	// TODO this is basically the same code as WaitingArea.Serve()
	// generalize code?
	log.Printf("room %s serving...\n", r.id)

	for {
		receive := <-r.receive

		var handler func(Message)
		var ok bool
		if r.inGame {
			handler, ok = r.gameHandlers[receive.Type]
		} else {
			handler, ok = r.lobbyHandlers[receive.Type]
		}

		if !ok {
			log.Printf("Unhandled message type %s for game started = %v\n", receive.Type, r.inGame)
			continue
		}

		handler(receive)

		// TODO is this special case needed? To shut down thread properly
		if receive.Type == "disconnect" && r.numPlayers() == 0 {
			log.Printf("serve() thread for room %s exiting\n", r.id)
			r.toWaitingArea <- r.id
			return
		}
	}
}

func (r *room) handleDisconnect(receive Message) {
	// TODO delete (or invalidate?) player from array
	disconnectedClientId := r.clientIdFromPlayerId(receive.Player.id)

	if disconnectedClientId == -1 {
		log.Printf("room handleDisconnect error - no player found")
		return
	}

	// TODO is this necessary to forward this here?
	receive.Player.toPlayer <- receive

	// move all players to fill in disconnected players
	r.players[disconnectedClientId] = nil
	for i := disconnectedClientId + 1; i < maxNumPlayersInRoom; i++ {
		r.players[i-1] = r.players[i]
	}

	// If no players left, return early
	if r.numPlayers() == 0 {
		return
	}

	// TODO what if game has started?
	if r.inGame {
		if r.numPlayers() == 1 {
			r.inGame = false
		}
		r.clientIdTurn = r.clientIdTurn % r.numPlayers()
		r.pushGameStateToPlayers()
	} else {
		r.pushRoomStateToPlayers()
	}
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
	log.Printf("Got request room state from player %d\n", receive.Player.displayName)
	data := r.roomStateData()
	data.ClientId = r.clientIdFromPlayerId(receive.Player.id)
	msg, err := NewMessage(receive.Player, "room_state", data)
	if err != nil {
		log.Printf("handleRequestRoomState err %v\n", err)
		return
	}
	receive.Player.toPlayer <- msg
}

func (r *room) handleStartGame(receive Message) {
	log.Printf("Got game start request from player %s\n", receive.Player.displayName)

	deck := game.NewDeck()

	for _, player := range r.players {
		if player == nil {
			break
		}

		// TODO make this a constant somewhere?
		// TODO check that player doesn't have 3 twos?
		player.currentHand = deck.Deal(13)
	}

	// TODO first player must have 3 of clubs?
	// or choose a random player to start?
	r.clientIdTurn = rand.Intn(r.numPlayers())
	r.inGame = true
	r.lastHandCleared = true

	inGameMessage, err := NewMessage(nil, "game_started", EmptyData{})

	if err != nil {
		log.Printf("handleStartGame could not create game started msg %v\n", err)
		return
	}

	for _, player := range r.players {
		if player == nil {
			break
		}
		player.toPlayer <- inGameMessage
	}

	// for players stuck on previous game page
	r.pushGameStateToPlayers()
}

func (r *room) handleRequestGameState(receive Message) {
	log.Printf("Got request game state message from player %s\n", receive.Player.displayName)

	gameState := r.gameStateData()
	gameState.UserHand = game.CardListToJson(receive.Player.currentHand)
	gameState.ClientId = r.clientIdFromPlayerId(receive.Player.id)

	send, err := NewMessage(nil, "game_state", gameState)
	if err != nil {
		log.Fatal("could not create game state %v\n", err)
		return
	}
	receive.Player.toPlayer <- send
}

func (r *room) handlePlayMove(receive Message) {
	log.Printf("Got play cards message %v\n", receive)

	receivedClientId := r.clientIdFromPlayerId(receive.Player.id)
	if receivedClientId != r.clientIdTurn {
		log.Printf("handlePlayMove got move from player %d, but turn is %d\n", receivedClientId, r.clientIdTurn)
		sendErrorToPlayer(receive.Player.toPlayer, "It is not your turn")
		return
	}

	var data PlayMoveData
	err := json.Unmarshal(receive.Data, &data)

	if err != nil {
		log.Printf("handlePlayCards failed to unmarshal data %v\n", err)
		sendErrorToPlayer(receive.Player.toPlayer, "Invalid hand")
		return
	}

	cards, err := game.CardListFromJson(data.Cards)
	if err != nil {
		log.Printf("handlePlayCards failed to get card list from data %v\n", err)
		sendErrorToPlayer(receive.Player.toPlayer, "Invalid hand")
		return
	}

	newHand, err := game.NewPlayedHand(cards)
	if err != nil {
		log.Printf("handlePlayCards failed to create played hand from cards %v, err=%v\n", cards, err)
		sendErrorToPlayer(receive.Player.toPlayer, "Invalid hand")
		return
	}

	// bypass check if last hand cleared
	if !r.lastHandCleared {
		newBeatsOld, err := newHand.Beats(r.lastPlayedHand)
		if err != nil || !newBeatsOld {
			log.Printf("handlePlayCards got %v, but doesn't beat last hand %v\n", newHand, r.lastPlayedHand)
			sendErrorToPlayer(receive.Player.toPlayer, "Invalid hand")
			return
		}
	}

	// If we're here, then all the hand can be played as long as the cards exist
	// in the player's current hand
	// RemoveCardsFromList returns the original hand if there was an error
	receive.Player.currentHand, err = game.RemoveCardsFromList(receive.Player.currentHand, cards)
	if err != nil {
		log.Printf("handlePlayCards cards from client not in hand: %v, %v\n", cards, receive.Player.currentHand)
		sendErrorToPlayer(receive.Player.toPlayer, "Invalid hand")
		return
	}

	if len(receive.Player.currentHand) == 0 {
		r.inGame = false
	}

	r.lastHandCleared = false
	r.numTurnsPassed = 0
	r.lastPlayedHand = newHand
	r.clientIdTurn = (r.clientIdTurn + 1) % r.numPlayers()
	r.pushGameStateToPlayers()
}

func (r *room) handlePassMove(receive Message) {
	log.Printf("Got pass move message %v\n", receive)
	receivedClientId := r.clientIdFromPlayerId(receive.Player.id)
	if receivedClientId != r.clientIdTurn {
		log.Printf("handlePlayMove got move from player %d, but turn is %d\n", receivedClientId, r.clientIdTurn)
		sendErrorToPlayer(receive.Player.toPlayer, "It is not your turn")
		return
	}

	if r.lastHandCleared {
		log.Printf("Can't pass on free move\n")
		sendErrorToPlayer(receive.Player.toPlayer, "Pass not allowed")
		return
	}

	r.numTurnsPassed += 1
	if r.numTurnsPassed >= r.numPlayers()-1 {
		r.lastHandCleared = true
	}
	r.clientIdTurn = (r.clientIdTurn + 1) % r.numPlayers()
	r.pushGameStateToPlayers()
}

func (r *room) pushRoomStateToPlayers() {
	if r.inGame {
		log.Printf("pushRoomStateToPlayers called when game started")
		return
	}

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

func (r *room) pushGameStateToPlayers() {

	data := r.gameStateData()
	for clientId, player := range r.players {
		if player == nil {
			break
		}
		data.ClientId = clientId
		data.UserHand = game.CardListToJson(player.currentHand)
		msg, err := NewMessage(player, "game_state", data)
		if err != nil {
			log.Printf("pushGameStateToPlayers err %v\n", err)
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

// creates a GameStateData instance with all the common state populated
// player-specific data (ie, UserHand) is left uninitialized
// Should only be called after the game has started
func (r *room) gameStateData() GameStateData {
	ret := GameStateData{
		AllPlayerHands: []OtherPlayerHand{},
	}

	if r.lastHandCleared {
		// Send empty array to frontend if the hand is cleared
		// ie start of game or everyone passed
		// indicates that any hand can be played by the next player
		ret.LastPlayedHand = []game.JsonCard{}
	} else {
		ret.LastPlayedHand = r.lastPlayedHand.ToJson()
	}

	ret.GameOver = !r.inGame
	for _, player := range r.players {
		if player == nil {
			break
		}

		playerHand := OtherPlayerHand{
			Name:  player.displayName,
			Count: len(player.currentHand),
		}
		ret.AllPlayerHands = append(ret.AllPlayerHands, playerHand)
	}

	if r.numTurnsPassed > 0 {
		ret.LastAction = "passed"
	} else {
		ret.LastAction = "played_hand"
	}

	if r.clientIdTurn < 0 || r.clientIdTurn >= r.numPlayers() {
		log.Printf("bad current client id %d\n", r.clientIdTurn)
		return ret
	}

	currentPlayer := r.players[r.clientIdTurn]

	if currentPlayer == nil {
		log.Printf("current client is nil %d\n", r.clientIdTurn)
		return ret
	}

	ret.CurrentUserTurn = r.players[r.clientIdTurn].displayName
	// TODO detect game over?
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

func generateRandomRoomCode() roomId {
	code := make([]byte, 4)

	for i := range code {
		code[i] = roomCodeLetters[rand.Intn(len(roomCodeLetters))]
	}

	return roomId(code)
}

func newRoom(toWaitingArea chan roomId) *room {
	r := room{
		id: generateRandomRoomCode(),
		// TODO don't initialize everything to nil
		players:         [maxNumPlayersInRoom]*player{nil, nil, nil, nil},
		receive:         make(chan Message),
		toWaitingArea:   toWaitingArea,
		lobbyHandlers:   make(map[string]func(Message)),
		gameHandlers:    make(map[string]func(Message)),
		inGame:          false,
		clientIdTurn:    0,
		numTurnsPassed:  0,
		lastHandCleared: true,
	}

	r.lobbyHandlers["disconnect"] = r.handleDisconnect
	r.lobbyHandlers["join_room"] = r.handleJoinRoom
	r.lobbyHandlers["request_room_state"] = r.handleRequestRoomState
	r.lobbyHandlers["start_game"] = r.handleStartGame

	r.gameHandlers["disconnect"] = r.handleDisconnect
	r.gameHandlers["request_game_state"] = r.handleRequestGameState
	r.gameHandlers["play_move"] = r.handlePlayMove
	r.gameHandlers["pass_move"] = r.handlePassMove

	go r.serve()
	return &r
}
