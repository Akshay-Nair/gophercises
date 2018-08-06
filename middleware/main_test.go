package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/stretchr/testify/assert"
)

type middlewareInterface interface {
	middleware(app http.Handler) http.HandlerFunc
}

func TestPanicFunc(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotEqualf(t, err, nil, "they should not be equal")
	}()
	var w http.ResponseWriter
	var r http.Request
	panicfunc(w, &r)
}

func TestHello(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8008/hello", nil)
	w := httptest.NewRecorder()
	hello(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEqualf(t, "hello world", string(body), "they should be equal")
}

func TestDebug(t *testing.T) {
	urlLink := "locahost:8008/debug?"
	testCase := []struct {
		subPath  string
		response string
		msg      string
	}{
		{
			"",
			"Invalid File Path Provided",
			"they should be equal",
		},
		{
			"path=./main.go&line=1",
			"func",
			"they should be equal",
		},
		{
			"path=./main.go",
			"Invalid File Path Provided",
			"they should be equal",
		},
		{
			"path=./testing.txt&line=1",
			"testing.txt: permission denied",
			"they should be equal",
		},
	}
	for _, test := range testCase {
		req := httptest.NewRequest("GET", urlLink+test.subPath, nil)
		w := httptest.NewRecorder()
		debugFunc(w, req)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		contentValidate, _ := regexp.Match(test.response, body)
		assert.Equalf(t, true, contentValidate, test.msg)
	}
	//assert.NotEqualf(t, "invalid ", string(body), "they should be equal")
}

func TestMiddleware(t *testing.T) {
	handler := http.HandlerFunc(panicfunc)
	executeRequest("Get", "/panic", middleware(handler), t, "<a href=.+</a>")
}

func executeRequest(method string, url string, handler http.Handler, t *testing.T, pattern string) {
	req, err := http.NewRequest(method, url, nil)
	if err == nil {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		outputValidate, _ := regexp.Match(pattern, rr.Body.Bytes())
		assert.Equalf(t, true, outputValidate, "they should be equal")
	}
}

func TestNegMiddleware(t *testing.T) {
	tempf := findAndReplacePath
	defer func() {
		findAndReplacePath = tempf
	}()
	findAndReplacePath = func(stacktrace string) (string, error) {
		return "", errors.New("error occured")
	}
	handler := http.HandlerFunc(panicfunc)
	executeRequest("Get", "/panic", middleware(handler), t, "error occured")
}
func TestM(t *testing.T) {
	tmplistenAndServe := listenAndServeFunc
	defer func() {
		listenAndServeFunc = tmplistenAndServe
	}()
	listenAndServeFunc = func(port string, hanle http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")
}

func TestMain(m *testing.M) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE, 0000)
	file.Close()
	dashtest.ControlCoverage(m)
	m.Run()
	os.Remove("testing.txt")
}
