package handler

import (
	"fmt"
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
	if r.Method != "POST" { // Read the form data and save the image if the user is has a valid session save the image to db else return the string for registration
		successful, reason := auth.WriteImage("../../images/profile", r)
		if !successful {
			w.Write([]byte(reason))
			return
		}
		// reason == img dirrectory
		fmt.Println("REASON ==== ", reason)

	}
}
