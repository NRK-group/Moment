package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/response"
	"backend/pkg/structs"
)

// CloseFriends adds/removes a user from the current users close friend list
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
		resp := closefriend.UpdateCloseFriend(closeFriend.UserId, closeFriend.CloseFriendId, *DB.Env)
		response.WriteMessage("Close Friend Updated", resp, w)
	}
}

func (DB *Env) CloseFriendList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getclosefriend" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == http.MethodGet {
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("Cookie not found", "Unauthorised", w)
			return
		}
		cookieSlc, slcErr := auth.SliceCookie(c.Value)
		if slcErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		closeFriendsList := closefriend.GetCloseFriends(cookieSlc[0], *DB.Env) // get the close friends of the current user
		result, marshErr := json.Marshal(closeFriendsList)
		if marshErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}

// func (DB *Env) CheckCloseFriend(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/checkclosefriend" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}
// 	SetupCorsResponse(w)
// 	if r.Method == http.MethodGet {
// 		c, err := r.Cookie("session_token")
// 		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
// 			response.WriteMessage("Cookie not found", "Unauthorised", w)
// 			return
// 		}
// 		cookieSlc, slcErr := auth.SliceCookie(c.Value)
// 		if slcErr != nil {
// 			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}
// 		//Check if the current user is a Close friend of the Profile viewing
// 		profileID := r.URL.Query().Get("profileID") // Get the parameter


// 	}
// }

