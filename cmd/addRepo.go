package cmd

import (
	"fmt"
	"mgit/internal/completion"
	"mgit/internal/run"
	"os"

	"github.com/spf13/cobra"
)

var addRepoCmd = &cobra.Command{
	Use:   "repo --path <path>",
	Short: "Add a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println("Usage: mgit add repo --path <path> --name <name>")
			fmt.Println("Optional: --namespace <namespace> (default: 'default') ")
			return
		}
		path, _ := cmd.Flags().GetString("path")
		name, _ := cmd.Flags().GetString("name")

		namespace, _ := cmd.Flags().GetString("namespace")
		if path == "" {
			path, _ = os.Getwd()
			fmt.Printf("No path provided, using current directory: %s\n", path)
		}
		run.AddRepoFunc(path, name, namespace)
	},
}

func init() {

	addRepoCmd.Flags().StringP("path", "p", "", "Local path to the repository")
	addRepoCmd.Flags().String("name", "", "Repository name")
	// addRepoCmd.MarkFlagRequired("path") // look into leaving the required flag for leaf children if it comes with problems too
	addRepoCmd.RegisterFlagCompletionFunc("path", completion.AddPathCompletion)
	addRepoCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	addCmd.AddCommand(addRepoCmd)
}
