package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLsCmd(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	addCommand.Run(addCommand, []string{"abc"})
	listCommand.Run(listCommand, []string{})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val1, _ := regexp.Match("List of task", content)
	val2, _ := regexp.Match("TODO list is empty", content)
	val := val1 || val2
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
}
