package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"task/db"
	"testing"

	"github.com/atrox/homedir"
	"github.com/stretchr/testify/assert"
)

func TestNegAddCmd(t *testing.T) {
	tempadd := addNewTask
	defer func() {
		addNewTask = tempadd
	}()
	addNewTask = func(task string) error {
		return errors.New("error occured")
	}
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	closeDB = func() {

	}
	addCommand.Run(addCommand, []string{"hello"})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("failed to add", content)
	assert.Equalf(t, val, true, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	file.Close()

}
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

func TestMain(m *testing.M) {
	closeDB()
	dir, _ := homedir.Dir()
	dir = dir + "/task_db.db"
	os.Remove(dir)
	db.Connection()
	m.Run()
}
