package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/response"
	"backend/pkg/structs"
)

//FollowReq will change the follow relationship between two users
func (DB *Env) FollowReq(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/followrequest" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "PUT" {
		// Get the cookie and authorise the session
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("Invalid cook/session", "Unauthorised", w)
			return
		}
		var followRequest structs.Follower
		if bodyErr := GetBody(&followRequest, w, r); bodyErr != nil {
			response.WriteMessage("Error marshalling the body", "Unauthorised", w)
			return
		}
		if result, checkErr := follow.FollowUser(followRequest.FollowerId, followRequest.FollowingId, DB.Env); checkErr != nil {
			response.WriteMessage("Error running the auth request", "500 internal server error", w)
			return
		} else {
			response.WriteMessage("Error marshalling the body", result, w)
			return
		}
	}
}
