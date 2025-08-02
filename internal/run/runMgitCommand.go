package run

import (
	"fmt"
	"log"
	"mgit/internal/logic"
	"mgit/internal/model"
	"mgit/internal/store"
	"strings"
)

func RunMgitCommand(repos []string, namespace string, args []string) {

	allRepos, err := store.GetAllReposInNamespace(namespace)
	if err != nil {
		log.Fatalf("Error fetching repositories in namespace %s: %v\n", namespace, err)
	}
	if len(allRepos) == 0 {
		log.Fatalf("No repositories found in the namespace %s.\n", namespace)
	}
	var approvedRepos []model.Repo
	if len(repos) == 1 && repos[0] == "all" {
		approvedRepos = allRepos
	} else {
		if ok, name := logic.AllSliceElemsExistInReposSlice(repos, allRepos); !ok {
			log.Fatalf("the repository %s doesn't exist in the namespace %s.\n", name, namespace)
		}
		approvedRepos = logic.NamesToRepoSlice(repos, allRepos)
	}

	err = applyCommandToRepos(approvedRepos, strings.Join(args, " "))
	if err != nil {
		log.Fatalf("Error applying command to repositories: %v\n", err)
	}
	fmt.Printf("Command '%s' executed successfully on repositories: %s\n", strings.Join(args, " "), strings.Join(logic.ReposToNamesSlice(approvedRepos), ", "))
}
