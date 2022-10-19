package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBody(b interface{}, w http.ResponseWriter,  r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&b) // unmarshall the userdata
	if err != nil {
		fmt.Print(err)
		http.Error(w, "500 Internal Server Error.", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}