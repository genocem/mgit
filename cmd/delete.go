package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete (repo or project)",
	Aliases: []string{"del"},
	Short:   "Delete a repository or project by name",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println("Usage: mgit delete repo --name <name> --project <project>")
			fmt.Println("or: mgit delete project --name <name>")
			return
		}

	}}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
