package main

import (
	"GolangWeb/entity"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml")) // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	_ = t.ExecuteTemplate(w, "range.gohtml", entity.Hobbies{
		Hobbies: []string{
			"Programming", "Reading", "Writing", "Listening to Music", "Watching Movies", "Playing Games",
		},
	})
}

func TestTemplateDataRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateDataRange(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println("Response Body:", string(body))
}
