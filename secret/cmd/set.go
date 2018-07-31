package cmd

import (
	"fmt"
	"secret/crypt"
	"secret/fileHandle"

	"github.com/spf13/cobra"
)

//setTemplate defines the help template for the set subcommand.
var setTemplate = `Store the key related to a service name in encrypted form

Usage:
  set [flags] [service_name] [secret_key]

Flags:
  -h, --help         help for secret
      --key string   encryption key (required)

Arguments:
service_name	name of the service of which the key is to be saved
secret_key	secret key related to the service name` + "\n\n\n"

//function saveKey performs the operations for saving the secret into the file in encrypted form.
func saveKey(Secret, ServiceName string) {
	encryptedKey, err := crypt.Encrypt(key, Secret)
	if err == nil {
		err = fileHandle.SetSecret(ServiceName, encryptedKey)
	}
	if err == nil {
		fmt.Println("Secret key saved successfully")
	} else {
		fmt.Println("Following error occured while saving the secret key : ", err)
	}

}

//setCmd defines the functionality of the set subcommand.
var setCmd = &cobra.Command{

	Use:   "set",
	Short: "store a secret key in encrypted form",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 || len(key) == 0 {
			fmt.Println(" invalid input ")
			return
		}
		ServiceName := args[0]
		Secret := args[1]
		saveKey(Secret, ServiceName)
	},
}

func init() {
	setCmd.SetUsageTemplate(setTemplate)
	MainCmd.AddCommand(setCmd)
}
