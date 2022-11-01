package member

import (
	"fmt"
	"time"

	l "backend/pkg/log"
	"backend/pkg/structs"
)

// Add member to a group.
func AddMember(groupId, userId string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()

	// check if the group exist
	group, err1 := GetMembers(groupId, database)
	if err1 != nil || len(group) <= 0 {
		fmt.Println("Error inside AddMember")
		return "error - group doesn't exist", err1
	}

	stmt, _ := database.DB.Prepare(`
		INSERT INTO GroupMember values (?, ?, ?)
	`)
	_, err := stmt.Exec(groupId, userId, createdAt)
	if err != nil {
		fmt.Println("inside AddMember", err)
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
		member = structs.Member{
			CreatedAt: CreatedAt,
			UserId:    userId,
			GroupId:   groupIds,
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
// type: the type of the invitation (join, invite)
// database: the database
func AddMemberNotif(groupId, userId, receiverId, typeNotif string, database *structs.DB) error {
	createdAt := time.Now().String()
	// check if the groupid and ReceiverId exists already
	rows, err := database.DB.Query("SELECT receiverId FROM InviteNotif WHERE groupId = '" + groupId + "' AND receiverId = '" + receiverId + "'")
	if err != nil {
		l.LogMessage("Member.go", "AddMemberNotif", err)
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
		l.LogMessage("Member.go", "AddMemberNotif", err)
	}
	_, err = stmt.Exec(groupId, userId, receiverId, createdAt, typeNotif, "pending", 0)
	if err != nil {
		l.LogMessage("Member.go", "AddMemberNotif", err)
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
func AcceptMemberNotif(groupId, userId, receiverId string, database *structs.DB) error {
	AddMember(groupId, userId, database)
	stmt, err := database.DB.Prepare("UPDATE InviteNotif SET status = ? WHERE groupId = ? AND userId = ? AND receiverId = ?")
	if err != nil {
		l.LogMessage("Member.go", "AcceptMemberNotif", err)
	}
	_, err = stmt.Exec("accepted", groupId, userId, receiverId)
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
func DeclineMemberNotif(groupId, userId, receiverId string, database *structs.DB) error {
	_, err := database.DB.Exec("DELETE FROM InviteNotif WHERE groupId = ? AND userId = ? AND receiverId = ?", groupId, userId, receiverId)
	if err != nil {
		l.LogMessage("Member.go", "DeclineMemberNotif", err)
		return err
	}
	return nil
}