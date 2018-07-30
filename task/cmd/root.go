package cmd

import (
	"github.com/spf13/cobra"
)

//MainCmd variable defines the task command
var MainCmd = &cobra.Command{

	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}
