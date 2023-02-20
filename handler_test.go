package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServeWithHandler(t *testing.T) {
	// HandlerFunc is an adapter to allow the use of ordinary functions as HTTP handlers.
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, r.Method) // Fprintln formats using the default formats for its operands and writes to w.
		_, _ = fmt.Fprintln(w, r.URL)    // Fprintln formats using the default formats for its operands and writes to w.
	}

	server := http.Server{ // Server is an HTTP server. It wraps an HTTP listener and optionally TLS.
		Addr:    "localhost:8888",
		Handler: handler,
	}

	// ListenAndServe listens on the TCP network address and
	//then calls Serve with handler to handle requests on incoming connections.
	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Server is running:", server.Addr)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// HandlerFunc is an adapter to allow the use of ordinary functions as HTTP handlers.
	// return json response with status code 200
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintln(w, `{"message": "Hello Handler!"}`)

	//_, _ = fmt.Fprintln(w, "Hello Handler!") // Fprintln formats using the default formats for its operands and writes to w.
}

func TestHelloHandler(t *testing.T) {
	// NewRequest returns a new Request given a method, URL, and optional body.
	request := httptest.NewRequest("GET", "http://localhost:8888/hello", nil)
	// NewRecorder returns an initialized ResponseRecorder.
	// ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
	recorder := httptest.NewRecorder()

	// ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.
	HelloHandler(recorder, request)

	response := recorder.Result()            // Result returns the Result of the ResponseRecorder.
	body, _ := ioutil.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read.

	fmt.Println(response.StatusCode)                 // StatusCode returns the HTTP status code of the result.
	fmt.Println(response.Header.Get("Content-Type")) // Get gets the first value associated with the given key.
	fmt.Println(string(body))                        // string(body) converts the body to string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	love := r.URL.Query().Get("love")
	// URL is the URL of the request.
	// URL.Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs.
	// Query returns the parsed query parameters. Query parses RawQuery and returns the corresponding values.
	if name == "" && love == "" {
		_, _ = fmt.Fprintln(w, "Hello!")
	} else if name == "" && love != "" {
		_, _ = fmt.Fprintf(w, "Hello, I love %s", love)
	} else if name != "" && love == "" {
		_, _ = fmt.Fprintf(w, "Hello %s", love)
	} else {
		_, _ = fmt.Fprintf(w, "Hello %s, I love %s", name, love)
	}
}

func TestSayHello(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/sayhello?name=Fathan&love=Golang", nil)
	// NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()
	// NewRecorder returns an initialized ResponseRecorder.

	// ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.
	SayHello(recorder, request)
	response := recorder.Result()            // Result returns the Result of the ResponseRecorder.
	body, _ := ioutil.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read.

	fmt.Println(response.StatusCode)                 // StatusCode returns the HTTP status code of the result.
	fmt.Println(response.Header.Get("Content-Type")) // Get gets the first value associated with the given key.
	fmt.Println(string(body))                        // string(body) converts the body to string
}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query() // URL is the URL of the request.
	var names []string = query["name"]   // query["name"] returns a slice of strings.

	_, _ = fmt.Fprintln(w, strings.Join(names, ", ")) // strings.Join concatenates the elements of a to create a single string.
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/multiplequeryparameter?name=Fathan&name=Yustika", nil)
	// NewRequest returns a new Request given a method, URL, and optional body. The provided body is closed by the function.
	recorder := httptest.NewRecorder() // NewRecorder returns an initialized ResponseRecorder.

	// ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.
	MultipleQueryParameter(recorder, request)
	result := recorder.Result()            // Result returns the Result of the ResponseRecorder.
	body, _ := ioutil.ReadAll(result.Body) // ReadAll reads from r until an error or EOF and returns the data it read.

	fmt.Println(result.StatusCode)                 // StatusCode returns the HTTP status code of the result.
	fmt.Println(result.Header.Get("Content-Type")) // Get gets the first value associated with the given key.
	fmt.Println(string(body))                      // string(body) converts the body to string
}
