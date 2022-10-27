package handler

import (
	"log"
	"net/http"
	"strconv"

	wSocket "backend/pkg/websocket"
)

// serveWs handles websocket requests from the peer.
func ServeWs(hub *wSocket.Hub, w http.ResponseWriter, r *http.Request) {
	wSocket.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wSocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// UserId will be replaced by the user id from the cookie
	client := &wSocket.Client{Hub: hub, UserId: strconv.Itoa(len(hub.Clients)), Conn: conn, Send: make(chan []byte, 1024)}
	client.Hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
