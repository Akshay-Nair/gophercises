package db

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRmnTsk(t *testing.T) {
	taskList, err := GetRemainingTask()
	assert.Equalf(t, reflect.TypeOf(taskList).String(), "[]string", "they should be equal")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestAddNwTsk(t *testing.T) {
	err := AddTask("jogging")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestMrkTsk(t *testing.T) {
	err := MarkTask("jogging")
	assert.Equalf(t, err, nil, "both should be equal")
}

func TestNegDo(t *testing.T) {
	tempFthRmnTsk := fetchRemainingTask
	mrkTstDn := markTestDone
	defer func() {
		fetchRemainingTask = tempFthRmnTsk
		markTestDone = mrkTstDn
	}()

	fetchRemainingTask = func() ([]string, error) {
		return []string{}, errors.New("error occured")
	}

	err := DoTask([]int{1})
	assert.Equalf(t, "error occured", err.Error(), "they should be equal")

	markTestDone = func(task string) error {
		return errors.New("error occured")
	}

	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file

	fetchRemainingTask = func() ([]string, error) {
		return []string{"a", "b"}, nil
	}

	DoTask([]int{1})
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("was not deleted", content)
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	os.Stdout = oldStdout
	file.Close()
	os.Remove("testing.txt")
}

func TestInvDoTsk(t *testing.T) {
	var err error
	var (
		list1, list2 []string
	)
	//testing with empty TODO
	list1, err = GetRemainingTask()
	if err != nil {
		t.Error("error occured while fetching tasks")
	}
	err = DoTask([]int{95})
	list2, err = GetRemainingTask()
	if err != nil {
		t.Error("error occured while fetching tasks")
	}
	assert.Equalf(t, len(list1), len(list2), "both should be equal")
	//testing with one element added
	AddTask("hiking")
	list1, err = GetRemainingTask()
	if err != nil {
		t.Error("error occured while fetching tasks")
	}
	err = DoTask([]int{95})
	list2, err = GetRemainingTask()
	if err != nil {
		t.Error("error occured while fetching tasks")
	}
	assert.Equalf(t, len(list1), len(list2), "both should be equal")

}

func TestVldDoTsk(t *testing.T) {
	err := DoTask([]int{1, 2, 3, 66})
	val := (err == nil) || (err.Error() == "TODO list is empty")
	assert.Equalf(t, val, true, "they should be equal")
	AddTask("hiking")
	err = DoTask([]int{1})
	assert.Equalf(t, val, true, "they should be equal")
}

func TestGetFnsTsk(t *testing.T) {
	taskList, err := GetFinishedTask()
	assert.Equalf(t, reflect.TypeOf(taskList).String(), "[]string", "they should be equal")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestNegConnection(t *testing.T) {
	tempExit := exit
	tempGetHome := getHomeDir
	defer func() {
		exit = tempExit
		getHomeDir = tempGetHome
	}()
	exit = func(i int) {

	}
	getHomeDir = func() (string, error) {
		return "", errors.New("error occured")
	}
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	Connection()
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	val, _ := regexp.Match("error occured", content)
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	os.Stdout = oldStdout
	file.Close()
	os.Remove("testing.txt")
}
