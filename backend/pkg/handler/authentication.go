package handler

import "net/http"

type User struct {
	UserId string `json:"UserId"`
	SessionId string `json:"SessionId"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	NickName string `json:"NickName"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	DateOfBirth string `json:"DateOfBirth"`
	AboutMe string `json:"AboutMe"`
	Avatar string `json:"Avatar"`
	CreatedAt string `json:"CreatedAt"`
	IsLoggedIn string `json:"IsLoggedIn"`
	IsPublic string `json:"IsPublic"`
	NumFollowers string `json:"NumFollowers"`
	NumFollowing string `json:"NumFollowing"`
	NumPosts string `json:"NumPosts"`
}

func (DB *DB) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	//Check if registration is correct
}