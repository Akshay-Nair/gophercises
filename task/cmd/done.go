package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

//variable doneCommand defines the functionality and usage of done subcommand
var doneCommand = &cobra.Command{

	Use:   "done",
	Short: "list completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetFinishedTask()
		if err != nil {
			fmt.Println("Following error occured during the operation : ", err)
		} else if len(tasks) == 0 {
			fmt.Println("No Tasks finised")
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
