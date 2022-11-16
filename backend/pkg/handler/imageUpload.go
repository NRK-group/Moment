package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/pkg/auth"
	commet "backend/pkg/commets"
)

func (database *Env) ImageUpload(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/imageUpload" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}
	SetupCorsResponse(w)
	if r.Method == http.MethodPost {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println("No current session")
			return
		}
		c, slcErr := auth.SliceCookie(cookie.Value)
		if slcErr != nil {
			log.Println("No session found")
			return
		}

		valid, seshErr := auth.CheckSession(c[2], c[0], *database.Env)

		if seshErr != nil || !valid {
			log.Println("Session not valid")
			return
		}

		table := r.FormValue("table")
		id := r.FormValue("id")
		idType := r.FormValue("idType")
		postId := r.FormValue("postid")

		successful, reason := auth.WriteImage("images/"+table+"/", r)

		if successful {
			updateErr := auth.Update(table, "imageUpload", reason, idType, id, *database.Env)
			if updateErr != nil {
				log.Println("Error updating profile image in db: ", updateErr)
			}

			commets, _ := commet.GetComments(postId, database.Env)
			arrComment, _ := json.Marshal(commets)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(arrComment))
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("error uploading file"))
			return
		}
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Method Not Allowed"))
}
