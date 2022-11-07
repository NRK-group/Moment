package chat

import (
	l "backend/pkg/log"
	"backend/pkg/structs"
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
