package cmd

import (
	"mgit/internal/completion"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var deleteNamespaceCmd = &cobra.Command{
	Use:               "namespace <name> ",
	Short:             "delete a namespace",
	ValidArgsFunction: completion.NamespaceCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		//add deleting namespaces and such and look into batch deleting
		run.DeleteNamespacesByNamesFunc(args)
	},
}

func init() {
	deleteCmd.AddCommand(deleteNamespaceCmd)
}
