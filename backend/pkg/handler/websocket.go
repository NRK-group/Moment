package handler

import (
	"log"
	"net/http"

	"backend/pkg/auth"
	l "backend/pkg/log"
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
	// get the user id from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		l.LogMessage("Websocket.go", "ServeWs - cokkie not found", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	arrCookie, err := auth.SliceCookie(cookie.Value)
	if err != nil {
		l.LogMessage("Websocket.go", "ServeWs - cokkie not found", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	l.LogMessage("Websocket.go", "ServeWs - cookie value", arrCookie[0])
	client := &wSocket.Client{Hub: hub, UserId: arrCookie[0], Conn: conn, Send: make(chan []byte, 1024)}
	client.Hub.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
