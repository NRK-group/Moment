package handler

import (
	"backend/pkg/auth"
	"backend/pkg/response"
	"net/http"
)

func (DB *Env) Following(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/following" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {
		// Validate the user session
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("Cookie not found", "Unauthorised", w)
			return
		}
		
		followingId := r.URL.Query().Get("followingID") // Get the query for the profile being checked 
		cookieSlc, slcErr := auth.SliceCookie(c.Value)
		if slcErr != nil {
			response.WriteMessage("Error slicing cookie", "Unauthorised", w)
			return
		}
		//Check if cookieSlc[0] is following  followingId
		


	}
}
