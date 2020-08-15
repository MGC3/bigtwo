package main

import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
    "strconv"
    "sync"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/gorilla/websocket"

    "github.com/MGC3/bigtwo/backend/app"
)

var frontDesk = app.NewFrontDesk()
var frontDeskLock = sync.Mutex{}
var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type CreateRoomResponseBody struct {
    RoomId int `json: "roomId"`
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    frontDeskLock.Lock()
    roomId, err := frontDesk.CreateRoom()
    frontDeskLock.Unlock()

    if err != nil {
        log.Fatal("error")
    }

    fmt.Printf("Got create room request. New room is %d\n", roomId)
    
    w.Header().Set("Content-Type", "application/json")
    response := CreateRoomResponseBody{
        RoomId: roomId,
    }

    json.NewEncoder(w).Encode(response)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roomId, err := strconv.Atoi(vars["roomId"])

    fmt.Printf("Got request to join room %s\n", roomId)
    if err != nil {
        log.Fatal("bad room id")
    }

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("User has connected on %v\n", conn)
    frontDeskLock.Lock()
    err = frontDesk.JoinRoom(roomId, conn)
    frontDeskLock.Unlock()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Successfully joined room %d\n", roomId)
}

func main() {

    // I am a web programmer
    upgrader.CheckOrigin = func(r *http.Request) bool {
        return true
    }

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/rooms", CreateRoomHandler).Methods("POST")
    r.HandleFunc("/rooms/{roomId}", JoinRoomHandler).Methods("GET")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
