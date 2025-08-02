package completion

import (
	"log"
	"mgit/internal/store"

	"github.com/spf13/cobra"
)

func NamespaceCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespaces, err := store.GetAllNamespaces()
	if err != nil {
		log.Printf("Error fetching namespaces: %v", err)
		return nil, cobra.ShellCompDirectiveError
	}
	return namespaces, cobra.ShellCompDirectiveNoFileComp
}
