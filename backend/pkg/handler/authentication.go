package handler

import "net/http"

type User struct {
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	NickName string `json:"NickName"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	DateOfBirth string `json:"DateOfBirth"`
	AboutMe string `json:"AboutMe"`
	Avatar string `json:"Avatar"`
}

func (DB *DB) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	//Check if registration is correct
}