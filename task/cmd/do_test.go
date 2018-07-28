package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoCmd(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	testCase := []struct {
		arg      []string
		expected string
		msg      string
	}{
		{
			[]string{},
			"invalid argument",
			"they must be equal",
		},
		{
			[]string{"1"},
			"marked as completed",
			"they must be equal",
		},
		{
			[]string{"100"},
			"Following task id were not deleted",
			"they should be equal",
		},
	}

	for _, test := range testCase {
		doCommand.Run(doCommand, test.arg)
		file.Seek(0, 0)
		content, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error("error occured while test case : ", err)
		}
		val1, _ := regexp.Match(test.expected, content)
		val2, _ := regexp.Match("TODO list is empty", content)
		val := val1 || val2
		assert.Equalf(t, true, val, "they should be equal")
		file.Truncate(0)
		file.Seek(0, 0)
	}
	os.Stdout = oldStdout
	file.Close()
}
