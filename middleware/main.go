package main

import (
	"fmt"
	"middleware/content"
	"middleware/filehandle"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var findAndReplacePath = content.FindAndReplacePath

var middlewareFunc = middleware

var listenAndServeFunc = http.ListenAndServe

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func panicfunc(w http.ResponseWriter, r *http.Request) {
	panic("error occured")
}

func debugFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("path")
	lineNumber := r.URL.Query().Get("line")
	var fileContent string
	line, err := strconv.Atoi(lineNumber)
	if len(filePath) > 0 && len(lineNumber) > 0 && err == nil {
		fileContent, err = filehandle.GetData(filePath)
		if err != nil {
			fileContent = err.Error()
		}
		lexer := lexers.Get("go")
		iterator, _ := lexer.Tokenise(nil, fileContent)
		formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.HighlightLines([][2]int{{line, line}}))
		style := styles.Get("github")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
		formatter.Format(w, style, iterator)
	} else {
		fmt.Fprintln(w, "Invalid File Path Provided")
	}
}

func middleware(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stackTrace := debug.Stack()
				data, err := findAndReplacePath(string(stackTrace))
				if err == nil {
					fmt.Fprintln(w, data)
				} else {
					fmt.Fprintln(w, err.Error())
				}
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func main() {
	app := http.NewServeMux()
	app.HandleFunc("/", hello)
	app.HandleFunc("/panic", panicfunc)
	app.HandleFunc("/debug", debugFunc)
	listenAndServeFunc(":8008", middlewareFunc(app))
}
