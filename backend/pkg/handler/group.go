package handler

import "net/http"

func (database *Env) Group(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)    
		w.Write([]byte("Group-Get")) 
		return
	}

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)    
		w.Write([]byte("Group-POST")) 
		return
	}


}
