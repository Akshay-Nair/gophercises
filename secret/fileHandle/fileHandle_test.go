package fileHandle

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tempName := fileName
	fileName = "../temp.csv"
	defer func() {
		fileName = tempName
	}()
	testCase := []struct {
		key      string
		expected string
	}{
		{
			"abc",
			"",
			//"c647e6a250b9a8cd61087c03f15eb914a98808a15ecc",
		},
		{
			"def",
			"",
		},
		{
			"google",
			"",
		},
		{
			"",
			"",
		},
	}
	for _, test := range testCase {
		key, err := GetSecret(test.key)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equalf(t, key, test.expected, "should be equal")
	}
}

func TestInvalid(t *testing.T) {
	tempName := fileName
	fileName = "../abc.csv" //testing for non existing file
	defer func() {
		fileName = tempName
	}()
	secret, _ := GetSecret("xyz")
	assert.Equalf(t, len(secret), 0, "they should be equal")
	err := SetSecret("abc", "def")
	assert.NotEqualf(t, err, "nil", "they should not be equal")
}

func TestWrite(t *testing.T) {
	tempName := fileName
	fileName = "../temp.csv"
	defer func() {
		fileName = tempName
	}()
	testCase := []struct {
		serviceName string
		key         string
		expected    error
	}{
		{
			"twitter",
			"abc",
			nil,
		},
		{
			"google",
			"",
			nil,
		},
		{
			"",
			"",
			nil,
		},
		{
			"",
			"abc123",
			nil,
		},
	}
	for _, test := range testCase {
		err := SetSecret(test.serviceName, test.key)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equalf(t, err, test.expected, "should be equal")
	}
}

func TestNegFile(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	tempFunc := getHomeDir
	tempExit := exit
	defer func() {
		getHomeDir = tempFunc
		exit = tempExit
	}()
	getHomeDir = func() (string, error) {
		return "", errors.New("error occured")
	}
	exit = func(i int) {

	}
	setHomeDir()
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	val, _ := regexp.Match("error", content)
	assert.Equalf(t, val, true, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
