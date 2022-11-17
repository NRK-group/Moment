package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/messages"
	"backend/pkg/response"
	"backend/pkg/structs"
)

// Profile handles all requests for a users own profile information.
func (DB *Env) Message(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/message" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {
		c, err := r.Cookie("session_token") // Check if a cookie is present
		if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
			response.WriteMessage("No cookie present user unauthorized", "Unauthorised", w)
			return
		}
		chatId := r.URL.Query().Get("chatId") // Get the parameter
		messageType := r.URL.Query().Get("type")
		var msgs []structs.Message
		if messageType == "privateMessage" {
			msgs, err = messages.GetPrivateMessages(chatId, *DB.Env)
		}
		if messageType == "groupMessage" {
			msgs, err = messages.GetGroupMessages(chatId, *DB.Env)
		}
		if err != nil {
			response.WriteMessage("Error getting messages", "Error", w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(msgs)
		if err != nil {
			response.WriteMessage("Error marshalling messages", "Error", w)
			return
		}
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}
