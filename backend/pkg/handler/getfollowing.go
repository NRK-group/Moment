package handler

import (
	"backend/pkg/auth"
	"backend/pkg/following"
	"backend/pkg/response"
	"encoding/json"
	"net/http"
)

// GetFollowing is a handler which accesses all the users that follow a user
func (DB *Env) GetFollowing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getfollowing" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == http.MethodGet {
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("Cookie not found", "Unauthorised", w)
			return
		}
		// Get the followers of the user in the cookie
		cookie, slcErr := auth.SliceCookie(c.Value)
		if slcErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		following, getErr := following.Get(cookie[0], *DB.Env)
		if getErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		result, resErr := json.Marshal(following)
		if resErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}
