package jsonWrite_test

import (
	"net/http"
	"testing"
)

func TestJsonError(t *testing.T) {
	req, err := http.NewRequest("GET","/")

}

func BenchmarkJsonError(b *testing.B) {

}