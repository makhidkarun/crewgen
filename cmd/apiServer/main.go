// apiServer/main.go

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	host := flag.String("h", "localhost", "Server host")
	port := flag.Int("p", 8080, "Server port")
	flag.Parse()

	baseRoute := "myThing"

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      newMux(baseRoute),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
