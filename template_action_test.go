package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml")) // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	_ = t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Title": "Hello HTML Template with Action",
		"Name":  "",
	})
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateAction(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	fmt.Println("Response Body:", string(body))       // string returns a string containing the contents of the slice.
}
