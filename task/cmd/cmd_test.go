package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCmd(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	testCase := []struct {
		arg      []string
		expexted string
		msg      string
	}{
		{
			[]string{},
			"invalid argument",
			"they must be equal",
		},
		{
			[]string{"swimming"},
			"added successfully",
			"they must be equal",
		},
	}
	closeDB = func() {

	}
	for _, test := range testCase {
		addCommand.Run(addCommand, test.arg)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error("error occured while test case : ", err)
		}
		val, _ := regexp.Match(test.expexted, content)
		assert.Equalf(t, val, true, test.msg)
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	file.Close()
}
