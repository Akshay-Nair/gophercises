package filehandle

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

//TempFile function would create a file with a unique and random name and with a given extension and return a file pointer and an error if one occurs
func TempFile(path string, ext string) (*os.File, error) {
	file, err := ioutil.TempFile(path, "")
	if err == nil {
		defer os.Remove(file.Name())
		file, err = os.Create(file.Name() + "." + ext)
	}

	return file, err
}

//ValFile function is to validate if the image name provided is a valid one with respect to the extension. it returns the validity as bool
func ValFile(fileName string) bool {
	valid := false
	tmp := strings.Split(fileName, ".")
	if len(tmp) == 2 {
		extCheck := reflect.DeepEqual(tmp[1], "png") || reflect.DeepEqual(tmp[1], "jpg") || reflect.DeepEqual(tmp[1], "jpeg")
		valid = extCheck
	}
	return valid
}
