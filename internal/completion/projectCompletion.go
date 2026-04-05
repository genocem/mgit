package completion

import (
	"log"
	"mgit/internal/store"

	"github.com/spf13/cobra"
)

func ProjectCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	projects, err := store.GetAllProjects()
	if err != nil {
		log.Printf("Error fetching projects: %v", err)
		return nil, cobra.ShellCompDirectiveError
	}
	return projects, cobra.ShellCompDirectiveNoFileComp
}
