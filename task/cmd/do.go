package cmd

import (
	"strconv"
	"github.com/spf13/cobra"
	"fmt"
	"../db"
)

var doCommand = &cobra.Command{

		Use		:	"do",
		Short	:	"mark a task as complete",
		Run		:	func(cmd *cobra.Command, args []string) {
						var idList []int
						var ( 
							i int
							err error
						)
						for _, id := range args{
							i, err = strconv.Atoi(id)
							if err != nil{
								fmt.Println(id, " is an invalid id")
							}else{
								idList = append(idList, i)
							}
						}
						if len(idList) != 0{ 
							err = db.DoTask(idList)
						}else{
							fmt.Println("invalid argument")
						}
						if err != nil{
							fmt.Println("Following error occured : ", err)
						}
						db.DbInstance.Close()
					},
					
	}

func init()	{

	MainCmd.AddCommand(doCommand)

}