package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	// ParseFiles creates a new Template and parses the template definitions from the files.
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	// ExecuteTemplate applies a parsed template to the specified data object, and writes the output to w.
	_ = t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Muhammad A. Fathan",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateLayout(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // Result returns the Result of the ResponseRecorder.
	fmt.Println("Response Body:", string(body))       // Response Body: <html>
}

func TemplateDefineLayout(w http.ResponseWriter, r *http.Request) {
	// ParseFiles creates a new Template and parses the template definitions from the files.
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml"))

	// ExecuteTemplate applies a parsed template to the specified data object, and writes the output to w.
	_ = t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Muhammad A. Fathan",
	})
}

func TestDefineTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateDefineLayout(recorder, request) // SimpleTemplate writes to the ResponseWriter.

	body, _ := ioutil.ReadAll(recorder.Result().Body) // Result returns the Result of the ResponseRecorder.
	fmt.Println("Response Body:", string(body))       // Response Body: <html>
}
