package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"backend/pkg/group"
	"backend/pkg/structs"
)

func (database *Env) Group(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		groups, err := group.AllGroups("fdsfdsff", database.Env)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		marshallPosts, err := json.Marshal(groups)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(marshallPosts))
		return
	}

	if r.Method == "POST" {
		var groupData structs.Group
		GetBody(&groupData, w, r)
		_, groupErr := group.CreateGroup(groupData.Name, groupData.Description, groupData.Admin, database.Env)
		if groupErr != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "successfully posted")
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
