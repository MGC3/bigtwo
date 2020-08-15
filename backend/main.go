package main

import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
    "sync"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    "github.com/MGC3/bigtwo/backend/app"
)

var frontDesk app.FrontDesk
var frontDeskLock sync.Mutex

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

func main() {
    frontDesk = app.FrontDesk{}
    frontDeskLock = sync.Mutex{}
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/rooms", CreateRoomHandler).Methods("POST")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
