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
	w.Header().Add("Content-Type", "application/text")
	if r.Method == "POST" { // Check if registration is correct
		var newUser *structs.User
		bodyErr := GetBody(&newUser, w, r) // Get the body of the request
		if bodyErr != nil {
			w.Write([]byte("500 Internal Server Error"))
			return
		}
		msg, valid := auth.ValidateValues(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.IsPublic) // Validate the user input here
		if !valid {
			w.Write([]byte(msg))
			return
		}
		auth.Capitalise(newUser)                  // Make all values lowercase
		err := auth.InsertUser(*newUser, *DB.Env) // Insert the new user into the database
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("Successfully Registered"))
	}
}


