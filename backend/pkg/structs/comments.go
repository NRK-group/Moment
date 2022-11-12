package structs

type Comment struct {
	CommentID string
	CommentName string
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`
	CreatedAt string
	Image     string `json:"image"`
	Content   string `json:"content"`
	NumLikes  int `json:"numLikes"`
}



type RetrunComment struct {
	CommentID string
	UserID    string
	PostID    string
	CreatedAt string
	Image     string
	Name   	  string
	Content   string
	NumLikes  int
}