package main

import (
	"bytes"
	_ "embed"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err.Error())
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100 << 20) // parse request to get multipart form size 100MB
	if err != nil {
		panic(err.Error())
	}

	// file is io.ReadCloser, fileHeader is *multipart.FileHeader, err is error
	// io.ReadCloser is an interface that has Read() and Close() method
	// multipart.FileHeader is a struct that has Name(), Size(), Header(), Open() method
	// error is an interface that has Error() method
	file, fileHeader, err := r.FormFile("file") // get file from request
	if err != nil {
		panic(err.Error())
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename) // create file destination
	if err != nil {
		panic(err.Error())
	}

	_, err = io.Copy(fileDestination, file) // copy file to destination
	if err != nil {
		panic(err.Error())
	}

	name := r.PostFormValue("name") // get name from request
	_ = myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	}) // execute template
}

func TestUploadFormServer(t *testing.T) { // test upload form
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources")))) // static file handler

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

//go:embed resources/noodle-burger.jpeg
var noodleBurger []byte

func TestUpload(t *testing.T) {
	body := new(bytes.Buffer)               // create new buffer, buffer is a byte slice
	writer := multipart.NewWriter(body)     // create new multipart writer
	_ = writer.WriteField("name", "Akhtar") // write name field

	file, _ := writer.CreateFormFile("file", "noodle-burger.jpeg") // create file field, file is io.Writer
	_, err := file.Write(noodleBurger)                             // write file field
	if err != nil {
		return
	}
	_ = writer.Close() // close writer

	request := httptest.NewRequest(http.MethodPost, "localhost:8888/upload", body) // create new request
	request.Header.Set("Content-Type", writer.FormDataContentType())               // set content type, writer.FormDataContentType() is multipart/form-data; boundary=...
	recorder := httptest.NewRecorder()                                             // create new recorder

	Upload(recorder, request) // call upload function

	response := recorder.Result() // get response from recorder
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(response.Body) // close response body

	if response.StatusCode != http.StatusOK { // check status code
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	} else {
		t.Logf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}
