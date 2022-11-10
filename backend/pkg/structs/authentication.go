package structs

type User struct {
	UserId       string `json:"UserId"`
	SessionId    string `json:"SessionId"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	NickName     string `json:"NickName"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	DateOfBirth  string `json:"DateOfBirth"`
	AboutMe      string `json:"AboutMe"`
	Avatar       string `json:"Avatar"`
	CreatedAt    string `json:"CreatedAt"`
	IsLoggedIn   int    `json:"IsLoggedIn"`
	IsPublic     int    `json:"IsPublic"`
	NumFollowers int    `json:"NumFollowers"`
	NumFollowing int    `json:"NumFollowing"`
	NumPosts     int    `json:"NumPosts"`
}