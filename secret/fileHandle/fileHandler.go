//Package fileHandle provides methods to read and write key and a value associated to it into a csv file.
package fileHandle

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/atrox/homedir"
)

//variable to set filename.
var fileName string

var getHomeDir = homedir.Dir

var exit = os.Exit

//SetSecret method is to save a key and a value associated into the csv file
//it returns an error if one occurs.
func SetSecret(ServiceName, Key string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err == nil {
		writer := csv.NewWriter(file)
		err = writer.Write([]string{ServiceName, Key})
		writer.Flush()
		file.Close()
	}
	return err
}

//GetSecret method reads a file and returns secret key associated with a service name
// along with an error if one occurs.
func GetSecret(ServiceName string) (string, error) {
	csvFile, _ := os.Open(fileName)
	var key string
	var err error
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		data, er := reader.Read()
		if er != nil {
			if er == io.EOF {
				er = nil
			}
			err = er
			break
		}
		if len(data) == 2 {
			if reflect.DeepEqual(data[0], ServiceName) {
				key = data[1]
				break
			}
		}
	}
	return key, err
}

func setHomeDir() {
	home, err := getHomeDir()
	if err != nil {
		fmt.Println("error while fetching home dir")
		exit(1)
	}
	fileName = home + "/secret.csv"
}

func init() {
	setHomeDir()
}
