package handler

import "net/http"

func (DB *Env) Update(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/updateprofile" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
}
