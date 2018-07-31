package cmd

import (
	"github.com/spf13/cobra"
)

//key variable is for storing the flag value of --key or -k
var key string

//MainCmd defines the root command secret
var MainCmd = &cobra.Command{

	Use:   "secret",
	Short: "manage secret keys",
	Long:  "store the keys in encrypted form",
}

func init() {
	MainCmd.PersistentFlags().StringVar(&key, "key", "", "encryption key (required)")
}
