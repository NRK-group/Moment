package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"backend/pkg/auth"
	commet "backend/pkg/commets"
	"backend/pkg/structs"
)

func (database *Env) Comment(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.URL.Path, "/")[1] != "comment" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}


	SetupCorsResponse(w)

	if r.Method == "GET" {

		id := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

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

		c, err := r.Cookie("session_token")
		if err != nil {
			log.Println("No cookie found in validate")
			io.WriteString(w, "Unauthorized")
			return
		}
		cookie, cErr := auth.SliceCookie(c.Value)
		if cErr != nil {
			fmt.Println("Error slicing the cookie")
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		var commentS structs.Comment
		err = GetBody(&commentS, w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		_, errc := commet.CreateComment(cookie[0] , commentS.PostID, commentS.Content, database.Env)
		if errc != nil {
			fmt.Println("Error inputing a comment in the db %v", err)
		}

		returstr, err := commet.GetComments(commentS.PostID, database.Env)

		if err != nil {
			http.Error(w, "500 Internal Server Error with getting the comments", http.StatusInternalServerError)
			return
		}
		
		marshallPage, err := json.Marshal(returstr)
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
