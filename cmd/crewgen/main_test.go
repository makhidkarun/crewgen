// crewgen/main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	templateDir = "web"
}

func TestHandleDocroot(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", crewGen)

	writer = httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
