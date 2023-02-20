package main

import (
	"GolangWeb/entity"
	"fmt"
	"net/http"
	"testing"
)

// Middleware is a function that takes a handler and returns a new handler
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler called")
		_, _ = fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic called")
	})

	logMiddleware := &entity.LogMiddleware{
		Handler: mux, // call mux
	}

	errorHandler := &entity.ErrorHandler{
		Handler: logMiddleware, // call log middleware
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler, // call error handler
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
