package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Post struct {
	PostID       string
	UserID       string `json:"UserID"`
	CreatedAt    string
	Content      string `json:"Content"`
	GroupID      string `json:"GroupID"`
	Image        string `json:"Image"`
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

func (database *DB) Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		posts, err := database.AllPost("6t78t8t87")
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		marshallPosts, err := json.Marshal(posts)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPosts)
		return
	}
	if r.Method == "POST" {
		var postData Post
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, postErr := database.CreatePost(postData.UserID, postData.Content, postData.GroupID, postData.Image)
		fmt.Println("----postData")
		if postErr != nil {
			fmt.Print("postErr - ", postErr)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "successfully posted")
		return
	}

	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}

// ----------------------- Post Queries -----------------------

// Get Comments
// is a method of forum that return all the comment with that specific postID
func (database *DB) GetComments(pID string) []Comment {
	rows, err := database.DB.Query("SELECT * FROM Comment WHERE postId = '" + pID + "'")
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
func (database *DB) AllPost(uID string) ([]Post, error) {
	var post Post
	var posts []Post
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
		post = Post{
			PostID:       postId,
			UserID:       userId,
			GroupID:      groupId,
			CreatedAt:    createdAt,
			Content:      content,
			Image:        image,
			NumLikes:     numLikes,
			NumOfComment: len(database.GetComments(postId)),
		}
		posts = append([]Post{post}, posts...)
	}

	return posts, nil
}

// CreatePost
// is a method of database that add post in it.
func (database *DB) CreatePost(userID, groupId, image, content string) (string, error) {
	createdAt := time.Now().Format("2006 January 02 3:4:5 pm")
	postID := uuid.NewV4()
	stmt, _ := database.DB.Prepare(`
		INSERT INTO Post (postId, userId, groupId, content, image, numLikes, createdAt ) values (?, ?, ?, ?, ?, ?, ?)
	`)
	_, err := stmt.Exec(postID, userID, groupId, content, image, 0, createdAt)
	if err != nil {
		fmt.Println("inside Create Post", err)
		return "", err
	}
	return postID.String(), nil
}
