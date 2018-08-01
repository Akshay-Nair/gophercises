package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var closeDB = func() {
	db.DbInstance.Close()
}

var addNewTask = db.AddTask

//variable addCommand defines the functionality and usage of add subcommand.
var addCommand = &cobra.Command{

	Use:   "add",
	Short: "add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if len(task) == 0 {
			fmt.Println("invalid argument")
			return
		}
		if addNewTask(task) != nil {
			fmt.Println("Failed to add ", task, "task")
		} else {
			fmt.Println(task, "added successfully")
		}
		closeDB()
	},
}

func init() {

	MainCmd.AddCommand(addCommand)

}
