package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/chat"
)

type Receiver struct {
	Id string `json:"receiverId"`
}

func (database *Env) Chat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(chats)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == http.MethodPost {
		var receiver Receiver
		err := GetBody(&receiver, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		isChat, info := chat.CheckIfChatExists(cookie[0], receiver.Id, database.Env)
		if isChat {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(info)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}
		chatData, err := chat.InsertNewChat(cookie[0], receiver.Id, database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(chatData)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
