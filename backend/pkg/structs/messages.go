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
	ChatId  string              `json:"chatId"`
	User    map[string]UserInfo `json:"user"`
	Content Message             `json:"content"`
}

type UserInfo struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Img      string `json:"img"`
}
