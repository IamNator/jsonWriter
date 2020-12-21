package jsonWriter

import (
	//jsonWrite "github.com/IamNator/JsonWrite"

	"errors"

	//"github.com/IamNator/mysql-golang-web/session"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorf(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var errs [3]error
	errs[0] = errors.New("test error, ")
	errs[1] = errors.New("test error 2, ")
	errs[2] = errors.New("test error 3")

	handler := http.HandlerFunc(func (w http.ResponseWriter, _ * http.Request){
		Errorf(w, "error : ", http.StatusOK, errs[0], errs[1], errs[2])
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("Status Code Test Passed : %v", status)
	}

	expected := `{"error":"error : %!(EXTRA []error=[test error,  test error 2,  test error 3])"}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Json Body test Passed : %v", rr.Body)
	}
}

func TestErrorf2(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var errs [3]error
	errs[0] = errors.New("test error, ")
	errs[1] = errors.New("test error 2, ")
	errs[2] = errors.New("test error 3")

	handler := http.HandlerFunc(func (w http.ResponseWriter, _ * http.Request){
		Errorf(w, "error : %v", http.StatusOK, errs[0])
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		t.Logf("Status Code Test Passed : %v", status)
	}

	expected := `{"error":"error : [test error, ]"}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Json Body test Passed : %v", rr.Body)
	}
}
//
//func BenchmarkJsonError(b *testing.B) {
//
//}