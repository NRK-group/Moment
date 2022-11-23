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
func AllPost(database *structs.DB) ([]structs.Post, error) {
	var post structs.Post
	var posts []structs.Post
	var err error

	rows, err := database.DB.Query("SELECT * FROM Post ")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&post.PostID, &post.UserID, &post.GroupID, &post.NickName, &post.Content, &post.Image, &post.ImageUpload, &post.NumLikes, &post.CreatedAt, &post.Privacy)
		if post.GroupID == "" {
		arr, _ := commets.GetComments(post.PostID, database)
		post.NumOfComment = len(arr)
		posts = append([]structs.Post{post}, posts...)
		}
	}

	return posts, nil
}


// return all user posts
func AllUserPost(uID string, database *structs.DB) ([]structs.Post, error) {
	var post structs.Post
	var posts []structs.Post
	var err error

	rows, err := database.DB.Query("SELECT * FROM Post ")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&post.PostID, &post.UserID, &post.GroupID, &post.NickName, &post.Content, &post.Image, &post.ImageUpload, &post.NumLikes, &post.CreatedAt, &post.Privacy)
		if post.UserID == uID {
		arr, _ := commets.GetComments(post.PostID, database)
		post.NumOfComment = len(arr)
		posts = append([]structs.Post{post}, posts...)
		}
	}

	return posts, nil
}


// CreatePost
// is a method of database that add post in it.
func CreatePost(userID, groupId, imageUpload, content string, privacy int, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	postID := uuid.NewV4()
	var reUser structs.User
	err2 := auth.GetUser("userId", userID, &reUser, *database)
	if err2 != nil {
		fmt.Print(err2)
		return "CreatePost", err2
	}

	stmt, _ := database.DB.Prepare(`
		INSERT INTO Post VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?)
	`)
	_, err := stmt.Exec(postID, userID, groupId, reUser.FirstName, content, reUser.Avatar, imageUpload, 0, createdAt, privacy)
	if err != nil {
		fmt.Println("inside Create Post", err)
		return "", err
	}
	return postID.String(), nil
}
