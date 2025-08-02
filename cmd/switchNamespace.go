package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var switchNamespaceCmd = &cobra.Command{
	Use:               "switch-namespace <name>",
	Short:             "Change the current namespace",
	ValidArgsFunction: completion.NamespaceCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: mgit switch-namespace <name>")
			return
		}
		name := args[0]
		run.SwitchNamespace(name)
	},
}

func init() {
	rootCmd.AddCommand(switchNamespaceCmd)
}
