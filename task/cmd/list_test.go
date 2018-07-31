package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"task/db"
	"testing"

	"github.com/atrox/homedir"
	"github.com/stretchr/testify/assert"
)

func TestEmptyList(t *testing.T) {
	db.DbInstance.Close()
	dir, _ := homedir.Dir()
	dir = dir + "/task_db.db"
	file, _ := os.OpenFile(dir, os.O_TRUNC, 0666)
	file.Truncate(0)
	file.Close()
	db.Connection()
	file, _ = os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	listCommand.Run(listCommand, []string{})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("empty", content)
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()

}

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
	file.Close()
}
