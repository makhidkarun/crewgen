// apiServer/server.go

package main

import (
	"net/http"
)

func newMux(baseRoute string) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)
	return m
}

func replyTextContent(w http.ResponseWriter, r *http.Request,
	status int, content string) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content))
}
