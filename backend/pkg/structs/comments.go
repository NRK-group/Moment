package structs

type Comment struct {
	CommentID string
	UserID    string
	PostID    string
	CreatedAt string
	Image     string
	Content   string
	NumLikes  int
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