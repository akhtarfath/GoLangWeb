package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	_ = myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML("<p> Selamat Belajar Go-Lang Web <script>alert('anda di hack!');</script></p>"),
	})
}

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	_ = myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  "<h1> Selamat Belajar Go-Lang Web <script>alert('anda di hack!');</script></h1>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	t.Log("Response Data:", string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err == nil {
		panic(err.Error())
	}
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err == nil {
		panic(err.Error())
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	_ = myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8888/?body=<p>alert!</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Result().Body)
	t.Log("Response Body:", string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err == nil {
		panic(err.Error())
	}
}
