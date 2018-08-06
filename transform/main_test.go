package main

import (
	"bytes"
	"gophercises/transform/primitive"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8008/", nil)
	w := httptest.NewRecorder()
	indexFunc(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	check, _ := regexp.Match("<html[.]+", body)
	assert.NotEqualf(t, true, check, "they should be equal")
}

func TestChoice(t *testing.T) {
	testCase := []struct {
		url    string
		status int
	}{
		{
			"localhost:8008/choice?image=hello&mode=1",
			400,
		},
		{
			"localhost:8008/choice?image=rafale.jpg&mode=1",
			200,
		},
		{
			"localhost:8008/choice?image=rafale.jpg",
			400,
		},
	}
	for _, test := range testCase {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()
		choiceFunc(w, req)
		resp := w.Code
		assert.Equalf(t, test.status, resp, "they should be equal")
	}
}

func TestUpload(t *testing.T) {
	testCase := []struct {
		filename string
		status   int
		invalid  bool
	}{
		{
			"rafale.jpg",
			200,
			false,
		},
		{
			"rafale.gif",
			400,
			false,
		},
		{
			"rafale",
			400,
			false,
		},
		{
			"",
			400,
			false,
		},
		{
			"",
			500,
			true,
		},
	}
	for _, test := range testCase {
		var r *http.Request
		if !test.invalid {
			file, _ := os.Open("./input/testing.jpg")
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("pic", test.filename)
			io.Copy(part, file)
			writer.Close()
			r, _ = http.NewRequest("POST", "localhost:8008/upload", body)
			r.Header.Set("Content-Type", writer.FormDataContentType())
		} else {
			r, _ = http.NewRequest("POST", "localhost:8008/upload", nil)
		}
		w := httptest.NewRecorder()
		uploadFunc(w, r)
		resp := w.Code
		assert.Equalf(t, test.status, resp, "they should be equal")
	}
}

func TestM(t *testing.T) {
	tmp := listenAndServe
	defer func() {
		listenAndServe = tmp
	}()
	listenAndServe = func(port string, mux http.Handler) error {
		return nil
	}
	assert.NotPanicsf(t, main, "panic should not occur")
}

func TestMain(m *testing.M) {
	tmp := batchTarnsformFunc
	defer func() {
		batchTarnsformFunc = tmp
	}()
	batchTarnsformFunc = func(abc string, i []int, r []int, ext string) []primitive.TransformInfo {
		os.Create("./input/testing.jpg")
		defer os.Remove("./input/testing.jpg")
		file, _ := os.Open("./input/testing.jpg")
		return []primitive.TransformInfo{
			{
				File: file,
				Mode: 1,
				Err:  nil,
			},
		}
	}
	dashtest.ControlCoverage(m)
	m.Run()
}
