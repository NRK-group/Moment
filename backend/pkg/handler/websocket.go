package handler

import (
	"log"
	"net/http"

	"backend/pkg/auth"
	l "backend/pkg/log"
	wSocket "backend/pkg/websocket"
)

// serveWs handles websocket requests from the peer.
//
// Param:
//
//	hub: the hub that contains the clients connected to the websocket
//	w: the response writer
//	r: the request
func (database *Env) ServeWs(hub *wSocket.Hub, w http.ResponseWriter, r *http.Request) {
	wSocket.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wSocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		l.LogMessage("Websocket.go", "ServeWs - cookie not found", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	arrCookie, err := auth.SliceCookie(cookie.Value)
	if err != nil {
		l.LogMessage("Websocket.go", "ServeWs - slicing the cookie", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	client := &wSocket.Client{Hub: hub, UserId: arrCookie[0], Conn: conn, Send: make(chan []byte, 1024)}
	client.Hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
