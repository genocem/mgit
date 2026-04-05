package completion

import (
	"mgit/internal/config"
	"mgit/internal/logic"
	"mgit/internal/store"
	"strings"

	"github.com/spf13/cobra"
)

func RepoCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if logic.DoubleDashExists() {
		dashIndex := cmd.ArgsLenAtDash()
		if dashIndex <= len(args) {
			return nil, cobra.ShellCompDirectiveDefault
		}
	}

	project, _ := cmd.Flags().GetString("project")
	if project == "" {
		project = config.GetCurrentProject()
	}

	availableRepos, err := store.GetAllReposInProject(project)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	availableRepoNames := logic.ReposToNamesSlice(availableRepos)

	if len(args) > 0 {
		availableRepoNames = logic.DeleteElemsFromSlice(availableRepoNames, args)
	}

	var completions []string
	for _, repo := range availableRepoNames {
		if strings.HasPrefix(repo, toComplete) {
			completions = append(completions, repo)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
