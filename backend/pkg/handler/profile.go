package handler

import "net/http"

func (DB *Env) Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
}
