package main

import (
	"fmt"
	"middleware/dir"
	"net/http"
	"runtime/debug"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func panicfunc(w http.ResponseWriter, r *http.Request) {
	panic("error occured")
}

func middleware(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stackTrace := debug.Stack()
				dir.ExtractFilePath(stackTrace)
				fmt.Fprintln(w, string(stackTrace))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func main() {
	app := http.NewServeMux()
	app.HandleFunc("/", hello)
	app.HandleFunc("/panic", panicfunc)
	http.ListenAndServe(":8008", middleware(app))
}
