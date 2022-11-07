package group

import (
	"fmt"
	"time"

	"backend/pkg/chat"
	"backend/pkg/member"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// AllGroups
// return all groups
func AllGroups(uID string, database *structs.DB) ([]structs.Group, error) {
	var group structs.Group
	var groups []structs.Group
	var err error
	rows, err := database.DB.Query("SELECT * FROM Groups ")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var groupId, admin, name, description, createdAt string
	for rows.Next() {
		rows.Scan(&groupId, &admin, &name, &description, &createdAt)
		group = structs.Group{
			CreatedAt:   createdAt,
			Name:        name,
			GroupID:     groupId,
			Description: description,
			Admin:       admin,
		}
		groups = append([]structs.Group{group}, groups...)
	}
	return groups, nil
}

// CreateGroup
// is a method of database that add a group.
func CreateGroup(name, description, admin string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	groupId := uuid.NewV4()
	stmt, _ := database.DB.Prepare(`
		INSERT INTO Groups (groupId, admin, name, description, createdAt ) values (?, ?, ?, ?, ?)
	`)
	_, err := stmt.Exec(groupId, admin, name, description, createdAt)
	if err != nil {
		fmt.Println("inside Create Group", err)
		return "", err
	}
	member.AddMember(groupId.String(), admin, database)
	chat.InsertNewGroupChat(groupId.String(), database)
	return groupId.String(), nil
}
