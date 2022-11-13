package response

import (
	"backend/pkg/structs"
	"encoding/json"
	"log"
	"net/http"
)

// Write message logs a value and writes another string to the response writer w
func WriteMessage(logs, write string, w http.ResponseWriter) {
	log.Println(logs)
	WriteErrStruct(write, w)
}

//WriteErrStruct marshalls an error into a struct and writes the message
func WriteErrStruct(val string, w http.ResponseWriter) {
	result := structs.ErrorResponse{Message: val}
	byteSlc, err := json.Marshal(result)
	if err != nil {
		log.Println("Error marshalling error response: ", err)
	}
	w.Write(byteSlc)
}
