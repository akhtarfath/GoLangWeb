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

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")   // New allocates a new, undefined template with the given name.
	t.Funcs(map[string]interface{}{ // Funcs adds the elements of the argument map to the template's function map.
		"upper": func(value string) string { // upper is a custom function that returns the upper case of a string.
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}`))     // Parse parses text as a template body.
	_ = t.ExecuteTemplate(w, "FUNCTION", entity.MyPage{ //	ExecuteTemplate applies a parsed template to the specified data object, and writes the output to w.
		Name: "Muhammad Fathan Aulia",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateFunctionCreateGlobal(recorder, request)   // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body) // Result returns the Result of the ResponseRecorder.
	t.Log("Response Body:", string(body))
}
