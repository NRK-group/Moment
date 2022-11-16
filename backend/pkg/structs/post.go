package structs

type Post struct {
	PostID       string `json:"PostID"`
	UserID       string `json:"UserID"`
	CreatedAt    string
	NickName     string
	ImageUpload  string
	Content      string `json:"Content"`
	GroupID      string `json:"GroupID"`
	Image        string `json:"Image"`
	NumOfComment int
	NumLikes     int
}
