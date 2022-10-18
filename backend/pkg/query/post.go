package query

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type DB struct {
	*sql.DB
}

type Post struct {
	PostID       string
	UserID       string
	Date         string
	CreatedAt    string
	Content      string
	GroupID      string
	Image        string
	NumOfComment int
	NumLikes     int
}

type Comment struct {
	CommentID string
	UserID    string
	PostID    string
	CreatedAt string
	Image     string
	Content   string
	NumLikes  int
}

// CreatePost
// is a method of database that add post in it.
func (databse *DB) CreatePost(userID, groupId, image, content string) (string, error) {
	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	postID := uuid.NewV4()
	stmt, _ := databse.DB.Prepare(`
		INSERT INTO Post (postId, userId, groupId, content, image, numLikes, createdAt ) values (?, ?, ?, ?, ?, ?, ?)
	`)
	_, err := stmt.Exec(postID, userID, groupId, content, image, 0, createdAt)
	if err != nil {
		return "", err
	}
	return postID.String(), nil
}

// Get Comments
// is a method of forum that return all the comment with that specific postID
func (databse *DB) GetComments(pID string) []Comment {
	rows, err := databse.DB.Query("SELECT * FROM Comment WHERE postId = '" + pID + "'")
	var comment Comment
	var comments []Comment
	if err != nil {
		fmt.Print(err)
		return comments
	}
	var numLikes int
	var commentId, postId, userId, createdAt, image, content string
	for rows.Next() {
		rows.Scan(&commentId, &postId, &userId, &content, &image, &numLikes, &createdAt)
		comment = Comment{
			CommentID: commentId,
			UserID:    userId,
			PostID:    postId,
			CreatedAt: createdAt,
			Image:     image,
			Content:   content,
			NumLikes:  numLikes,
		}

		comments = append([]Comment{comment}, comments...)
	}
	return comments
}

// AllPost
// return all post
func (databse *DB) AllPost(uID string) []Post {
	var post Post
	var posts []Post
	var err error

	rows, err := databse.DB.Query("SELECT * FROM Post ")
	if err != nil {
		fmt.Print(err)
		return posts
	}
	var numLikes int
	var postId, userId, groupId, content, image, createdAt string
	for rows.Next() {
		rows.Scan(&postId, &userId, &groupId, &content, &image, &numLikes, &createdAt)
		post = Post{
			PostID:       postId,
			UserID:       userId,
			GroupID:      groupId,
			CreatedAt:    createdAt,
			Content:      content,
			Image:        image,
			NumLikes:     numLikes,
			NumOfComment: len(databse.GetComments(postId)),
		}
		posts = append([]Post{post}, posts...)
	}

	return posts
}
