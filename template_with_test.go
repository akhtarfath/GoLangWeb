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

func TemplateDataWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/with.gohtml")) // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	_ = t.ExecuteTemplate(w, "with.gohtml", entity.Page{
		Title: "With in Go Template",
		Name: entity.Name{
			First:  "Muhammad",
			Middle: "Fathan",
			Last:   "Aulia",
		},
		Address: entity.Address{
			Street:   "Jl. Tanah Merdeka",
			City:     "Jakarta Timur",
			Province: "DKI Jakarta",
			Country:  "Indonesia",
		},
		Phone: "08978329974",
	})
}

func TestTemplateDataWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateDataWith(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println("Response Body:", string(body))
}
