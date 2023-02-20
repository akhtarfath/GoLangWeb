package main

import (
	"GolangWeb/entity"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))

	_ = t.ExecuteTemplate(w, "FUNCTION", entity.MyPage{
		Name: "Fathan",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8888", nil) // NewRequest returns a new Request given a method, URL, and optional body.
	recorder := httptest.NewRecorder()                                           // NewRecorder returns an initialized ResponseRecorder.

	TemplateFunction(recorder, request) // SimpleTemplate writes to the ResponseWriter.
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	t.Log("Response Body:", string(body))
}
