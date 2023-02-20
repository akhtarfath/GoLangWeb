package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml")) // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	_ = t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 100,
	})
}

func TestTemplateDataOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateDataOperator(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println("Response Body:", string(body))
}
