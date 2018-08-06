package filehandle

import (
	"io/ioutil"
	"os"
)

//GetData function would read a file from the path provided
//and return the content of the file and error if one occurs
func GetData(path string) (string, error) {
	var fileContent []byte
	file, err := os.Open(path)
	if err == nil {
		fileContent, err = ioutil.ReadAll(file)
	}
	return string(fileContent), err
}
