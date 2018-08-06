//Package db is to perform the read write operations on boltDB.
package db

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/atrox/homedir"
	"github.com/boltdb/bolt"
)

//DbInstance is the variable which would be containing the instance created while opening the connection with database
var DbInstance *bolt.DB

const bucketName = "task"

var markTestDone = MarkTask

var fetchRemainingTask = GetRemainingTask

var exit = os.Exit

var getHomeDir = homedir.Dir

//AddTask is to add new task into the database.
func AddTask(task string) error {
	err := DbInstance.Update(func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err == nil {
			err = BucketInstance.Put([]byte(task), []byte("todo"))
		}
		return err
	})
	return err
}

//GetRemainingTask for fetching the list of task from the database
//which are yet to be marked as completed.
func GetRemainingTask() ([]string, error) {
	complete := []byte("todo")
	var tasks []string
	err := DbInstance.Update(func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err == nil {
			cursor := BucketInstance.Cursor()
			for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
				if reflect.DeepEqual(value, complete) {
					tasks = append(tasks, string(key))
				}
			}
		}
		return err
	})
	return tasks, err
}

//MarkTask is to mark a task as completed
func MarkTask(task string) error {
	err := DbInstance.Update(func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err == nil {
			err = BucketInstance.Put([]byte(task), []byte("complete"))
		}
		return err
	})
	return err
}

//DoTask takes a list of indexes and marks the tasks corresponding to as done.
//It returns an error if it encounters one.
func DoTask(idIndex []int) error {
	tasks, err := fetchRemainingTask()
	if err != nil {
		return err
	} else if len(tasks) == 0 {
		return errors.New("TODO list is empty")
	}
	var invalidID []int
	for _, id := range idIndex {
		if id < 1 || id > len(tasks) {
			invalidID = append(invalidID, id)
		} else if markTestDone(tasks[id-1]) != nil {
			fmt.Println("Task ", id, " was not deleted")
		} else {
			fmt.Println("Task ", tasks[id-1], "marked as completed")
		}
	}
	if len(invalidID) != 0 {
		fmt.Print("Following task id were not deleted : ")
		for _, i := range invalidID {
			fmt.Print(i, " ")
		}
		fmt.Print("\n")
	}
	return nil
}

//GetFinishedTask returns list of tasks from the database
// which have been marked as completed.
func GetFinishedTask() ([]string, error) {
	complete := []byte("complete")
	var tasks = make([]string, 0)
	err := DbInstance.Update(func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err == nil {
			cursor := BucketInstance.Cursor()
			for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
				if reflect.DeepEqual(value, complete) {
					tasks = append(tasks, string(key))
				}
			}
		}
		return err
	})
	return tasks, err
}

//Connection is the function to create connection with the boltDB
func Connection() {
	const bucketName = "task"
	var err error
	var dir string
	dir, err = getHomeDir()
	if err == nil {
		dir = dir + "/task_db.db"
		DbInstance, err = bolt.Open(dir, 0644, nil)
	}
	if err != nil {
		fmt.Println("Following error occured while DB connection ", err)
		exit(1)
	}
}

func init() {
	Connection()
}
