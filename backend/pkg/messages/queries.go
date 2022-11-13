package messages

import (
	"time"

	l "backend/pkg/log"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// InsertMessage inserts a new message into the database
//
// # Param
//
// senderId: the id of the user sending the message
// chatId: the id of the chat the message is being sent to
// receiverId: the id of the user receiving the message
// message: the message to be sent
// database: the database to insert the message into
func InsertMessage(message structs.Message, database structs.DB) (structs.Message, error) {
	messageId := uuid.NewV4().String()
	createdAt := time.Now()
	msg := structs.Message{}
	stmt, err := database.DB.Prepare("INSERT INTO PrivateMessage (messageId, chatId, senderId, receiverId, content, createdAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		l.LogMessage("Messages.go", "InsertMessage - Inserting", err)
		return msg, err
	}
	_, err = stmt.Exec(messageId, message.ChatId, message.SenderId, message.ReceiverId, message.Content, createdAt)
	if err != nil {
		l.LogMessage("Messages.go", "InsertMessage - Executing", err)
		return msg, err
	}
	msg = structs.Message{
		MessageId:   messageId,
		MessageType: "privateMessage",
		ChatId:      message.ChatId,
		SenderId:    message.SenderId,
		ReceiverId:  message.ReceiverId,
		Img:         message.Img,
		Content:     message.Content,
		CreatedAt:   createdAt,
	}
	return msg, nil
}

// GetMessages returns the messages from a chat
//
// Param:
//
// chatId: the id of the chat
//
//	database: the database
func GetPrivateMessages(chatId string, database structs.DB) ([]structs.Message, error) {
	var message structs.Message
	var messages []structs.Message
	row, err := database.DB.Query("SELECT * FROM PrivateMessage WHERE chatId = ?", chatId)
	if err != nil {
		l.LogMessage("Messages.go", "GetMessages - Query", err)
		return messages, err
	}
	for row.Next() {
		err = row.Scan(&message.MessageId, &message.ChatId, &message.SenderId, &message.ReceiverId, &message.Content, &message.CreatedAt)
		if err != nil {
			l.LogMessage("Messages.go", "GetMessages - Scan", err)
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
