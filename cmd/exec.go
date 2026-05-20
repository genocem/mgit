package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/config"
	"mgit/internal/run"
	"os"

	"github.com/spf13/cobra"
)

var execCommand = &cobra.Command{
	Use:               "exec [repo1 repo2...] -- command to run",
	Short:             "execute a command on multiple git repositories",
	ValidArgsFunction: completion.RepoCompletion,
	PreRun: func(cmd *cobra.Command, args []string) {
		if cmd.ArgsLenAtDash() < 0 || cmd.ArgsLenAtDash() == len(args)  {
			fmt.Print("wrong command usage")
			cmd.Help()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		doubeDashPosition := cmd.ArgsLenAtDash()
		repos := args[:doubeDashPosition]
		command := args[doubeDashPosition:]

		project, _ := cmd.Flags().GetString("project")
		if project == "" {
			project = config.GetCurrentProject()
		}

		run.RunMgitCommand(repos, project, command)
	},
}

func init() {
	rootCmd.AddCommand(execCommand)
}
