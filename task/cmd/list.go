package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"../db"
)

var listCommand = &cobra.Command{

		Use		:	"list",
		Short	:	"list all the tasks",
		Run		:	func(cmd *cobra.Command, args []string) {
						tasks, err := db.GetRemainingTask()
						if err != nil {
							fmt.Println("Following error occured during the operation : ", err)
						}else if len(tasks) == 0{
							fmt.Println("TODO list is empty")
						}else{
							fmt.Println("List of tasks : ")
							for i, task := range tasks {
								fmt.Println(i+1, ") ", task)
							}
						}
						db.DbInstance.Close()
					},
					
	}

func init()	{

	MainCmd.AddCommand(listCommand)

}