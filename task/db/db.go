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

var DbInstance *bolt.DB

const bucketName = "task"

//AddTask is to add new task into the database.
func AddTask(task string) error {
	err := DbInstance.Update(func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		err = BucketInstance.Put([]byte(task), []byte("todo"))
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
		if err != nil {
			return err
		}
		cursor := BucketInstance.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			if reflect.DeepEqual(value, complete) {
				tasks = append(tasks, string(key))
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
		if err != nil {
			return err
		}
		err = BucketInstance.Put([]byte(task), []byte("complete"))
		return err
	})
	return err
}

//DoTask takes a list of indexes and marks the tasks corresponding to as done.
//It returns an error if it encounters one.
func DoTask(idIndex []int) error {
	tasks, err := GetRemainingTask()
	if err != nil {
		return err
	} else if len(tasks) == 0 {
		return errors.New("TODO list is empty")
	}
	var invalidId []int
	for _, id := range idIndex {
		if id < 1 || id > len(tasks) {
			invalidId = append(invalidId, id)
		} else if MarkTask(tasks[id-1]) != nil {
			invalidId = append(invalidId, id)
		} else {
			fmt.Println("Task ", tasks[id-1], "marked as completed")
		}
	}
	if len(invalidId) != 0 {
		fmt.Print("Following task id were not deleted : ")
		for _, i := range invalidId {
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
		if err != nil {
			return err
		}
		cursor := BucketInstance.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			if reflect.DeepEqual(value, complete) {
				tasks = append(tasks, string(key))
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
	dir, err = homedir.Dir()
	if err != nil {
		fmt.Println("error occured while fetching home directory")
		os.Exit(1)
	}
	dir = dir + "/task_db.db"
	DbInstance, err = bolt.Open(dir, 0644, nil)
	if err != nil {
		fmt.Println("Following error occured while opening a file ", err)
		os.Exit(1)
	}
}

func init() {
	Connection()
}
