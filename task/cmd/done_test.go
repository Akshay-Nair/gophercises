package cmd

import (
	"errors"
	"gophercises/task/db"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/atrox/homedir"
	"github.com/stretchr/testify/assert"
)

func TestNegDn(t *testing.T) {
	tempdone := fetchFinishedTask
	defer func() {
		fetchFinishedTask = tempdone
	}()
	fetchFinishedTask = func() ([]string, error) {
		return []string{}, errors.New("error occured")
	}
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	doneCommand.Run(doneCommand, []string{"hello"})
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
func TestDnCmd(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	doneCommand.Run(doneCommand, []string{})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("List of Completed tasks", content)
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}

func TestEmptyDone(t *testing.T) {
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
	doneCommand.Run(doneCommand, []string{})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("No Tasks finished", content)
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()
}
