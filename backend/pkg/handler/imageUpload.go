package handler

import (
	"fmt"
	"log"
	"net/http"

	"backend/pkg/auth"
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

		successful, reason := auth.WriteImage("images/"+table+"/", r)
		fmt.Println("successful - ", successful)
		if successful {
			updateErr := auth.Update(table, "imageUpload", reason, idType, id, *database.Env)
			if updateErr != nil {
				log.Println("Error updating image in db: ", updateErr)
			}
			w.Header().Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Image upload"))
			return
		} else {
			w.Header().Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("error uploading Image"))
			return
		}
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Method Not Allowed"))
}
