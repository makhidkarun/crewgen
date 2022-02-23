// crewgen/main_test.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	//"path/filepath"
	"runtime"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder
var binName = "crewgen"

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	fmt.Printf("Building %s...\n", binName)
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build %s: %s\n", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running crewgen tests...")
	result := m.Run()
	fmt.Println("Cleanng up")
	//os.Remove(binName)
	os.Exit(result)
}

/*
func TestHandleDocroot(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tDir := t.TempDir()
	//cmdPath := filepath.Join(dir, binName)
	fmt.Printf("Running in %s, with %s as dir.\n", tDir, dir)
	t.Run("BaseCrewgenRun", func(t *testing.T) {
		mux = http.NewServeMux()
		mux.HandleFunc("/", crewGen)

		writer = httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)
		mux.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code is %v", writer.Code)
		}
	})
}
*/
