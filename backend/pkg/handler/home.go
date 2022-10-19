package handler

import (
	"database/sql"
	"net/http"
)

type DB struct {
	DB *sql.DB
}

// Home is the handler for the documentation
func (database *DB) Home(w http.ResponseWriter, r *http.Request) {
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
