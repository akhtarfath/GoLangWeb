package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-GO-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	_, _ = fmt.Fprintln(w, "Cookie has been set")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-GO-Name")
	if err != nil {
		_, _ = fmt.Fprintln(w, "Cookie not found")
		return
	}

	if cookie.Value == "" {
		_, _ = fmt.Fprintln(w, "Cookie value is empty")
	} else {
		_, _ = fmt.Fprintf(w, "Hello %s\n", cookie.Value)
	}
}

func TestCookieWithHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func TestSetCookie(t *testing.T) {
	// Create a new request
	request := httptest.NewRequest("GET", "localhost:8888/set-cookie?name=Muhammad%20Fathan%20A.", nil)
	recorder := httptest.NewRecorder() // NewRecorder returns an initialized ResponseRecorder.

	SetCookie(recorder, request) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.

	cookies := recorder.Result().Cookies() // Result returns the Result of the ResponseRecorder.

	fmt.Println("Response Cookies:", cookies) // [X-GO-Name=Muhammad Fathan]
}

func TestAddCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8888/", nil)
	recorder := httptest.NewRecorder() // NewRecorder returns an initialized ResponseRecorder.
	request.AddCookie(&http.Cookie{    // AddCookie adds a cookie to the request.
		Name:  "X-GO-Name",
		Value: "Muhammad Fathan A.",
	})

	GetCookie(recorder, request) // ServeHTTP calls f(w, r). ServeHTTP implements http.Handler.

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body)) // X-GO-Name=Muhammad Fathan
}
