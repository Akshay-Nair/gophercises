package main

import (
	"fmt"
	"gophercises/transform/filehandle"
	"gophercises/transform/primitive"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var html = `<html>
<body>
	<div style="align-content: center">
		<form action="http://127.0.0.1:8008/upload" method="POST" enctype="multipart/form-data">
			<input type="file" name="pic" accept="image/*">
			<br>
			<input type="submit">
		</form> 
	</div>
</body>
</html>`

var htmlBegin = "<html>  <body>"
var htmlEnd = "</body> </html>"
var batchTarnsformFunc = primitive.BatchTransform
var listenAndServe = http.ListenAndServe

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, html)
}

func htmlContent(link string, fileInfo []primitive.TransformInfo, src string, reqLink bool) string {
	content := " "
	tmp := strings.Split(src, "/")
	srcName := tmp[len(tmp)-1]
	for _, fileData := range fileInfo {
		if fileData.Err == nil {
			fileName := fileData.File.Name()
			data := strings.Split(fileName, "/")
			if reqLink {
				fileLocation := fmt.Sprintf("<img src=\"http://127.0.0.1:8008/output/%s\" alt=\"hello\" width=\"400\" height=\"400\">", data[len(data)-1])
				link = fmt.Sprintf("/choice?image=%s&mode=%d", srcName, fileData.Mode)
				content = content + fmt.Sprintf("<a href=\"%s\">%s</a> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", link, fileLocation)
			} else {
				fileLocation := fmt.Sprintf("<img src=\"http://127.0.0.1:8008/output/%s\" alt=\"hello\" width=\"400\" height=\"400\">", data[len(data)-1])
				content = content + fmt.Sprintf("%s &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", fileLocation)
			}
		}
	}
	return htmlBegin + content + htmlEnd
}

func uploadFunc(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("pic")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "following error occured : "+err.Error())
		return
	}
	if len(header.Filename) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	fileData := strings.Split(header.Filename, ".")
	if len(fileData) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	ext := fileData[1]
	switch ext {
	case "jpg", "png", "jpeg":
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	outputFile, err := filehandle.TempFile("./input", ext)
	io.Copy(outputFile, file)
	file.Close()
	outputFile.Close()
	fileSlice := batchTarnsformFunc(outputFile.Name(), []int{0, 1, 2, 3}, []int{100}, ext)
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, htmlContent("/", fileSlice, outputFile.Name(), true))
}

func choiceFunc(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("image")
	imageMode := r.URL.Query().Get("mode")
	if len(imageMode) == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	mode, err := strconv.Atoi(imageMode)
	if !filehandle.ValFile(imageName) || err != nil || len(imageName) == 0 || mode < 0 || mode > 8 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	filePath := fmt.Sprintf("./input/%s", imageName)
	fileSlice := batchTarnsformFunc(filePath, []int{mode}, []int{100, 150, 200, 250}, strings.Split(imageName, ".")[1])
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(200)
	fmt.Fprintf(w, htmlContent("", fileSlice, "", false))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexFunc)
	mux.HandleFunc("/upload", uploadFunc).Methods("POST")
	fsOutput := http.FileServer(http.Dir("./output/"))
	mux.PathPrefix("/output/").Handler(http.StripPrefix("/output/", fsOutput))
	fsInput := http.FileServer(http.Dir("./input/"))
	mux.PathPrefix("/input/").Handler(http.StripPrefix("/input/", fsInput))
	mux.HandleFunc("/choice", choiceFunc)
	listenAndServe(":8008", mux)
}
