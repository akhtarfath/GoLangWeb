package entity

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(w, r) // call next middleware
	fmt.Println("After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil { // catch error
			fmt.Println("Error: ", err)
			w.WriteHeader(http.StatusInternalServerError) // 500
			_, _ = fmt.Fprint(w, "Something went wrong, error: ", err)
		}
	}()

	middleware.Handler.ServeHTTP(w, r)
}
