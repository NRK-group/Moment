package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func Reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}
