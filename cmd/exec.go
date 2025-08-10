package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/config"
	"mgit/internal/logic"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var execCommand = &cobra.Command{
	Use:               "exec [repo1 repo2...] -- command to run",
	Short:             "execute a command on multiple git repositories",
	ValidArgsFunction: completion.RepoCompletion,
	PreRun: func(cmd *cobra.Command, args []string) {
		if !logic.DoubleDashExists() {
			fmt.Print("please provide a command after --")
			cmd.Help()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		doubeDashPosition := cmd.ArgsLenAtDash()
		repos := args[:doubeDashPosition]
		command := args[doubeDashPosition:]

		namespace, _ := cmd.Flags().GetString("namespace")
		if namespace == "" {
			namespace = config.GetCurrentNamespace()
		}

		run.RunMgitCommand(repos, namespace, command)
	},
}

func init() {
	rootCmd.AddCommand(execCommand)
}
