package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var fetchRemainingTask = db.GetRemainingTask

//variable lisCommand defines the functionality and usage of list subcommand
var listCommand = &cobra.Command{

	Use:   "list",
	Short: "list all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := fetchRemainingTask()
		if err != nil {
			fmt.Println("Following error occured during the operation : ", err)
		} else if len(tasks) == 0 {
			fmt.Println("TODO list is empty")
		} else {
			fmt.Println("List of tasks : ")
			for i, task := range tasks {
				fmt.Println(i+1, ") ", task)
			}
		}
		closeDB()
	},
}

func init() {

	MainCmd.AddCommand(listCommand)

}
