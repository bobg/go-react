package main

// This is a simple HTTP server that serves files from the current directory.
// It listens on localhost:3001 by default.
// Requesting "/foo" serves file foo. Requesting "/" serves "index.html".

import (
	"flag"
	"net/http"
	"strings"
)

func main() {
	addr := flag.String("addr", ":3001", "listen address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := strings.TrimLeft(req.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}
		http.ServeFile(w, req, path)
	})

	http.ListenAndServe(*addr, nil)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
