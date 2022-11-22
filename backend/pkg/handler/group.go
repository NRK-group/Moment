package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/group"
	"backend/pkg/member"
	"backend/pkg/structs"
)

func (database *Env) Group(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group" {
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

		groups, err := group.AllGroups(cookie[0], database.Env)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(groups) == 0 {
			groups = append([]structs.Group{{Name: "Threre are no groups", Description: "", Admin: "none", Member: true}}, groups...)
		}

		for _, group := range groups {
			members, err := member.GetMembers(group.GroupID, database.Env)
			if err != nil {
				fmt.Print(err)
				http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			for i, member := range members {
				if member.UserId == cookie[0] {
					groups[i].Member = true
				}
			}

		}

		marshallGroups, err := json.Marshal(groups)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(marshallGroups))
		return
	}

	if r.Method == "POST" {
		var groupData structs.Group
		GetBody(&groupData, w, r)
		_, groupErr := group.CreateGroup(groupData.Name, groupData.Description, cookie[0], database.Env)
		if groupErr != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		io.WriteString(w, "successfully in creating a group")
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
