package member

import (
	"fmt"
	"time"

	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/structs"
)

func AllEventByGroup(groupId string, database *structs.DB) ([]structs.Event, error) {
	var event structs.Event
	var events []structs.Event
	var err error
	rows, err := database.DB.Query("SELECT * FROM Event WHERE groupId = '" + groupId + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&event.EventId, &event.UserId, &event.GroupId, &event.Name, &event.ImageUpload, &event.Description, &event.Location, &event.StartTime, &event.EndTime, &event.CreatedAt)

		events = append([]structs.Event{event}, events...)
	}
	return events, nil
}

func AddEventParticipant(eventId, userId string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	stmt, _ := database.DB.Prepare(`
	INSERT INTO EventParticipant values (?, ?, ?, ?)
`)

	_, err := stmt.Exec(eventId, userId, 0, createdAt)
	if err != nil {
		fmt.Println("inside Create Add Event Participant", err)
		return "", err
	}
	return eventId, nil
}

func CheckIfUserInEvent(eventId, userId string, database *structs.DB) (bool, structs.EventParticipant, error) {
	var holder structs.EventParticipant

	rows, err := database.DB.Query("SELECT * FROM EventParticipant WHERE eventId = '" + eventId + "' AND userId = '" + userId + "'")
	if err != nil {
		fmt.Println(err)
		return false, holder, err
	}
	for rows.Next() {
		rows.Scan(&holder.EventId, &holder.UserId, &holder.Status, &holder.CreatedAt)
	}
	if holder.CreatedAt != "" {
		return true, holder, err
	}
	return false, holder, nil
}

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

	events, erre := AllEventByGroup(groupId, database)
	if erre != nil {
		fmt.Println("inside member - AllEventByGroup", err)
		return "", err
	}

	for _, eventg := range events {
		AddEventParticipant(eventg.EventId, userId, database)
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
		user, err := helper.GetUserInfo(userId, database)
		if err != nil {
			fmt.Print("Get Members", err)
			return members, err
		}
		member = structs.Member{
			CreatedAt: CreatedAt,
			UserId:    userId,
			GroupId:   groupIds,
			UserName:  user.Name,
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
	rows, err := database.DB.Query("SELECT receiverId FROM InviteNotif WHERE groupId = '" + groupId + "' AND receiverId = '" + receiverId + "' AND userId = '" + userId + "'")
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
func AcceptInvitationNotif(groupId, senderId, receiverId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE InviteNotif SET status = ? WHERE groupId = ? AND receiverId = ? AND userId = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec("accepted", groupId, senderId, receiverId)
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
func DeclineInvitationNotif(groupId, senderId, reciverId string, database *structs.DB) error {
	_, err := database.DB.Exec("DELETE FROM InviteNotif WHERE groupId = ? AND receiverId = ? AND userId = ?", groupId, senderId, reciverId)
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

func GetAllInvitationNotif(database *structs.DB) ([]structs.GroupNotif, error) {
	var notif structs.GroupNotif
	var notifs []structs.GroupNotif
	var err error
	rows, err := database.DB.Query("SELECT * FROM InviteNotif")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&notif.GroupID, &notif.UserId, &notif.ReceiverId, &notif.CreatedAt, &notif.Type, &notif.Status, &notif.Read)
		if err != nil {
			return nil, err
		}
		notifs = append([]structs.GroupNotif{notif}, notifs...)
	}
	return notifs, nil
}
