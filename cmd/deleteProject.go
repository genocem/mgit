package cmd

import (
	"mgit/internal/completion"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var deleteProjectCmd = &cobra.Command{
	Use:               "project <name> ",
	Short:             "delete a project",
	ValidArgsFunction: completion.ProjectCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		//add deleting projects and such and look into batch deleting
		run.DeleteProjectsByNamesFunc(args)
	},
}

func init() {
	deleteCmd.AddCommand(deleteProjectCmd)
}
