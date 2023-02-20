package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		http.ServeFile(w, r, "./resources/notfound.html")
	} else {
		http.ServeFile(w, r, "./resources/ok.html")
	}
}

func TestServerFileServe(t *testing.T) {
	server := http.Server{ // Server is an HTTP server. It wraps an http.Server.
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(ServeFile), // HandlerFunc converts a function to an http.Handler.
	}

	err := server.ListenAndServe()
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	if err != nil {
		panic(err.Error())
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		_, _ = fmt.Fprint(w, resourceNotFound)
	} else {
		_, _ = fmt.Fprint(w, resourceOK)
	}
}

func TestServerFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
