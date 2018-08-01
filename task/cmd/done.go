package cmd

import (
	"fmt"
	"gophercises/task/db"

	"github.com/spf13/cobra"
)

var fetchFinishedTask = db.GetFinishedTask

//variable doneCommand defines the functionality and usage of done subcommand
var doneCommand = &cobra.Command{

	Use:   "done",
	Short: "list completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := fetchFinishedTask()
		if err != nil {
			fmt.Println("Following error occured during the operation : ", err)
		} else if len(tasks) == 0 {
			fmt.Println("No Tasks finished")
		} else {
			fmt.Println("List of Completed tasks : ")
			for i, task := range tasks {
				fmt.Println(i+1, ") ", task)
			}
		}
		closeDB()
	},
}

func init() {

	MainCmd.AddCommand(doneCommand)

}
