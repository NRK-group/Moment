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
		MessageType: message.MessageType,
		ChatId:      message.ChatId,
		SenderId:    message.SenderId,
		ReceiverId:  message.ReceiverId,
		Img:         message.Img,
		Content:     message.Content,
		CreatedAt:   createdAt,
	}
	return msg, nil
}