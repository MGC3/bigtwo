package app

import (
    "github.com/gorilla/websocket"
)

type player struct {
    id int
    displayName string
    conn *websocket.Conn
}
