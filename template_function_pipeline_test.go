package main

import (
	"GolangWeb/entity"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string { // sayHello is a custom function that returns the upper case of a string.
			return "Hello " + value
		},
		"upper": func(value string) string { // upper is a custom function that returns the upper case of a string.
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`)) // Parse parses text as a template body.
	_ = t.ExecuteTemplate(w, "FUNCTION", entity.MyPage{
		Name: "Muhammad Fathan Aulia",
	}) //	ExecuteTemplate applies a parsed template to the specified data object, and writes the output to w.
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateFunctionPipelines(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	t.Log("Response Body:", string(body))
}
