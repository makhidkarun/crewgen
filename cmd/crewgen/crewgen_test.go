// crewgen_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHandleRecruit(t *testing.T) {
	templateDir = "../../web"
	mux := http.NewServeMux()
	mux.HandleFunc("/recruit", recruitCrew)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/recruit", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
