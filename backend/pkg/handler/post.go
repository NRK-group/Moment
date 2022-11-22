package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/closefriend"
	"backend/pkg/follow"
	"backend/pkg/post"
	"backend/pkg/response"
	"backend/pkg/structs"
)

func (database *Env) Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil || !auth.ValidateCookie(c, database.Env, w) {
		log.Println("No cookie found in validate")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	SetupCorsResponse(w)
	w.Header().Add("Content-Type", "application/json")
	cookie, _ := auth.SliceCookie(c.Value)
	
	if r.Method == "GET" {
		var returnPost []structs.Post
		posts, err := post.AllPost(database.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, post := range posts {
			if post.UserID == cookie[0] {
				returnPost = append([]structs.Post{post}, returnPost...)
			} else if post.Privacy == 0 {
				returnPost = append([]structs.Post{post}, returnPost...)
			} else if post.Privacy == 1 && follow.CheckIfFollow(cookie[0], post.UserID, database.Env) {
				returnPost = append([]structs.Post{post}, returnPost...)
			} else if post.Privacy == -1 {
				closeFriendsList := closefriend.GetCloseFriends(cookie[0], *database.Env)
				for _, closeF := range closeFriendsList {
					if closeF.Id == cookie[0] {
						returnPost = append([]structs.Post{post}, returnPost...)
					}
				}
			}
		}

		marshallPosts, err := json.Marshal(returnPost)
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
