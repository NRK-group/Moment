package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/response"
)

func (DB *Env) CloseFriends(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/closefriend" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "POST" {
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("Cookie not found", "Unauthorised", w)
			return
		}
		//Add closefriend to the database
	}
}
