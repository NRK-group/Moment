package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/event"
	"backend/pkg/group"
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

	cookie, _ := auth.SliceCookie(c.Value)

	if r.Method == "GET" {

		var retEvents []structs.Event
		groups, err := group.AllUserGroups(cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error. AllUserGroups", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, group := range groups {
			events, err := event.AllEventByGroup(group.GroupID, database.Env)
			if err != nil {
				http.Error(w, "500 Internal Server Error. AllEventByGroup", http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			for _, eventl := range events {
				EventP, _ := event.AllEventParticipant(eventl.EventId, database.Env)
				for i, user := range EventP {
					if user.UserId == cookie[0] {
						events[i].Status = "Going"
					}
				}
			}
			retEvents = events
		}
		marshallEvents, err := json.Marshal(retEvents)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshallEvents)
		return
	}

	if r.Method == "POST" {
		var eventS structs.Event
		err := GetBody(&eventS, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		eventId, err := event.AddEvent(eventS.GroupId, eventS, database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error with getting the events", http.StatusInternalServerError)
			return
		}

		_, errE := event.AddEventParticipant(eventId, eventS.UserId, database.Env)

		if errE != nil {
			http.Error(w, "500 Internal Server Error with AddEventParticipant ", http.StatusInternalServerError)
			return
		}

		marshallEvent, err := json.Marshal(eventId)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallEvent)
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
