package jsonWriter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func (w http.ResponseWriter, _ * http.Request){
		Response(w, "error", "error", http.StatusOK)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("Status Code Test Passed : %v", status)
	}

	expected := `{"title":"error","message":"error"}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Json Body test Passed : %v", rr.Body)
	}
}

//func BenchmarkJsonError(b *testing.B) {
//
//}
