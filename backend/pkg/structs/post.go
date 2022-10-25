package structs


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

