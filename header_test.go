package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-Powered-By", "AkhtarFath")
	_, _ = fmt.Fprintln(w, "Ok!")
}

func TestResponseHeader(t *testing.T) {
	requestHeader := httptest.NewRequest("GET", "http://localhost:8888/request-header", nil)
	requestHeader.Header.Add("content-type", "Application/Json")
	recorder := httptest.NewRecorder()      // NewRecorder returns an initialized ResponseRecorder.
	ResponseHeader(recorder, requestHeader) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.

	fmt.Println("Response Content-Type:", requestHeader.Header.Get("content-type"))
	fmt.Println("Response X-Powered-By:", requestHeader.Header.Get("X-Powered-By"))
}
