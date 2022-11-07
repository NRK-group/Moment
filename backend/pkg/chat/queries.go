package chat

import (
	"time"

	l "backend/pkg/log"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// GetPreviousChat returns the previous chat messages
//
// Param:
//
//	userId: the user id
//	database: the database
func GetPreviousPrivateChat(userId string, database *structs.DB) ([]structs.Chat, error) {
	var prevChat structs.Chat
	var prevChats []structs.Chat
	row, err := database.DB.Query("SELECT * FROM Chat WHERE user1 = ? or user2 = ?", userId, userId)
	if err != nil {
		l.LogMessage("Chat", "GetPreviousPrivateChat - Query", err)
		return prevChats, err
	}
	for row.Next() {
		err = row.Scan(&prevChat.ChatId, &prevChat.GroupId, &prevChat.User1, &prevChat.User2, &prevChat.UpdatedAt)
		if err != nil {
			l.LogMessage("Chat", "GetPreviousPrivateChat - Scan", err)
			return prevChats, err
		}
		prevChats = append([]structs.Chat{prevChat}, prevChats...)

	}
	return prevChats, nil
}

// InsertNewChat inserts a new chat message
//
// Param:
//
//	user1Id: the user id
//	user2Id: the user id
//	database: the database
func InsertNewChat(user1Id string, user2Id string, database *structs.DB) (string, error) {
	stmt, err := database.DB.Prepare("INSERT INTO Chat (chatId, user1, user2, groupId, updatedAt) VALUES (?, ?, ?, ?, ?)")
	chatId := uuid.NewV4().String()
	updateAt := time.Now().String()
	if err != nil {
		l.LogMessage("Chat", "InsertNewChat - Insert Error", err)
		return "", err
	}
	_, err = stmt.Exec(chatId, user1Id, user2Id, "", updateAt)
	if err != nil {
		l.LogMessage("Chat", "InsertNewChat - Exec Error", err)
		return "", err
	}
	return chatId, nil
}
