package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/response"
	"backend/pkg/structs"
)

// Profile handles all requests for a users own profile information.
func (DB *Env) Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {
		c, err := r.Cookie("session_token") // Check if a cookie is present
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		userID := r.URL.Query().Get("userID") // Get the parameter
		var result structs.User
		getErr := auth.GetUser("userId", userID, &result, *DB.Env)
		if getErr != nil {
			response.WriteMessage("Error getting user: "+getErr.Error(), "User Not Found", w)
			return
		}
		result.Password = ""
		profileDetails, marshErr := json.Marshal(result)
		if marshErr != nil {
			response.WriteMessage("Error marshalling user profile data", "500 Internal Server Error", w)
			return
		}
		w.Write(profileDetails)
	}
}
