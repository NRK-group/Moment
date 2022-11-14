package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/chat"
)

func (database *Env) Chat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
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
		cookie, _ := auth.SliceCookie(c.Value)
		chats, err := chat.GetPreviousPrivateChat(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		groupChat, err := chat.GetPreviousGroupChat(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		chats = append(chats, groupChat...)
		chatlist, err := json.Marshal(chats)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(chatlist)
		return
	}

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		// w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("successfully add chat"))
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
