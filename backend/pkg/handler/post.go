package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/post"
	"backend/pkg/response"
	"backend/pkg/structs"
)

func (database *Env) Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	if r.Method == "GET" {
		posts, err := post.AllPost(database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		marshallPosts, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshallPosts)
		return
	}
	if r.Method == "POST" {
		var postData structs.Post
		GetBody(&postData, w, r)

		postID, postErr := post.CreatePost(postData.UserID, postData.GroupID, "", postData.Content, postData.Privacy, database.Env)
		if postErr != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err := post.IncreasePostNumber(postData.UserID, *database.Env)
		if err != nil{
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("Error updating the post number: ", err)
			return
		} 
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response.WriteMessage("Successfully Posted", postID, w)
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
