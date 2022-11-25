package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/event"
	"backend/pkg/member"
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
		groupId := r.URL.Query().Get("groupId")

		events, err := event.AllEventByGroup(groupId, cookie[0], database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error. AllEventByGroup", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for i, eventS := range events {
			fmt.Println()
			fmt.Println()
			fmt.Println("Events === ", eventS)
			fmt.Println()
			fmt.Println()
			userp, _ := event.GetEventParticipant(eventS.EventId, cookie[0], database.Env)
			if userp.Status == 1 {
				events[i].Status = "Going"
			} else {
				events[i].Status = "Not Going"
			}
		}

		marshallEvents, err := json.Marshal(events)
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

		members, errm := member.GetMembers(eventS.GroupId, database.Env)
		if errm != nil {
			fmt.Println("inside Create Addevent", err)
			return
		}

		for _, member := range members {
			_, errE := event.AddEventParticipant(eventId, member.UserId, database.Env)
			if errE != nil {
				http.Error(w, "500 Internal Server Error with AddEventParticipant ", http.StatusInternalServerError)
				return
			}

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
