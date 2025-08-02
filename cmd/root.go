package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/config"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mgit [flags] -- [command to run on a repo]",
	Short: "run a command on multiple git repositories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a command.")
			cmd.Help()
			return
		}
		repos, _ := cmd.Flags().GetStringSlice("repos")
		namespace, _ := cmd.Flags().GetString("namespace")
		if namespace == "" {
			namespace = config.GetCurrentNamespace()
		}

		run.RunMgitCommand(repos, namespace, args)
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("error executing mgit command: %v", err)
	}
	return nil
}

func init() {
	rootCmd.Flags().StringSliceP("repos", "r", []string{}, "List of repositories to run commands on")
	rootCmd.PersistentFlags().StringP("namespace", "n", config.GetCurrentNamespace(), "Namespace for the resource")
	rootCmd.RegisterFlagCompletionFunc("repos", completion.RepoCompletion)
	rootCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}
