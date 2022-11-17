package commets

import (
	"fmt"
	"time"

	"backend/pkg/auth"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

// Get Comments
// is a method of forum that return all the comment with that specific postID
func GetComments(pID string, database *structs.DB) ([]structs.Comment, error) {
	rows, err := database.DB.Query("SELECT * FROM Comment WHERE postId = '" + pID + "'")
	var comment structs.Comment
	var comments []structs.Comment
	if err != nil {
		fmt.Print(err)
		return comments, err
	}

	for rows.Next() {
		rows.Scan(&comment.CommentID, &comment.CommentName, &comment.PostID, &comment.UserID, &comment.Content, &comment.Image, &comment.ImageUpload, &comment.NumLikes, &comment.CreatedAt)
		comments = append([]structs.Comment{comment}, comments...)
	}
	return comments, err
}

// CreateComment is a method that add a comment.
func CreateComment(userID, postID, content string, database *structs.DB) (string, error) {
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	commentID := uuid.NewV4()
	var reUser structs.User
	err := auth.GetUser("userId", userID, &reUser, *database)
	if err != nil {
		fmt.Print(err)
		return "comments", err
	}

	image := reUser.Avatar 

	stmt, _ := database.DB.Prepare(`
		INSERT INTO Comment values (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	_, err = stmt.Exec(commentID, reUser.NickName, postID, userID, content, image, "", 0, createdAt)
	if err != nil {
		return "", err
	}
	return commentID.String(), nil
}