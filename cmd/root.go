package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mgit [command] [flags] ",
	Short: "run a command on multiple git repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("error executing mgit command: %v", err)
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringP("namespace", "n", config.GetCurrentNamespace(), "Namespace for the repos")
	rootCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}
