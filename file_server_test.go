package main

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")     // Dir is an implementation of FileSystem using the native fileServer system restricted to a specific directory tree.
	fileServer := http.FileServer(directory) // FileServer returns a handler that serves HTTP requests with the contents of the fileServer system rooted at root.

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // Handle registers the handler for the given pattern.
	// http.StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the request URLs Path and invoking the handler h.

	server := http.Server{ // Server is an HTTP server. It wraps a http.Handler to provide HTTP-specific functionality.
		Addr:    "localhost:8889",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

//go:embed resources
var resources embed.FS // FS is an interface that abstracts the access to a collection of named files.

func TestServerWithGolangEmbed(t *testing.T) {
	fileServer := http.FileServer(http.FS(resources)) // http.FS returns a http.FileSystem implementation for the given embed.FS.
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // http.FS returns a http.FileSystem implementation for the given embed.FS.
	// resources is a variable of type embed.FS. It is a read-only file system.

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func TestServerWithGolangEmbedSubDir(t *testing.T) {
	// fs.Sub is a helper function that returns a sub-filesystem of fs, rooted at dir.
	directory, _ := fs.Sub(resources, "resources") // Sub returns a sub-filesystem of fs, rooted at dir.
	// http.FileServer returns a handler that serves HTTP requests with the contents of the fileServer system rooted at root.
	fileServer := http.FileServer(http.FS(directory)) // http.FS returns a http.FileSystem implementation for the given embed.FS.
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // http.FS returns a http.FileSystem implementation for the given embed.FS.
	// resources is a variable of type embed.FS. It is a read-only file system.

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
