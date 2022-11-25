package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/event"
	"backend/pkg/structs"
)

func (database *Env) UpdateEventParticipant(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/updateEventParticipant" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
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

	cookie, _ := auth.SliceCookie(c.Value)

	if r.Method == "POST" {

		var eventS structs.Event
		err := GetBody(&eventS, w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		resEvent, err := event.UpdateEventParticipant(eventS, cookie[0], *database.Env)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		marshallEvent, err := json.Marshal(resEvent)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallEvent)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Method Not Allowed"))
}
