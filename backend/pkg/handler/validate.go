package handler

import (
	"io"
	"log"
	"net/http"

	"backend/pkg/auth"
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
		if err != nil || !auth.ValidateCookie(c, database.Env, w){
			log.Println("No cookie found in validate")
			io.WriteString(w, "Unauthorized")
			return
		}
		io.WriteString(w, "Validated")
	}
}
