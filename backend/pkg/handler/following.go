package handler

import (
	"fmt"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/response"
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
		// Check if cookieSlc[0] is following  followingId
		fmt.Println("UserId === ", cookieSlc[0])
		fmt.Println()
		fmt.Println()
		fmt.Println("FollowingID === ", followingId)
		fmt.Println()
		fmt.Println()

		if follow.CheckIfFollow(cookieSlc[0], followingId, DB.Env) { // User is following the profile
			response.WriteMessage(cookieSlc[0]+" follows "+followingId, "Following", w)
			return
		}
		if follow.CheckIfFollowPending(cookieSlc[0], followingId, DB.Env) {
		response.WriteMessage(cookieSlc[0]+" pending "+followingId, "pending", w)

		}
		response.WriteMessage(cookieSlc[0]+" not following "+followingId, "Not Following", w)
	}
}
