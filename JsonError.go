package jsWrite

import (
	"encoding/json"
	"net/http"
)

//Writes Error Message to ResponseWriter in the form a json
//Similar to http.Error(w ResponseWriter, error string, code int)
func JsonError(w http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	w.WriteHeader(ErrorCode)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
}

