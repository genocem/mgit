package cmd

import (
	"fmt"
	"mgit/internal/config"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var listRepos = &cobra.Command{
	Use:   "list",
	Short: "List all tracked repositories",
	Run: func(cmd *cobra.Command, args []string) {
		allRepos, _ := cmd.Flags().GetBool("all")
		namespace, _ := cmd.Flags().GetString("namespace")

		if allRepos {
			fmt.Println("Listing all repositories across all namespaces.")
		} else {
			if namespace == "" {
				namespace = config.GetCurrentNamespace()
			}
			fmt.Printf("Listing repositories in namespace: %s\n", namespace)
		}
		run.ListReposInNamespaceFunc(namespace, allRepos)
	},
}

func init() {
	listRepos.Flags().BoolP("all", "A", false, "list repos in *all* namespaces")
	rootCmd.AddCommand(listRepos)
}
