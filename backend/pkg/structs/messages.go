package structs

type Message struct {
	MessageId   string `json:"messageId"`
	MessageType string `json:"type"` // "privateMessage", "groupMessage", or "typing"
	ReceiverId  string `json:"receiverId"`
	SenderId    string `json:"senderId"`
	ChatId      string `json:"chatId"`
	Img         string `json:"img"`
	Content     string `json:"content"`
	CreateAt    string `json:"createAt"`
}
