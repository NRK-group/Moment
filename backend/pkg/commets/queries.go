package commets

import (
	"fmt"
	"backend/pkg/structs"
)

// Get Comments
// is a method of forum that return all the comment with that specific postID
func GetComments(pID string, database *structs.DB) []structs.Comment {
	rows, err := database.DB.Query("SELECT * FROM Comment WHERE postId = '" + pID + "'")
	var comment structs.Comment
	var comments []structs.Comment
	if err != nil {
		fmt.Print(err)
		return comments
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
	return comments
}
