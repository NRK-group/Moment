package handler

import (
	"io"
	"net/http"
)

func (database *Env) Comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("comment get"))
		return
	}

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "successfully add event")
		return
	}
	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
}
