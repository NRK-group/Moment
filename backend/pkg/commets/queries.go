package commets

import (
	"fmt"
	"time"

	"backend/pkg/structs"
	"backend/pkg/auth"

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
		return  comments, err
	}
	var numLikes int
	var commentId, postId, userId, createdAt, image, content string
	for rows.Next() {
		rows.Scan(&commentId, &postId, &userId, &content, &image, &numLikes, &createdAt)
		comment = structs.Comment{
			CommentID: commentId,
			UserID:    userId,
			PostID:    postId,
			CreatedAt: createdAt,
			Image:     image,
			Content:   content,
			NumLikes:  numLikes,
		}

		comments = append([]structs.Comment{comment}, comments...)
	}
	return comments, err
}

// CreateComment is a method that add a comment.
func CreateComment(userID, postID, content string, database *structs.DB) (string, error) {
	createdAt := time.Now().String()
	commentID := uuid.NewV4()
	var reUser structs.User
	auth.GetUser("userId",userID, &reUser, *database)
	image := reUser.Avatar
	stmt, _ := database.DB.Prepare(`
		INSERT INTO Comment values (?, ?, ?,? ,? ,? ,? )
	`)
	_, err := stmt.Exec(commentID, postID, userID, content, image, 0, createdAt)
	if err != nil {
		return "", err
	}
	return commentID.String(), nil
}
