package jsonWrite

import (
	"fmt"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, Title string, Message string, ErrorCode int) {
	(w).WriteHeader(ErrorCode)
	//jmsg := "{ \""+Title+"\""+":"+""""+Message+ "\"}"
	jmsg := fmt.Sprintf("{%v:%v}",Title, Message)
	//json.NewEncoder(w).Encode(struct {
	//	eYror string  Title
	//}{
	//	eYror: Message,
	//})

	//json.NewEncoder(w).Encode(jmsg)
}
