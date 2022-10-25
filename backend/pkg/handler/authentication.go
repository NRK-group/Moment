package handler

import (
	"backend/pkg/auth"
	"backend/pkg/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

// Registration is a handler where all registration functions are done
func (DB *Env) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if registration is correct
	if r.Method == "POST" {
		var newUser structs.User
		// Get the body of the request
		GetBody(&newUser, w, r)
		// Insert the new user into the database
		err := auth.InsertUser(newUser, *DB.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}




// GetBody marshalls the body of a request into a struct
func GetBody(b interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&b) // unmarshall the userdata
	if err != nil {
		fmt.Print(err)
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}



