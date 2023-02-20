package main

import (
	"fmt"
	"net/http"
	"testing"
)

// Content-Disposition: attachment; filename="filename.jpg" is used to force download the file instead of opening it.
// Content-Disposition: inline; filename="filename.jpg" is used to open the file in the browser.
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file") // get file name from query string, e.g. /download?file=filename.jpg
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)          // 400
		_, _ = fmt.Fprint(w, "File name is required") // print error message
		return                                        // stop execution
	}

	//w.Header().Add("Content-Disposition", "inline; filename=\""+filename+"\"") // open in browser
	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"") // force download, not open in browser
	http.ServeFile(w, r, "./resources/"+filename)                                  // serve file
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{ // create server
		Addr:    "localhost:8888",
		Handler: http.HandlerFunc(DownloadFile), // use http.HandlerFunc to convert function to http.Handler
	}

	err := server.ListenAndServe() // start server
	if err != nil {
		panic(err.Error())
	}
}
