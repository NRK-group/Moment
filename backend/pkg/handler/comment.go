package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	commet "backend/pkg/commets"
	"backend/pkg/structs"
)

func (database *Env) Comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	id := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	fmt.Println("Id----", id)

	SetupCorsResponse(w)
	if r.Method == "GET" {

		commets, err := commet.GetComments(id, database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error with getting the comments", http.StatusInternalServerError)
			return
		}

		marshallPage, err := json.Marshal(commets)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPage)
		return
	}

	if r.Method == "POST" {
		var post structs.Post
		err := GetBody(&post, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		commets, err := commet.GetComments(post.PostID, database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error with getting the comments", http.StatusInternalServerError)
			return
		}

		marshallPage, err := json.Marshal(commets)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPage)
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
