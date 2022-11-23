package member

import (
	"fmt"
	"time"

	"backend/pkg/auth"
	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/structs"
)

// Add member to a group.
func AddMember(groupId, userId string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	// check if the group exist
	// decide for the future of the this code
	// group, err1 := GetMembers(groupId, database)
	// if err1 != nil || len(group) <= 0 {
	// 	fmt.Println("Error inside AddMember")
	// 	return "error - group doesn't exist", err1
	// }
	stmt, err := database.DB.Prepare("INSERT INTO GroupMember values (?, ?, ?)")
	if err != nil {
		l.LogMessage("Member.go", "AddMember", err)
		return "", err
	}
	_, err = stmt.Exec(groupId, userId, createdAt)
	if err != nil {
		l.LogMessage("Member.go", "AddMember", err)
		return "", err
	}
	return groupId, nil
}

// Get Members
// is a method of forum that return all the members from a group
func GetMembers(groupId string, database *structs.DB) ([]structs.Member, error) {
	rows, err := database.DB.Query("SELECT * FROM GroupMember WHERE groupId = '" + groupId + "'")

	var member structs.Member
	var members []structs.Member

	if err != nil {
		fmt.Print(err)
		return members, err
	}

	var groupIds, userId, CreatedAt string
	for rows.Next() {
		rows.Scan(&groupIds, &userId, &CreatedAt)
		var reUser structs.User
		err := auth.GetUser("userId", userId, &reUser, *database)
		if err != nil {
			fmt.Print("Get Members", err)
			return members, err
		}
		member = structs.Member{
			CreatedAt: CreatedAt,
			UserId:    userId,
			GroupId:   groupIds,
			UserName:  reUser.NickName,
		}

		members = append([]structs.Member{member}, members...)
	}
	return members, nil
}

// AddMemberNotif
//
// Param:
//
// groupId: the id of the group
// userId: the id of the user
// receiverId: the id of the user that will receive the invitation
// typeNotif: the type of the invitation (join, invite)
// database: the database
func AddInvitationNotif(groupId, userId, receiverId, typeNotif string, database *structs.DB) error {
	createdAt := time.Now()
	// check if the groupid and ReceiverId exists already
	rows, err := database.DB.Query("SELECT receiverId FROM InviteNotif WHERE groupId = '" + groupId + "' AND receiverId = '" + receiverId + "'")
	if err != nil {
		return err
	}
	var receiverIds string
	for rows.Next() {
		rows.Scan(&receiverIds)
	}
	if receiverIds != "" {
		return nil
	}
	stmt, err := database.DB.Prepare(`
		INSERT INTO InviteNotif values (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(groupId, userId, receiverId, createdAt, typeNotif, "pending", 0)
	if err != nil {
		return err
	}
	return nil
}

// AcceptMemberNotif
//
// Param:
//
// groupId: the id of the group
// userId: the id of the user
// receiverId: the id of the user that will receive the invitation
// database: the database
func AcceptInvitationNotif(groupId, userId string, database *structs.DB) error {
	AddMember(groupId, userId, database)
	stmt, err := database.DB.Prepare("UPDATE InviteNotif SET status = ? WHERE groupId = ? AND receiverId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec("accepted", groupId, userId)
	if err != nil {
		l.LogMessage("Member.go", "AcceptMemberNotif", err)
		return err
	}
	return nil
}

// DeclineMemberNotif
//
// Param:
//
// groupId: the id of the group
// userId: the id of the user
// receiverId: the id of the user that will receive the invitation
// database: the database
func DeclineInvitationNotif(groupId, userId string, database *structs.DB) error {
	_, err := database.DB.Exec("DELETE FROM InviteNotif WHERE groupId = ? AND receiverId = ?", groupId, userId)
	if err != nil {
		l.LogMessage("Member.go", "DeclineMemberNotif", err)
		return err
	}
	return nil
}

// GetInvitationNotif get all invite notifications
//
//	Param:
//
//	userId: the user id
//	database: the database
func GetInvitationNotif(userId string, database *structs.DB) ([]structs.GroupNotifWriter, error) {
	var notif structs.GroupNotif
	var notifs []structs.GroupNotifWriter
	var err error
	rows, err := database.DB.Query("SELECT * FROM InviteNotif WHERE receiverId = '" + userId + "'")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&notif.GroupID, &notif.UserId, &notif.ReceiverId, &notif.CreatedAt, &notif.Type, &notif.Status, &notif.Read)
		if err != nil {
			return nil, err
		}
		user, err := helper.GetUserInfo(notif.UserId, database)
		if err != nil {
			return nil, err
		}
		receiver, err := helper.GetUserInfo(notif.ReceiverId, database)
		if err != nil {
			return nil, err
		}
		group, err := helper.GetGroupInfo(notif.GroupID, database)
		if err != nil {
			return nil, err
		}
		notifWriter := structs.GroupNotifWriter{
			GroupId:    group,
			UserId:     user,
			ReceiverId: receiver,
			CreatedAt:  notif.CreatedAt,
			NotifType:  notif.Type,
			Status:     notif.Status,
			Read:       notif.Read,
		}
		notifs = append(notifs, notifWriter)
	}
	return notifs, nil
}
