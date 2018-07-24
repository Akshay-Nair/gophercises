package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"strings"
	"../db"
)

var addCommand = &cobra.Command{

		Use		:	"add",
		Short	:	"add a new task",
		Run		:	func(cmd *cobra.Command, args []string) {
						task := strings.Join(args," ")
						if len(task) == 0{
							fmt.Println("invalid argument")							
							os.Exit(1)
						}
						if db.AddTask(task) != nil{
							fmt.Println("Failed to add ", task, "task")
						}else{
							fmt.Println(task, "added successfully")
						}
						db.DbInstance.Close()
					},

	}

func init()	{

	MainCmd.AddCommand(addCommand)

}