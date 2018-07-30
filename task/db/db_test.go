package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRmnTsk(t *testing.T) {
	taskList, err := GetRemainingTask()
	assert.Equalf(t, reflect.TypeOf(taskList).String(), "[]string", "they should be equal")
	assert.Equalf(t, err, nil, "they should be equal")
}

func TestGetFnsTsk(t *testing.T) {
	taskList, err := GetFinishedTask()
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
