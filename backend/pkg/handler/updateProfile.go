package handler

import (
	"log"
	"net/http"

	"backend/pkg/auth"
)

// Update is a handler where a users profile can be updated
func (DB *Env) UpdateImage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/updateprofileimg" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	w.Header().Add("Content-Type", "application/text")
	if r.Method == "POST" { // Read the form data and save the image if the user is has a valid session save the image to db else return the string for registration
		successful, reason := auth.WriteImage("images/profile/", r)
		w.Write([]byte(reason))
		if !successful {
			return
		}
		cookie, err := r.Cookie("session_token") // If valid session update the db
		if err != nil {
			log.Println("No current session")
			return
		}
		c, slcErr := auth.SliceCookie(cookie.Value)
		if slcErr != nil {
			log.Println("No session found")
			return
		}

		valid, seshErr := auth.CheckSession(c[2], c[0], *DB.Env)

		if seshErr != nil || !valid {
			log.Println("Session not valid")
			return
		}

		updateErr := auth.Update("User", "avatar", reason, "userId", c[0], *DB.Env)
		if updateErr != nil {
			log.Println("Error updating profile image in db: ", updateErr)
		}

	}
}
