package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/follow"
	"backend/pkg/response"
	"backend/pkg/structs"
)

func (DB *Env) ProfileChange(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/updateprofileinfo" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "PUT" {
		c, err := r.Cookie("session_token")
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) { // No cookie found
			response.WriteMessage("No cookie found: ", "Unauthorised", w)
			return
		}
		var data structs.User
		if getErr := GetBody(&data, w, r); getErr != nil {
			response.WriteMessage("Could not get request body", "Unauthorised", w)
			return
		}
		if resp, valid := auth.ValidateValues(data.FirstName, data.LastName, data.Email, "NotChecking1", data.IsPublic); !valid {
			response.WriteMessage("Update data not valid", resp, w)
			return
		}
		cookieSlc, cookErr := auth.SliceCookie(c.Value)
		if cookErr != nil {
			response.WriteMessage("Error slicing cookie", "Unauthorised", w)
			return
		}
		if auth.ActiveEmail(cookieSlc[0], data.Email, *DB.Env) {
			response.WriteMessage("Error updating the user profile", "Email already in use", w)
			return
		}
		if updateErr := auth.UpdateUserProfile(cookieSlc[0], data, *DB.Env); updateErr != nil { // Update the values in the db
			response.WriteMessage("Error updating the user profile", "Unauthorised", w)
			return
		}
		if data.IsPublic == 1 {
			if deleteErr := follow.DeletePendingRequests(cookieSlc[0], *DB.Env); deleteErr != nil {
				response.WriteMessage("Error Deleting all pending follow requests", "Error", w)
				return
			}
		}
		response.WriteMessage("Update user profile func run", "Updated", w)
	}
}
