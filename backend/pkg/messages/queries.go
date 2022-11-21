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
func InsertMessage(message structs.Message, database *structs.DB) (structs.Message, error) {
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
	UpdateChat(message.ChatId, database)
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
	UpdateChat(message.ChatId, database)
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

// GetLastGroupMessage returns the last message from a chat
//
// Param:
//
// chatId: the id of the chat
// database: the database
func GetLastGroupMessage(chatId string, database *structs.DB) structs.Message {
	var message structs.Message
	row := database.DB.QueryRow("SELECT * FROM GroupMessage WHERE chatId = ? ORDER BY createdAt DESC LIMIT 1", chatId)
	err := row.Scan(&message.MessageId, &message.ReceiverId, &message.SenderId, &message.ChatId, &message.Content, &message.CreatedAt)
	if err != nil {
		return message
	}
	return message
}

// UpdateChat updates the chat
//
// Param:
//
//	chatId: the chat id
//	database: the database
func UpdateChat(chatId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE Chat SET updatedAt = ? WHERE chatId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(time.Now(), chatId)
	if err != nil {
		return err
	}
	return nil
}

// InsertMessageNotif inserts a new message notification into the database or updates it
//
// Param:
//
// chatId: the id of the chat
// receiverId: the id of the receiver
// database: the database
func InsertOrUpdateMessageNotif(chatId string, receiverId string, database *structs.DB) error {
	// Check if the notification already exists in the database
	row := database.DB.QueryRow("SELECT * FROM MessageNotif WHERE chatId = ? AND receiverId = ?", chatId, receiverId)
	var notif structs.MessageNotif
	err := row.Scan(&notif.ChatId, &notif.ReceiverId, &notif.Notif)
	if err != nil {
		// Insert the notification into the database
		stmt, err := database.DB.Prepare("INSERT INTO MessageNotif (chatId, receiverId, notif) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(chatId, receiverId, 1)
		if err != nil {
			return err
		}
		return nil
	}
	// Update the notification in the database
	stmt, err := database.DB.Prepare("UPDATE MessageNotif SET notif = ? WHERE chatId = ? AND receiverId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(notif.Notif+1, chatId, receiverId)
	if err != nil {
		return err
	}
	return nil
}

// GetNotif returns the number of notifications for a user
//
// Param:
//
// chatId: the id of the chat
// receiverId: the id of the receiver
// database: the database
func GetNotif(chatId string, receiverId string, database *structs.DB) int {
	row := database.DB.QueryRow("SELECT notif FROM MessageNotif WHERE chatId = ? AND receiverId = ?", chatId, receiverId)
	var notif int
	err := row.Scan(&notif)
	if err != nil {
		return 0
	}
	return notif
}

// DeleteNotif reset the number of the notification to 0
//
// Param:
//
// chatId: the id of the chat
// currentUser: the id of the current user
// database: the database
func DeleteNotif(chatId string, currentUser string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE MessageNotif SET notif = ? WHERE chatId = ? AND receiverId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, chatId, currentUser)
	if err != nil {
		return err
	}
	return nil
}
