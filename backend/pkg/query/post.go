package query

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type DB struct {
	*sql.DB
}

// CreatePost
// is a method of database that add post in it.
func (databse *DB) CreatePost(userID, groupId, image, content string) (string, error) {
	time := time.Now().Format("2006 January 02 3:4:5 pm")
	postID := uuid.NewV4()
	stmt, _ := databse.DB.Prepare(`
		INSERT INTO Post (postId, userId, groupId, content, image, numLikes, createdAt ) values (?, ?, ?, ?, ?, ?, ?)
	`)
	_, err := stmt.Exec(postID, userID, groupId, content, image, 0 ,time)
	if err != nil {
		return "", err
	}
	return postID.String(), nil
}
