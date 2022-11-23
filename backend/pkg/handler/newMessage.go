package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/users"
)

func (database *Env) NewMessage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/message/new" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	c, err := r.Cookie("session_token")
	if err != nil || !auth.ValidateCookie(c, database.Env, w) {
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method == http.MethodGet {
		cookie, _ := auth.SliceCookie(c.Value)
		following, err := chat.GetFollowingInfo(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		pUsers, err := users.GetAllPublicUsersNotFollowed(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		following = append(following, pUsers...)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(following)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
