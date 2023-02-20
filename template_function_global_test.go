package main

import (
	"GolangWeb/entity"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len "Budi" }}`)) // len is a built-in function that returns the length of a string.
	_ = t.ExecuteTemplate(w, "FUNCTION", entity.MyPage{                    // ExecuteTemplate applies a parsed template to the specified data object, and writes the output to w.
		Name: "Fathan",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateFunctionGlobal(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	t.Log("Response Body:", string(body))
}
