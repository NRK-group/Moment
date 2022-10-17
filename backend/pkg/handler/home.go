package handler

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DB struct {
	*sql.DB
}

// Home is the handler for the documentation
func (databse *DB) Home(w http.ResponseWriter, r *http.Request) {
	// 'curl -v localhost:5070' run this command to see the response
	// or open the browser and go to localhost:5070
	w.WriteHeader(http.StatusOK)     // 200
	w.Write([]byte("Documentation")) // send here the api data
	d, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err) // log the error
	}
	fmt.Fprint(w, string(d)) // print the data
}
