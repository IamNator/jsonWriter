package jsonWrite

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, Title string, Message string, ErrorCode int) {
	(w).WriteHeader(ErrorCode)
	//jmsg := "{ \""+Title+"\""+":"+""""+Message+ "\"}"
	//jmsg := fmt.Sprintf("{%v:%v}",Title, Message)
	json.NewEncoder(w).Encode(struct {
		Title string  `json:"title"`
		Message string `json:"message"`
	}{
		Title: Title,
		Message: Message,
	})
}
