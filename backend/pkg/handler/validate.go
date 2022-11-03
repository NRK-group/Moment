package handler

import (
	"backend/pkg/auth"
	"io"
	"log"
	"net/http"
)

func (database *Env) Validate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	w.Header().Add("Content-Type", "application/text")
	if r.Method == "GET" {
		c, err := r.Cookie("session_token")
		if err != nil {
			log.Println("No cookie found in validate")
			io.WriteString(w, "Unauthorized")
			return
		}
		cookie, cErr := auth.SliceCookie(c.Value)
		if cErr != nil {
			log.Println("Error slicing the cookie")
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		valid, seshErr := auth.CheckSession(cookie[2], cookie[0], *database.Env)
		if seshErr != nil {
			log.Println("Error searching for session")
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		if !valid {
			auth.RemoveCookie(w)
			io.WriteString(w, "Unauthorized")
			return
		}
		io.WriteString(w, "Validated")
	}
}
