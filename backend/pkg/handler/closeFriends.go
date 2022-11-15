package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/response"
	"backend/pkg/structs"
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
		var closeFriend structs.CloseFriend
		if bodyErr := GetBody(&closeFriend, w, r); bodyErr != nil { // Get the body of the request
			response.WriteMessage("Error getting close friend body", "Unauthorised", w)
			return
		} 
		// Add closefriend to the database
		resp := closefriend.UpdateCloseFriend(closeFriend.UserId, closeFriend.CloseFriendId, DB.Env)
		response.WriteMessage("Close Friend Updated", resp, w)
	}
}
