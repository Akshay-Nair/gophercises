package cmd

import (
	"github.com/spf13/cobra"
)

var MainCmd = &cobra.Command{

		Use 	:	"task",
		Short 	:	"task is a CLI for managing your TODOs.",
		Run 	:	func(cmd *cobra.Command, args []string) {},

	}