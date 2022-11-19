package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/group"
)

func (database *Env) GetGroupPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getGroupPost" {
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

		groupId := r.URL.Query().Get("groupId")
		posts, err := group.AllGroupPosts(groupId, database.Env)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		marshallPosts, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(marshallPosts))
		return
	}

	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
