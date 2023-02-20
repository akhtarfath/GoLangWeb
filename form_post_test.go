package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	// r.PostForm is a map of our POST form values
	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	_, _ = fmt.Fprintf(w, "%s %s", firstName, lastName) // Fprintf formats according to a format specifier and writes to w.
}

func FormPostWithoutParseForm(w http.ResponseWriter, r *http.Request) {
	// r.FormValue returns the first value for the named component of the query.
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	_, _ = fmt.Fprintf(w, "%s %s", firstName, lastName) // Fprintf formats according to a format specifier and writes to w.
}

func TestFormPost(t *testing.T) {
	// strings.NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString but more efficient and read-only.
	requestBody := strings.NewReader("first_name=Muhammad&last_name=Fathan")
	// requestBody is io.Reader (strings.Reader)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8888/form-post", requestBody)
	// Add adds the key, value pair to the header.x-www-form-urlencoded is the default content type for HTML forms.
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// NewRecorder returns an initialized ResponseRecorder.
	recorder := httptest.NewRecorder()

	FormPost(recorder, request) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.
	//FormPostWithoutParseForm(recorder, request) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.

	response := recorder.Result()        // Result returns the Result of the ResponseRecorder.
	body, _ := io.ReadAll(response.Body) // ReadAll reads from r until an error or EOF and returns the data it read.

	fmt.Println("Response Status:", response.Status)                     // Status returns the HTTP status code of the result.
	fmt.Println("Response Headers:", request.Header.Get("Content-Type")) // Get gets the first value associated with the given key.
	fmt.Println("Response Body:", string(body))                          // string returns a string containing the contents of the slice.
}
