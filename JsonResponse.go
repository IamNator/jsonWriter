package jsonWrite

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, Title string, Message string, ErrorCode int) {
	(w).WriteHeader(ErrorCode)
	jmsg := []byte("{ "+Title+":"+Message+ "}")
	json.NewEncoder(w).Encode(jmsg)
}
