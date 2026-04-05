package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var switchProjectCmd = &cobra.Command{
	Use:               "switch-project <name>",
	Short:             "Change the current project",
	ValidArgsFunction: completion.ProjectCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: mgit switch-project <name>")
			return
		}
		name := args[0]
		run.SwitchProject(name)
	},
}

func init() {
	rootCmd.AddCommand(switchProjectCmd)
}
