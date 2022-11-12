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
func GetPreviousPrivateChat(userId string, database *structs.DB) ([]structs.ChatWriter, error) {
	var prevChat structs.Chat
	var chatList []structs.ChatWriter
	row, err := database.DB.Query("SELECT * FROM Chat WHERE user1 = ? or user2 = ?", userId, userId)
	if err != nil {
		l.LogMessage("Chat", "GetPreviousPrivateChat - Query", err)
		return chatList, err
	}
	for row.Next() {
		err = row.Scan(&prevChat.ChatId, &prevChat.GroupId, &prevChat.User1, &prevChat.User2, &prevChat.UpdatedAt)
		if err != nil {
			l.LogMessage("Chat", "GetPreviousPrivateChat - Scan", err)
			return chatList, err
		}
		m := make(map[string]structs.UserInfo)
		if prevChat.User1 == userId {
			userInfo, _ := GetUserInfo(prevChat.User2, database)
			m[prevChat.User2] = userInfo
			chatList = append([]structs.ChatWriter{{
				ChatId: prevChat.ChatId,
				User:   m,
			}}, chatList...)
		} else {
			userInfo, _ := GetUserInfo(prevChat.User1, database)
			m[prevChat.User1] = userInfo
			chatList = append([]structs.ChatWriter{{
				ChatId: prevChat.ChatId,
				User:   m,
			}}, chatList...)
		}
	}
	l.LogMessage("Chat", "GetPreviousPrivateChat - ChatList", chatList)
	return chatList, nil
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
	updateAt := time.Now()
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

// InsertNewGroupChat inserts a new group chat message
//
// Param:
//
//	groupId: the group id
//	database: the database
func InsertNewGroupChat(groupId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("INSERT INTO Chat (chatId, groupId, user1, user2, updatedAt) VALUES (?, ?, ?, ?, ?)")
	chatId := uuid.NewV4().String()
	updateAt := time.Now().String()
	if err != nil {
		l.LogMessage("Chat", "InsertNewGroupChat - Insert Error", err)
		return err
	}
	_, err = stmt.Exec(chatId, groupId, "", "", updateAt)
	if err != nil {
		l.LogMessage("Chat", "InsertNewGroupChat - Exec Error", err)
		return err
	}
	return nil
}

// GetUserInfo returns the user info for the chat writer
//
// Param:
//
//	userId: the user id
//	database: the database
func GetUserInfo(userId string, database *structs.DB) (structs.UserInfo, error) {
	var userInfo structs.UserInfo
	var user structs.User
	stmt, err := database.DB.Query("SELECT userId, firstName, lastName, nickName, avatar FROM User WHERE userId = ?", userId)
	if err != nil {
		l.LogMessage("Chat", "GetUserInfo - Query Error", err)
		return userInfo, err
	}
	for stmt.Next() {
		err = stmt.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.NickName, &user.Avatar)
		if err != nil {
			l.LogMessage("Chat", "GetUserInfo - Scan Error", err)
			return structs.UserInfo{}, err
		}
		userInfo = structs.UserInfo{
			UserId: user.UserId,
			Img:    user.Avatar,
		}
		if user.NickName != "" {
			userInfo.Username = user.NickName
		} else {
			userInfo.Username = user.FirstName + " " + user.LastName
		}
	}
	return userInfo, nil
}
