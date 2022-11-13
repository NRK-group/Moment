package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/pkg/structs"
)

type Env struct {
	Env *structs.DB
}

// GetBody unmarshalls the body of a request into a struct, b must be a struct
func GetBody(b interface{}, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&b) // unmarshall the userdata
	if err != nil {
		fmt.Println("Error unmarshalling req", err)
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}

// Home is the handler for the documentation
func (database *Env) Home(w http.ResponseWriter, r *http.Request) {
	// 'curl -v localhost:5070' run this command to see the response
	// or open the browser and go to localhost:5070
	if r.URL.Path != "/" {
		// if the path is not the expected one, then return a 404 error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}
	if r.Method == http.MethodGet {
		// if the request method is GET, then return the documentation
		w.WriteHeader(http.StatusOK)     // 200
		w.Write([]byte("Documentation")) // send here the api data
		return
	}
	// if the request method is not GET, then return a 405 error
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Method Not Allowed"))
}
