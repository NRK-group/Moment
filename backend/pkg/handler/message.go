package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	l "backend/pkg/log"
	"backend/pkg/messages"
	"backend/pkg/response"
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
		// var result structs.Message
		messages, err := messages.GetPrivateMessages(chatId, *DB.Env)
		if err != nil {
			l.LogMessage("message.go", "Error getting messages", err)
			response.WriteMessage("Error getting messages", "Error", w)
			return
		}
		l.LogMessage("Messages.go", "Messages", messages)
		resp, err := json.Marshal(messages)
		if err != nil {
			l.LogMessage("message.go", "Error marshalling messages", err)
			response.WriteMessage("Error marshalling messages", "Error", w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}
