package config

import (
	"fmt"
	"net/http"
)

func Routes(handler *http.ServeMux) {
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, _ := fmt.Fprintf(w, "Golang Web")
		fmt.Println(fprintf)
	})
	handler.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, _ := fmt.Fprintf(w, "About Golang Web")
		fmt.Println(fprintf)
	})
	handler.HandleFunc("/about/fathan", func(w http.ResponseWriter, r *http.Request) {
		fprintf, _ := fmt.Fprintf(w, "Fathan Learning Golang Web")
		fmt.Println(fprintf)
	})
}
