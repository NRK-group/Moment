package structs

import "time"

type Message struct {
	MessageId   string    `json:"messageId"`
	MessageType string    `json:"type"` // "privateMessage", "groupMessage", or "typing"
	ReceiverId  string    `json:"receiverId"`
	SenderId    string    `json:"senderId"`
	ChatId      string    `json:"chatId"`
	Img         string    `json:"img"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createAt"`
}

type Chat struct {
	ChatId    string    `json:"chatId"`
	GroupId   string    `json:"groupId"`
	User1     string    `json:"user1"`
	User2     string    `json:"user2"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChatWriter struct {
	ChatId  string          `json:"chatId"`
	Type    string          `json:"type"`
	Details Info            `json:"details"`
	Member  map[string]Info `json:"member"`
	Content Message         `json:"content"`
}

type Info struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Name      string `json:"name"`
	Img       string `json:"img"`
}
