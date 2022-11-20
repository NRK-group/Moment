package handler

import (
	"io"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/structs"
)

func (database *Env) Event(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/event" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	SetupCorsResponse(w)

	c, err := r.Cookie("session_token")
	if err != nil {
		log.Println("No cookie found in validate")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	if !auth.ValidateCookie(c, database.Env, w) {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("event get"))
		return
	}

	if r.Method == "POST" {
		var eventS structs.Event
		err := GetBody(&eventS, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		// w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "successfully add event")
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
