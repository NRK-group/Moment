package messages

import (
	"time"

	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// InsertMessage inserts a new message into the database
//
// # Param
//
// message struct: the message information
// database: the database
func InsertMessage(message structs.Message, database structs.DB) (structs.Message, error) {
	messageId := uuid.NewV4().String()
	createdAt := time.Now()
	msg := structs.Message{}
	stmt, err := database.DB.Prepare("INSERT INTO PrivateMessage (messageId, chatId, senderId, receiverId, content, createdAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return msg, err
	}
	_, err = stmt.Exec(messageId, message.ChatId, message.SenderId, message.ReceiverId, message.Content, createdAt)
	if err != nil {
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
		return messages, err
	}
	for row.Next() {
		err = row.Scan(&message.MessageId, &message.ChatId, &message.SenderId, &message.ReceiverId, &message.Content, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// InsertGroupMessage inserts a new message into the database
//
// Param:
//
// message struct: the message information
// database: the database
func InsertGroupMessage(message structs.Message, database *structs.DB) (structs.Message, error) {
	messageId := uuid.NewV4().String()
	createdAt := time.Now()
	msg := structs.Message{}
	stmt, err := database.DB.Prepare("INSERT INTO GroupMessage (messageId, groupId, senderId, chatId, content, createdAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return msg, err
	}
	_, err = stmt.Exec(messageId, message.ReceiverId, message.SenderId, message.ChatId, message.Content, createdAt)
	if err != nil {
		return msg, err
	}
	msg = structs.Message{
		MessageId:   messageId,
		MessageType: "groupMessage",
		ChatId:      message.ChatId,
		SenderId:    message.SenderId,
		ReceiverId:  message.ReceiverId,
		Img:         message.Img,
		Content:     message.Content,
		CreatedAt:   createdAt,
	}
	return msg, nil
}

func GetGroupMessages(chatId string, database structs.DB) ([]structs.Message, error) {
	var message structs.Message
	var messages []structs.Message
	row, err := database.DB.Query("SELECT * FROM GroupMessage WHERE ChatId = ?", chatId)
	if err != nil {
		return messages, err
	}
	for row.Next() {
		err = row.Scan(&message.MessageId, &message.ReceiverId, &message.SenderId, &message.ChatId, &message.Content, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// GetLastMessage returns the last message from a chat
//
// Param:
//
// chatId: the id of the chat
// database: the database
func GetLastMessage(chatId string, database *structs.DB) structs.Message {
	var message structs.Message
	row := database.DB.QueryRow("SELECT * FROM PrivateMessage WHERE chatId = ? ORDER BY createdAt DESC LIMIT 1", chatId)
	err := row.Scan(&message.MessageId, &message.ChatId, &message.SenderId, &message.ReceiverId, &message.Content, &message.CreatedAt)
	if err != nil {
		return message
	}
	return message
}
