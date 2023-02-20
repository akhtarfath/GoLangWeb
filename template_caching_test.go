package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ParseFS creates a new Template and parses the template definitions from the files identified by the pattern.
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

// Speed up the template execution by caching the template. because the template is parsed only once.
func MyTemplateCaching(w http.ResponseWriter, r *http.Request) {
	_ = myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	MyTemplateCaching(recorder, request)              // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body) // Result returns the Result of the ResponseRecorder.
	t.Log("Response Body:", string(body))             // Response Body: Hello HTML Template
}
