package handler

import (
	"net/http"

	"backend/pkg/auth"
	"backend/pkg/structs"
)

func SetupCorsResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8070")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
}

// Registration is a handler where all registration functions are done
func (DB *Env) Registration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	SetupCorsResponse(w)
	// Check if registration is correct
	if r.Method == "POST" {
		var newUser structs.User
		// Get the body of the request
		bodyErr := GetBody(&newUser, w, r)
		if bodyErr != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Validate the user input here
		msg, valid := auth.ValidateValues(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
		if !valid {
			w.Write([]byte(msg))
			return
		}
		// Insert the new user into the database
		err := auth.InsertUser(newUser, *DB.Env)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
