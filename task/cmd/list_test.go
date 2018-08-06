package cmd

import (
	"errors"
	"fmt"
	"gophercises/task/db"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/atrox/homedir"
	"github.com/stretchr/testify/assert"
)

func TestNegList(t *testing.T) {
	tempdone := fetchRemainingTask

	defer func() {
		fetchRemainingTask = tempdone
	}()

	fetchRemainingTask = func() ([]string, error) {
		return []string{}, errors.New("error occured")
	}

	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	listCommand.Run(listCommand, []string{})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)

	if err != nil {
		t.Error("error occured while test case : ", err)
	}

	val, _ := regexp.Match("error occured", content)
	assert.Equalf(t, val, true, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}
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
