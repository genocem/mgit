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
		project, _ := cmd.Flags().GetString("project")

		if allRepos {
			fmt.Println("Listing all repositories across all projects.")
		} else {
			if project == "" {
				project = config.GetCurrentProject()
			}
			fmt.Printf("Listing repositories in project: %s\n", project)
		}
		run.ListReposInProjectFunc(project, allRepos)
	},
}

func init() {
	listRepos.Flags().BoolP("all", "A", false, "list repos in *all* projects")
	rootCmd.AddCommand(listRepos)
}
