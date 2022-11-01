package member

import (
	"errors"
	"fmt"
	"time"

	"backend/pkg/structs"
)

// Add member to a group.
func AddMember(groupId, userId string, database *structs.DB) (bool, error) {
	var member structs.Member

	createdAt := time.Now().String()
	// check if the group exist
	group, err1 := GetMembers(groupId, database)
	if err1 != nil || len(group) <= 0 {
		fmt.Println("Error inside AddMember")
		return false, err1
	}
	rows, err := database.DB.Query("SELECT userID FROM GroupMember WHERE groupId = '" + groupId + "' AND userId = '" + userId + "'")
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	for rows.Next() {
		rows.Scan(&member.UserId)
	}
	if member.UserId == "" {

		stmt, _ := database.DB.Prepare(`
		INSERT INTO GroupMember values (?, ?, ?)
	`)
		_, err := stmt.Exec(groupId, userId, createdAt)
		if err != nil {
			fmt.Println("inside AddMember", err)
			return false, err
		}
	} else {
		return false, errors.New("user already exist")
	}
	return true, nil
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
