package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	Clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()

		if err != nil {
			delete(Clients, conn)

			conn.Close()
			break
		}
	}
}

func NotifyAll() {
	for client := range Clients {
		client.WriteJSON(map[string]string{
			"event": "update",
		})
	}
}
