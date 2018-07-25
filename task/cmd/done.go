package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"../db"
)

var doneCommand = &cobra.Command{

		Use		:	"done",
		Short	:	"list completed tasks",
		Run		:	func(cmd *cobra.Command, args []string) {
						tasks, err := db.GetFinishedTask()
						if err != nil {
							fmt.Println("Following error occured during the operation : ", err)
						}else if len(tasks) == 0{
							fmt.Println("No Tasks finised")
						}else{
							fmt.Println("List of Completed tasks : ")
							for i, task := range tasks {
								fmt.Println(i+1, ") ", task)
							}
						}
						db.DbInstance.Close()
					},
					
	}

func init()	{

	MainCmd.AddCommand(doneCommand)

}