package handler

import (
	"backend/pkg/auth"
	"net/http"
	"time"
)

// Logout is a handler that runs all functions to logout the user
func (DB *Env) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		c, err := r.Cookie("session_token") // Access the cookie
		if err == nil {                     // Cookie is present so remove
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()})
		} else { // The user isnt logged in
			http.Error(w, "401 unauthorized", http.StatusUnauthorized)
			return
		}
		emailSlc, slcErr := auth.SliceCookie(c.Value)
		if slcErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = auth.UpdateSessionId(emailSlc[1], "-", *DB.Env)// Update the sessionId update in users table and remove from sessions table
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		return
	}
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}