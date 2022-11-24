package chat

import (
	"fmt"
	"time"

	"backend/pkg/follow"
	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/messages"
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
	row, err := database.DB.Query("SELECT * FROM Chat WHERE user1 = ? or user2 = ? ORDER BY updatedAt DESC", userId, userId)
	if err != nil {
		l.LogMessage("Chat", "GetPreviousPrivateChat - Query", err)
		return chatList, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&prevChat.ChatId, &prevChat.GroupId, &prevChat.User1, &prevChat.User2, &prevChat.UpdatedAt)
		if err != nil {
			l.LogMessage("Chat", "GetPreviousPrivateChat - Scan", err)
			return chatList, err
		}
		lastMessage := messages.GetLastMessage(prevChat.ChatId, database)
		var msg structs.Message
		if lastMessage != msg {
			m := make(map[string]structs.Info)
			if prevChat.User1 == userId {
				userInfo, _ := helper.GetUserInfo(prevChat.User2, database)

				m[prevChat.User2] = userInfo
				chatList = append(chatList, structs.ChatWriter{
					Type:      "privateMessage",
					ChatId:    prevChat.ChatId,
					Details:   userInfo,
					Member:    m,
					Content:   lastMessage,
					UpdatedAt: prevChat.UpdatedAt,
					Notif:     messages.GetNotif(prevChat.ChatId, userId, database),
				})
			} else {
				userInfo, _ := helper.GetUserInfo(prevChat.User1, database)
				m[prevChat.User1] = userInfo
				chatList = append(chatList, structs.ChatWriter{
					Type:      "privateMessage",
					ChatId:    prevChat.ChatId,
					Details:   userInfo,
					Member:    m,
					Content:   lastMessage,
					UpdatedAt: prevChat.UpdatedAt,
					Notif:     messages.GetNotif(prevChat.ChatId, userId, database),
				})
			}
		}
	}
	return chatList, nil
}

// InsertNewChat inserts a new chat message
//
// Param:
//
//	user1Id: the user id
//	user2Id: the user id
//	database: the database
func InsertNewChat(user1Id string, user2Id string, database *structs.DB) (structs.ChatWriter, error) {
	stmt, err := database.DB.Prepare("INSERT INTO Chat (chatId, user1, user2, groupId, updatedAt) VALUES (?, ?, ?, ?, ?)")
	chatId := uuid.NewV4().String()
	updateAt := time.Now()
	var chat structs.ChatWriter
	if err != nil {
		l.LogMessage("Chat", "InsertNewChat - Insert Error", err)
		return chat, err
	}
	_, err = stmt.Exec(chatId, user1Id, user2Id, "", updateAt)
	if err != nil {
		l.LogMessage("Chat", "InsertNewChat - Exec Error", err)
		return chat, err
	}
	m := make(map[string]structs.Info)
	userInfo, _ := helper.GetUserInfo(user2Id, database)
	m[user2Id] = userInfo
	chat = structs.ChatWriter{
		Type:      "privateMessage",
		ChatId:    chatId,
		Details:   userInfo,
		Member:    m,
		UpdatedAt: updateAt,
	}
	return chat, nil
}

// GetPreviousGroupChat returns the previous chat messages
//
// Param:
//
//	groupId: the group id
//	database: the database
func GetPreviousGroupChat(userId string, database *structs.DB) ([]structs.ChatWriter, error) {
	var prevChatlist []structs.ChatWriter
	groups, err := GetUserGroups(userId, database)
	if err != nil {
		return prevChatlist, err
	}
	for _, group := range groups {
		var prevChat structs.Chat
		var info structs.Info
		row, err := database.DB.Query("SELECT chatId, groupId, updatedAt FROM Chat WHERE groupId = ? ORDER BY updatedAt DESC", group.GroupID)
		if err != nil {
			return prevChatlist, err
		}
		for row.Next() {
			err = row.Scan(&prevChat.ChatId, &prevChat.GroupId, &prevChat.UpdatedAt)
			if err != nil {
				return prevChatlist, err
			}
			m := make(map[string]structs.Info)
			for _, member := range group.Members {
				fmt.Println(member.UserId)
				userInfo, err := helper.GetUserInfo(member.UserId, database)
				if err != nil {
					return prevChatlist, err
				}
				m[member.UserId] = userInfo
			}
			info = structs.Info{
				Id:   group.GroupID,
				Name: group.Name,
				Img:  "images/profile/default-user.svg",
			}
			prevChatlist = append(prevChatlist, structs.ChatWriter{
				Type:      "groupMessage",
				ChatId:    prevChat.ChatId,
				Details:   info,
				Member:    m,
				Content:   messages.GetLastGroupMessage(prevChat.ChatId, database),
				UpdatedAt: prevChat.UpdatedAt,
				Notif:     messages.GetNotif(prevChat.ChatId, userId, database),
			})
		}
	}
	return prevChatlist, nil
}

// GetUserGroups returns all groups that the user is a member of and the member of the group
//
// Param:
//
//	userId: the user id
//	database: the database
func GetUserGroups(userId string, database *structs.DB) ([]structs.Group, error) {
	var group structs.Group
	var groups []structs.Group
	rows, err := database.DB.Query("SELECT * FROM Groups WHERE groupId IN (SELECT groupId FROM GroupMember WHERE userId = ?)", userId)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var groupId, admin, name, description, createdAt string
	for rows.Next() {
		rows.Scan(&groupId, &admin, &name, &description, &createdAt)
		member, err := GetAllMembersOfGroup(groupId, database)
		if err != nil {
			return nil, err
		}
		group = structs.Group{
			CreatedAt:   createdAt,
			Name:        name,
			GroupID:     groupId,
			Description: description,
			Admin:       admin,
			Members:     member,
		}
		groups = append([]structs.Group{group}, groups...)
	}
	return groups, nil
}

// GetAllMembersOfGroup returns all members of a group
//
// Param:
//
//	groupId: the group id
//	database: the database
func GetAllMembersOfGroup(id string, database *structs.DB) ([]structs.Member, error) {
	var members []structs.Member
	var member structs.Member
	rows, err := database.DB.Query("SELECT * FROM GroupMember WHERE groupId = ?", id)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&member.GroupId, &member.UserId, &member.CreatedAt)
		members = append([]structs.Member{member}, members...)
	}
	return members, nil
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
	updateAt := time.Now()
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

// GetFollowingInfo returns the following and follower info of the user
//
// return the following and follower information of the current user
//
// Param:
//
//	userId: the user id
//	database: the database
func GetFollowingInfo(userId string, database *structs.DB) ([]structs.Info, error) {
	var userInfos []structs.Info
	following, err := follow.GetFollowing(userId, database)
	if err != nil {
		return nil, err
	}
	for _, follower := range following {
		userInfo, err := helper.GetUserInfo(follower.FollowingId, database)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}
	return userInfos, nil
}

// DeleteChat deletes a chat
//
// Param:
//
//	chatId: the chat id
//	database: the database
func DeleteChat(chatId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("DELETE FROM Chat WHERE chatId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(chatId)
	if err != nil {
		return err
	}
	return nil
}

// CheckIfChatExists checks if a chat exists and return the chat writer
//
// Param:
//
//	user1: the user id
//	user2: the user id
//	database: the database
func CheckIfChatExists(user1 string, user2 string, database *structs.DB) (bool, structs.ChatWriter) {
	var prevChat structs.Chat
	var chat structs.ChatWriter
	row := database.DB.QueryRow("SELECT * FROM Chat WHERE user1 = ? AND user2 = ? OR user1 = ? AND user2 = ?", user1, user2, user2, user1)
	err := row.Scan(&prevChat.ChatId, &prevChat.GroupId, &prevChat.User1, &prevChat.User2, &prevChat.UpdatedAt)
	if err != nil {
		return false, chat
	}
	if prevChat.User1 == user1 {
		userInfo, err := helper.GetUserInfo(prevChat.User2, database)
		if err != nil {
			return false, chat
		}
		chat = structs.ChatWriter{
			Type:      "privateMessage",
			ChatId:    prevChat.ChatId,
			Details:   userInfo,
			UpdatedAt: time.Now(),
		}
	} else {
		userInfo, err := helper.GetUserInfo(prevChat.User1, database)
		if err != nil {
			return false, chat
		}
		chat = structs.ChatWriter{
			Type:      "privateMessage",
			ChatId:    prevChat.ChatId,
			Details:   userInfo,
			UpdatedAt: time.Now(),
		}
	}
	return true, chat
}

// insert group to the chat
// merge sort the chat by updated at
// return the chat
func ArrangeChat(chat []structs.ChatWriter, group []structs.ChatWriter) []structs.ChatWriter {
	result := []structs.ChatWriter{}
	if len(chat) == 0 {
		return group
	}
	if len(group) == 0 {
		return chat
	}
	// merge sort
	if chat[0].UpdatedAt.After(group[0].UpdatedAt) {
		result = append(result, chat[0])
		result = append(result, ArrangeChat(chat[1:], group)...)
	} else {
		result = append(result, group[0])
		result = append(result, ArrangeChat(chat, group[1:])...)
	}
	return result
}
