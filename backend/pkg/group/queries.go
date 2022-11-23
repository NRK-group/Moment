package group

import (
	"fmt"
	"time"

	"backend/pkg/chat"
	"backend/pkg/commets"
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

// AllPost
// return all Group post
func AllGroupPosts(groupID string, database *structs.DB) ([]structs.Post, error) {
	var post structs.Post
	var posts []structs.Post
	var err error

	rows, err := database.DB.Query("SELECT * FROM Post WHERE groupid = '" + groupID + "'")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&post.PostID, &post.UserID, &post.GroupID, &post.NickName, &post.Content, &post.Image, &post.ImageUpload, &post.NumLikes, &post.CreatedAt, &post.Privacy)
		arr, err := commets.GetComments(post.PostID, database)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		post.NumOfComment = len(arr)
		posts = append([]structs.Post{post}, posts...)
	}

	return posts, nil
}

// AllGroups
// return all groups
func AllUserGroups(uID string, database *structs.DB) ([]structs.Group, error) {
	var group structs.Group
	var groups []structs.Group
	var err error
	var flag bool
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
		members, err := member.GetMembers(groupId, database)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}

		for _, m := range members {
			if m.UserId == uID {
				flag = true
			}
		}

		if flag {
			group.Members = members
			groups = append([]structs.Group{group}, groups...)
			flag = false
		}
	}
	return groups, nil
}

// ReadGroupNotif is a function that read a group notification.
//
// Param:
//
// userId: the id of the user that read the notification.
// database: the database that the function will use.
func ReadGroupNotif(userId string, database *structs.DB) error {
	_, err := database.DB.Exec("UPDATE InviteNotif SET read = ? WHERE receiverId = ? AND read = ?", 1, userId, 0)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec("UPDATE EventNotif SET read = ? WHERE userId = ? AND read = ?", 1, userId, 0)
	if err != nil {
		return err
	}
	return nil
}
