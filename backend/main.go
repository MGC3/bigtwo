package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	//    "encoding/json"
	//    "strconv"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/MGC3/bigtwo/backend/app"
)

var waitingArea = app.NewWaitingArea()
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO sync better?
var bigLock = sync.Mutex{}

// TODO only one GET request is serviced at a time
func EstablishWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	bigLock.Lock()
	defer bigLock.Unlock()
	fmt.Println("EstablishWebsocketConnection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Could not establish websocket connection")
		return
	}

	// TODO can I convert this to use channels? It would be nice to get rid of the locks.
	uninitializedPlayer := app.UninitializedPlayer(conn)
	toWaitingArea, err := app.NewMessage(uninitializedPlayer, "new_connected_player", app.EmptyData{})
	if err != nil {
		log.Printf("EstablishWebsocketConnection failed %v\n", err)
		return
	}
	waitingArea.Receive <- toWaitingArea
	fmt.Println("Added new player to waiting area")
}

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	log.Printf("Backend running with seed %v\n", seed)
	go waitingArea.Serve()

	// I am a web programmer
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", EstablishWebsocketConnection).Methods("GET")
	//r.HandleFunc("/rooms/{roomId}", JoinRoomHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
