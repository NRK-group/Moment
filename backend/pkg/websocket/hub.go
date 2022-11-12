package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"backend/pkg/messages"
	"backend/pkg/structs"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[string]*Client
	// Inbound messages from the clients.
	Broadcast chan []byte
	// Register requests from the clients.
	Register chan *Client
	// Unregister requests from clients.
	Unregister chan *Client
}

// NewHub creates a new hub
func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
	}
}

// Run the hub and listen for new connections
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserId] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserId]; ok {
				delete(h.Clients, client.UserId)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			var msg structs.Message
			json.Unmarshal(message, &msg)
			if msg.MessageType == "privateMessage" {
				if _, valid := h.Clients[msg.ReceiverId]; valid {
					msg, err := messages.InsertMessage(msg, *h.Clients[msg.ReceiverId].Database)
					if err != nil {
						fmt.Println("error inserting message", err)
					}
					resp, _ := json.Marshal(msg)
					h.Clients[msg.ReceiverId].Send <- resp
				}
			}
			if msg.MessageType == "typing" {
				if _, valid := h.Clients[msg.ReceiverId]; valid {
					h.Clients[msg.ReceiverId].Send <- message
				}
			}
			// for userId, client := range h.Clients {
			// 	select {
			// 	case client.Send <- message:
			// 	default:
			// 		close(client.Send)
			// 		delete(h.Clients, userId)
			// 	}
			// }
		}
	}
}

func (h *Hub) LogConns() {
	for {
		fmt.Println(len(h.Clients), "clients connected")
		for userId := range h.Clients {
			fmt.Printf("client %v have %v connections\n", userId, len(h.Clients))
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
