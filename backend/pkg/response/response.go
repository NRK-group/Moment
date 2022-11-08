package response

import (
	"log"
	"net/http"
)

// Write message logs a value and writes another string to the response writer w
func WrtieMessage(logs, write string, w http.ResponseWriter) {
	log.Println(logs)
	w.Header().Add("Content-Type", "application/text")
	w.Write([]byte(write))
}
