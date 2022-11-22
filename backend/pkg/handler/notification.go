package handler

import (
	"encoding/json"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/event"
	"backend/pkg/follow"
	"backend/pkg/member"
	"backend/pkg/response"
	"backend/pkg/structs"
)

// Notification is a notifiation handler
//
// Param:
//
//	w: The response writer
//	r: The request
func (DB *Env) Notification(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notification" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	c, err := r.Cookie("session_token") // Check if a cookie is present
	if err != nil || !auth.ValidateCookie(c, DB.Env, w) {
		response.WriteMessage("No cookie present user unauthorized", "Unauthorised", w)
		return
	}
	cookie, _ := auth.SliceCookie(c.Value)
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		notifType := r.URL.Query().Get("notifType")
		if notifType == "follow" {
			var notif []structs.FollowerNotif
			notif, err = follow.GetFollowNotifs(cookie[0], DB.Env)
			if err != nil {
				response.WriteMessage("Error getting follow notifs", "Error", w)
				return
			}
			err = json.NewEncoder(w).Encode(notif)
			if err != nil {
				response.WriteMessage("Error encoding notif", "Error", w)
				return
			}
		} else if notifType == "group" {
			var notif []structs.GroupNotifWriter
			notif, err = member.GetInvitationNotif(cookie[0], DB.Env)
			if err != nil {
				response.WriteMessage("Error getting notifs", "Error", w)
			}
			event, err := event.GetEventNotifications(cookie[0], DB.Env)
			if err != nil {
				response.WriteMessage("Error getting event notifs", "Error", w)
			}
			notif = append(notif, event...)
			err = json.NewEncoder(w).Encode(notif)
			if err != nil {
				response.WriteMessage("Error encoding group notif", "Error", w)
				return
			}
		} else {
			err = json.NewEncoder(w).Encode(notifType)
			if err != nil {
				response.WriteMessage("Error encoding notif", "Error", w)
				return
			}
		}
		return
	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}
