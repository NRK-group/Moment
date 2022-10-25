package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"backend/pkg/structs"
	"backend/pkg/post"
)



func (database *Env) Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		posts, err := post.      AllPost("6t78t8t87", database.Env)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		marshallPosts, err := json.Marshal(posts)
		if err != nil {
			fmt.Println("Error marshalling the data: ", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPosts)
		return
	}
	if r.Method == "POST" {
		var postData structs.Post
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			fmt.Print(err)
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, postErr := post.CreatePost(postData.UserID, postData.Content, postData.GroupID, postData.Image, database.Env)
		fmt.Println("----postData")
		if postErr != nil {
			fmt.Print("postErr - ", postErr)
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
