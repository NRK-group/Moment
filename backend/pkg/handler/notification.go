package handler

import (
	"fmt"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/response"
)

// Notification is a notifiation handler
//
// Param:
//
//	w: The response writer
//	r: The request
func (DB *Env) Notification(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notification" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == http.MethodGet {
		c, err := r.Cookie("session_token") // Check if a cookie is present
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		notifType := r.URL.Query().Get("notifType")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Print("-------------------", notifType)
		w.Write([]byte(notifType))
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}
