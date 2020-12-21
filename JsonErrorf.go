package jsonWriter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Writes Error Message to ResponseWriter in json format
//Similar Errorf( ResponseWriter interface, custom error message, http.StatusCode, err error)
//Example of json written is --> {"error":ErrorMessage}
func Errorf(w http.ResponseWriter, ErrorMessage string, ErrorCode int, error ...error) {
	w.WriteHeader(ErrorCode)
	w.Header().Set("Content-Type", "application/json")

	ErrorMessage = fmt.Sprintf(ErrorMessage, error)
	//contains := 0
	for i,_ := range error {
		if !strings.Contains(ErrorMessage, error[i].Error())/*if ErrorMessage does not contain error */ {
			ErrorMessage = fmt.Sprintf(ErrorMessage + " : %v ", error[i] )
		}
	}


	err := json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to encode json response"))
	}
}

