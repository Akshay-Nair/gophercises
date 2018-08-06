package cmd

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

func TestSetCmd(t *testing.T) {
	key = "hello123"
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	key = "hello123"
	testCase := []struct {
		args     []string
		expected string
		msg      string
	}{
		{
			[]string{"abc", "abc123"},
			"saved",
			"they should be equal",
		},
		{
			[]string{"abc"},
			"invalid input",
			"they should be equal",
		},
		{
			[]string{},
			"invalid input",
			"they should be equal",
		},
	}

	for _, test := range testCase {
		setCmd.Run(setCmd, test.args)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		val, _ := regexp.Match(test.expected, content)
		assert.Equalf(t, val, true, test.msg)
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	file.Close()
}

func TestNegSet(t *testing.T) {
	key = "hello123"
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	key = "hello123"
	tempFunc := encryptFunc
	defer func() {
		encryptFunc = tempFunc
	}()
	//setting the encrypFunc with a function which would always return an error
	encryptFunc = func(key string, text string) (string, error) {
		return "", errors.New("error occured")
	}

	setCmd.Run(setCmd, []string{"abc", "abc123"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	val, _ := regexp.Match("Following error occured", content)
	assert.Equalf(t, val, true, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}

func TestIvdGetCmd(t *testing.T) {
	key = "hello123"
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	testCase := []struct {
		args     []string
		expected string
		msg      string
	}{
		{
			[]string{},
			"invalid input",
			"they should be equal",
		},
		{
			[]string{"twitte11r"},
			"not found",
			"they should be equal",
		},
	}

	for _, test := range testCase {
		getCmd.Run(getCmd, test.args)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		val, _ := regexp.Match(test.expected, content)
		assert.Equalf(t, val, true, test.msg)
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	file.Close()
}

func TestVldGet(t *testing.T) {
	key = "hello123"
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	var output []string
	testCase := []struct {
		args     []string
		expected string
		msg      string
	}{
		{
			[]string{"abc"},
			"Secret Key",
			"they should be equal",
		},
	}

	for _, test := range testCase {
		getCmd.Run(getCmd, test.args)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		val, _ := regexp.Match(test.expected, content)
		output = append(output, string(content))
		assert.Equalf(t, val, true, test.msg)
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	fmt.Println(output)
	file.Close()

}

func TestNegGet(t *testing.T) {
	key = "hello123"
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	key = "hello123"
	tempFunc := decryptFunc
	defer func() {
		decryptFunc = tempFunc
	}()
	//setting the encrypFunc with a function which would always return an error
	decryptFunc = func(key string, hexCode string) (string, error) {
		return "", errors.New("error occured")
	}

	getCmd.Run(getCmd, []string{"abc", "abc123"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	val, _ := regexp.Match("following error occured", content)
	assert.Equalf(t, val, true, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
