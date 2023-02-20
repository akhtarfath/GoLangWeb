package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml") // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	if err != nil {
		panic(err.Error())
	}

	_ = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateDirectory(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	fmt.Println("Response Body:", string(body))       // string returns a string containing the contents of the slice.
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml") // ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern.
	if err != nil {
		panic(err.Error())
	}

	_ = t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template") // ExecuteTemplate applies a parsed template to the specified data object, writing the output to wr.
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateEmbed(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // ReadAll reads from r until an error or EOF and returns the data it read.
	fmt.Println("Response Body:", string(body))       // string returns a string containing the contents of the slice.
}
