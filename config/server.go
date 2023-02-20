package config

import (
	"fmt"
	"net/http"
)

func Serve() {
	// NewServeMux look like router in express js
	handler := http.NewServeMux() // NewServeMux allocates and returns a new ServeMux.
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request
	//against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
	Routes(handler)

	server := http.Server{ // Server is an HTTP server. It wraps an HTTP listener and optionally TLS.
		Addr:    "localhost:8888",
		Handler: handler,
	}

	err := server.ListenAndServe()
	// ListenAndServe listens on the TCP network address and
	//then calls Serve with handler to handle requests on incoming connections.
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Server is running:", server.Addr)
	}
}
