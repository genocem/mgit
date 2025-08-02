package cmd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add (repo or namespace)",
	Short: "Add smth new",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println("Usage: mgit add repo --path <path> --name <name> --namespace <namespace>")
			fmt.Println("or: mgit add namespace --name <name>")
			return
		}
	},
}

func init() {

	rootCmd.AddCommand(addCmd)
}
