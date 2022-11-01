package handler

import (
	"fmt"
	"io"
	"net/http"

	"backend/pkg/structs"
)

func (database *Env) Comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	
	SetupCorsResponse(w)
	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("comment get"))
		return
	}
	
	if r.Method == "POST" {
		var post structs.Post
		err := GetBody(&post, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println(post)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "successfully add event")
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
