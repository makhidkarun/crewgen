// crewgen_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var	mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	templateDir = "web"
}

func TestHandleRecruit(t *testing.T) {
	//templateDir = "../../web"
	mux = http.NewServeMux()
	mux.HandleFunc("/recruit", recruitCrew)

	writer = httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/recruit", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleShow(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/show", recruitCrew)

	writer = httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/show", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
