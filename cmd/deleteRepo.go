package cmd

import (
	"mgit/internal/completion"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var deleteRepoCmd = &cobra.Command{
	Use:               "repo <name>",
	Short:             "delete a repository",
	ValidArgsFunction: completion.RepoCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		run.DeleteReposByNamesFunc(args)
	},
}

func init() {
	deleteCmd.AddCommand(deleteRepoCmd)
}
