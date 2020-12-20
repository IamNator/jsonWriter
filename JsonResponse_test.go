package jsonWrite

import (
	//jsonWrite "github.com/IamNator/JsonWrite"

	//"github.com/IamNator/mysql-golang-web/session"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func (w http.ResponseWriter, _ * http.Request){
		JsonResponse(w, "error", "error", http.StatusOK)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{title:"error", message:"error"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
//
//func BenchmarkJsonError(b *testing.B) {
//
//}