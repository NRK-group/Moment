package handler

import (
	"encoding/json"
	"log"
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
			response.WrtieMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		cookie, slcErr := auth.SliceCookie(c.Value)// Valid session so return details for the user
		if slcErr != nil {
			log.Println("No cookie present user unauthorized")
			w.Write([]byte("Unauthorised"))
			return
		}
		var result structs.User
		getErr := auth.GetUser("email", cookie[1], &result, *DB.Env)
		if getErr != nil {
			response.WrtieMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		sendBack, marshErr := json.Marshal(result)
		if marshErr != nil {
			response.WrtieMessage("Error marshalling user profile data", "500 Internal Server Error", w)
			return
		}
		w.Write(sendBack)
	}
}
