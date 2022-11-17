package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"backend/pkg/chat"
	"backend/pkg/follow"
	l "backend/pkg/log"
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
	// Database comnectiomn
	Database *structs.DB
}

// NewHub creates a new hub
func NewHub(DB *structs.DB) *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Database:   DB,
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
				msg, err := messages.InsertMessage(msg, *h.Database)
				if err != nil {
					l.LogMessage("Hub.go", "Run() - InsertMessage", err)
				}
				if _, valid := h.Clients[msg.ReceiverId]; valid {
					resp, _ := json.Marshal(msg)
					h.Clients[msg.ReceiverId].Send <- resp
				}
			}
			if msg.MessageType == "groupMessage" {
				msg, err := messages.InsertGroupMessage(msg, h.Database)
				if err != nil {
					l.LogMessage("Hub.go", "Run() - InsertGroupMessage", err)
				}
				members, err := chat.GetAllMembersOfGroup(msg.ReceiverId, h.Database)
				if err != nil {
					l.LogMessage("Hub.go", "Run() - GetAllMembersOfGroup", err)
				}
				for _, member := range members {
					if member.UserId != msg.SenderId {
						if _, valid := h.Clients[member.UserId]; valid {
							resp, _ := json.Marshal(msg)
							h.Clients[member.UserId].Send <- resp
						}
					}
				}
			}
			if msg.MessageType == "privateMessagetyping" {
				if _, valid := h.Clients[msg.ReceiverId]; valid {
					h.Clients[msg.ReceiverId].Send <- message
				}
			}
			if msg.MessageType == "groupMessagetyping" {
				members, err := chat.GetAllMembersOfGroup(msg.ReceiverId, h.Database)
				if err != nil {
					l.LogMessage("Hub.go", "Run() - GetAllMembersOfGroup", err)
				}
				for _, member := range members {
					if member.UserId != msg.SenderId {
						if _, valid := h.Clients[member.UserId]; valid {
							resp, _ := json.Marshal(msg)
							h.Clients[member.UserId].Send <- resp
						}
					}
				}
			}
			if msg.MessageType == "acceptFollowRequest" {
				follow.AcceptFollow(msg.ReceiverId, msg.SenderId, h.Database)
				if _, valid := h.Clients[msg.ReceiverId]; valid {
					h.Clients[msg.ReceiverId].Send <- message
				}
			}
			if msg.MessageType == "declineFollowRequest" {
				follow.DeclineFollow(msg.ReceiverId, msg.SenderId, h.Database)
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
