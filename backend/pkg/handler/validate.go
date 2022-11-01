package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (database *Env) Validate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/validate" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		c, err := r.Cookie("session_token")
		if err != nil {
			log.Println("No cookie found in validate")
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println(c)
	}
}
