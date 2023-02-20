package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 100 - 199 Informational responses
// 200 - 299 Successful responses
// 300 - 399 Redirects
// 400 - 499 Client errors
// 500 - 599 Server errors

// http.ResponseWriter is an interface that allows an HTTP handler to construct an HTTP response.
// *http.Request is an interface that describes a request to a server. It is an interface so that it can be easily mocked for testing.
func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(400) // Bad Request
		// Fprintln is similar to Println, but writes to w instead of the standard output.
		_, _ = fmt.Fprintln(w, "name is required")
	} else {
		w.WriteHeader(200) // OK
		_, _ = fmt.Fprintf(w, "Hi, %s", name)
	}
}

// *testing.T is a type passed to Test functions to manage test state and support formatted test logs.
func TestResponseCode(t *testing.T) {
	// example success request
	//request := httptest.NewRequest("GET", "http://localhost:8888/response-code?name=Muhammad%20Fathan", nil)

	// example bad request, name is empty
	request := httptest.NewRequest("GET", "http://localhost:8888/response-code", nil)
	recorder := httptest.NewRecorder() // NewRecorder returns an initialized ResponseRecorder.

	ResponseCode(recorder, request) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.

	response := recorder.Result()        // Result returns the Result of the ResponseRecorder.
	body, _ := io.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read.

	fmt.Println("Response Code:", response.StatusCode) // StatusCode is the HTTP response status code of the response.
	fmt.Println("Response Body:", string(body))        // string(body) converts the body to string
}
