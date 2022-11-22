package structs

import "time"

type Group struct {
	CreatedAt   string   `json:"CreatedAt"`
	Name        string   `json:"Name"`
	GroupID     string   `json:"GroupID"`
	Img         string   `json:"Img"`
	Description string   `json:"Description"`
	Admin       string   `json:"Admin"`
	Members     []Member `json:"Members"`
	Member      bool   `json:"member"`
}
type GroupNotif struct {
	GroupID    string    `json:"groupId"`
	UserId     string    `json:"userId"`
	ReceiverId string    `json:"receiverId"`
	CreatedAt  time.Time `json:"createdAt"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	Read       int       `json:"read"`
}
type GroupNotifWriter struct {
	GroupId    Info      `json:"groupId"`
	UserId     Info      `json:"userId"`
	EventId    Event     `json:"eventId"`
	ReceiverId Info      `json:"receiverId"`
	CreatedAt  time.Time `json:"createdAt"`
	NotifType  string    `json:"type"`
	Status     string    `json:"status"`
	Read       int       `json:"read"`
}
