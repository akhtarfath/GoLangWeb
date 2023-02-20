package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleTemplate(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`    // . is the current object, in this case a string.
	t, err := template.New("SIMPLE").Parse(templateText) // Parse parses text as a template.
	//t := template.Must(template.New("SIMPLE").Parse(templateText)) // Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as
	if err != nil {
		panic(err.Error()) // panic is equivalent to Print() followed by a call to panic().
	}
	_ = t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
	// ExecuteTemplate applies a parsed template to the specified data object, writing the output to wr.
	// The name identifies the template to apply. If the template cannot be found, an error will be returned.
}

func TestSimpleTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	SimpleTemplate(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	fmt.Println("Response Body:", string(body))       // string returns a string containing the contents of the slice.
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml") // ParseFiles creates a new Template and parses the template definitions from the named files.
	if err != nil {
		panic(err.Error())
	}
	_ = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template") // ExecuteTemplate applies a parsed template to the specified data object, writing the output to wr.
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	SimpleHTMLFile(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	fmt.Println("Response Body:", string(body))       // string returns a string containing the contents of the slice.
}
