package db

import (
	"os"
	"github.com/boltdb/bolt"
	"fmt"
	"reflect"
	"errors"
)

var DbInstance *bolt.DB
const bucketName = "task"

func AddTask( task string ) error {
	err := DbInstance.Update( func(tx *bolt.Tx)error{
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err!=nil{
			return err
		}
		err = BucketInstance.Put([]byte(task), []byte("todo"))
		return err
	})
	return err
}

func GetRemainingTask() ( []string, error ) {
	complete := []byte("todo") 
	var tasks []string
	err := DbInstance.Update( func(tx *bolt.Tx) error {
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err!=nil{
			return err
		}
		cursor := BucketInstance.Cursor()
		for key, value := cursor.First(); key != nil ; key, value = cursor.Next(){
			if reflect.DeepEqual(value, complete){
				tasks = append(tasks, string(key))
			}
		}
		return err
	})
	return tasks, err
}

func MarkTask( task string ) error {
	err := DbInstance.Update( func(tx *bolt.Tx)error{
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err!=nil{
			return err
		}
		err = BucketInstance.Put([]byte(task), []byte("complete"))
		return err
	})
	return err
}

func DoTask( idIndex []int ) error {
	tasks, err := GetRemainingTask()
	if err != nil{
		return err
	}else if len(tasks) == 0{
		return errors.New("TODO list is empty")
	}
	var invalidId []int
	for _, id := range idIndex{
		if id < 1 || id > len(invalidId){
			invalidId = append(invalidId, id)
		} else if MarkTask( tasks[id-1] ) != nil {
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

func GetFinishedTask() ( []string, error ) {
	complete := []byte("complete") 
	var tasks = make([]string, 0)
	err := DbInstance.Update( func(tx *bolt.Tx)error{
		BucketInstance, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err!=nil{
			return err
		}
		cursor := BucketInstance.Cursor()
		for key, value := cursor.First(); key != nil ; key, value = cursor.Next(){
			if reflect.DeepEqual(value, complete){
				tasks = append(tasks, string(key))
			}
		}
		return err
	})
	return tasks, err
} 

func init(){
	const bucketName = "task"
	var err error
	DbInstance, err = bolt.Open("task_db.db", 0644, nil)
	if err!=nil {
		fmt.Println("Following error occured while opening a file ", err)
		os.Exit(1)
	}
}