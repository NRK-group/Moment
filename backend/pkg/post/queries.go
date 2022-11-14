package post

import (
	"fmt"
	"time"

	"backend/pkg/auth"
	"backend/pkg/commets"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// AllPost
// return all post
func AllPost(uID string, database *structs.DB) ([]structs.Post, error) {
	var post structs.Post
	var posts []structs.Post
	var err error

	rows, err := database.DB.Query("SELECT * FROM Post ")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var numLikes int
	var postId, userId, groupId, content, image, createdAt string
	for rows.Next() {
		rows.Scan(&postId, &userId, &groupId, &content, &image, &numLikes, &createdAt)
		post = structs.Post{
			PostID:    postId,
			UserID:    userId,
			GroupID:   groupId,
			CreatedAt: createdAt,
			Content:   content,
			Image:     image,
			NumLikes:  numLikes,
		}
		arr, _ := commets.GetComments(postId, database)
		post.NumOfComment = len(arr)
		posts = append([]structs.Post{post}, posts...)
	}

	return posts, nil
}

// CreatePost
// is a method of database that add post in it.
func CreatePost(userID, groupId, image, content string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	postID := uuid.NewV4()
	var reUser structs.User
	err2 := auth.GetUser("userId", userID, &reUser, *database)
	if err2 != nil {
		fmt.Print(err2)
		return  "CreatePost", err2
	}

	stmt, _ := database.DB.Prepare(`
		INSERT INTO Post (postId, userId, groupId, content, image, numLikes, createdAt ) values (?, ?, ?, ?, ?, ?, ?)
	`)
	_, err := stmt.Exec(postID, userID+"-"+reUser.NickName, groupId, content, image, 0, createdAt)
	if err != nil {
		fmt.Println("inside Create Post", err)
		return "", err
	}
	return postID.String(), nil
}
