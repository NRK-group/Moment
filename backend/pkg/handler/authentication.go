package handler

import (
	"backend/pkg/functions"
	"fmt"
	"net/http"
	// "backend/pkg/queries"
)

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
	IsLoggedIn int `json:"IsLoggedIn"`
	IsPublic int `json:"IsPublic"`
	NumFollowers int `json:"NumFollowers"`
	NumFollowing int `json:"NumFollowing"`
	NumPosts int `json:"NumPosts"`
}

func (DB DB) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	//Check if registration is correct
	if r.Method == "POST" {
		var newUser User
		//Get the body of the request
		functions.GetBody(&newUser, w, r)
		//Insert the new user into the database
		
	}
}