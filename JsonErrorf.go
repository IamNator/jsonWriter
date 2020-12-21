package jsonWriter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Writes ErrorMessage to ResponseWriter in json format
//
//Errorf( ResponseWriter interface, custom error message, http.StatusCode, err error)
//
//Example of json written is --> {"error": "(ErrorMessage : error)" }
func Errorf(w http.ResponseWriter, ErrorMessage string, ErrorCode int, error ...error) {
	w.WriteHeader(ErrorCode)
	w.Header().Set("Content-Type", "application/json")

	ErrorMessage = fmt.Sprintf(ErrorMessage, error) //BUG(IamNator) error format not really cool
	//contains := 0
	for _,v := range error {
		//if ErrorMessage does not contain error message
		if !strings.Contains(ErrorMessage, v.Error()) {
			ErrorMessage = fmt.Sprintf(ErrorMessage + " : %s ", v.Error() )
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

