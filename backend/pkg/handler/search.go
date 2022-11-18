package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/search"
)

func (database *Env) Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil || !auth.ValidateCookie(c, database.Env, w) {
		log.Println("No cookie found in validate")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	SetupCorsResponse(w)
	w.Header().Add("Content-Type", "application/json")
	cookie, _ := auth.SliceCookie(c.Value)
	if r.Method == "GET" {
		users, err := search.GetAllUsers(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}
