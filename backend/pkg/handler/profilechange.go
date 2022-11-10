package handler

import (
	"fmt"
	"net/http"

	"backend/pkg/auth"
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
		if err != nil { // No cookie found
			response.WriteMessage("No cookie found: ", "Unauthorised", w)
			return
		}
		if !auth.ValidateCookie(c, DB.Env, w) { // Cookie doesn't have a valid session ID
			response.WriteMessage("Invalid Session", "Unauthorised", w)
			return
		}
		var data structs.User // Get the values from the request
		getErr := GetBody(&data, w, r)
		if getErr != nil {
			response.WriteMessage("Could not get request body", "Unauthorised", w)
			return
		}
		fmt.Println()
		fmt.Println("-=-=-+-+-+-+-+GOT here")
		fmt.Println()
		resp, valid := auth.ValidateValues(data.FirstName, data.LastName, data.Email, "NotChecking1", data.IsPublic)
		if !valid {
			response.WriteMessage("Update data not valid", resp, w)
			return
		}
		cookieSlc, cookErr := auth.SliceCookie(c.Value)
		if cookErr != nil {
			response.WriteMessage("Error slicing cookie", "Unauthorised", w)
			return
		}
		updateErr := auth.UpdateUserProfile(cookieSlc[0], data, *DB.Env) // Update the values in the db
		if updateErr != nil {
			response.WriteMessage("Error updating the user profile", "Unauthorised", w)
			return
	}
		response.WriteMessage("Update user profile func run", "Updated", w)
	}
}
