package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/post"
)

func (database *Env) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getUserPosts" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {

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
		userID := r.URL.Query().Get("userID") // Get the parameter


		posts, err := post.AllUserPost(userID, database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		marshallPosts, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPosts)
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
