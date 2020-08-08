package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/rooms", CreateRoomHandler).Methods("POST")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
