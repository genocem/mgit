package run

import (
	"fmt"
	"log"
	"mgit/internal/logic"
	"mgit/internal/model"
	"mgit/internal/store"
	"strings"
)

func RunMgitCommand(repos []string, project string, args []string) {

	allRepos, err := store.GetAllReposInProject(project)

	if err != nil {
		log.Fatalf("Error fetching repositories in project %s: %v\n", project, err)
	}
	if len(allRepos) == 0 {
		log.Fatalf("No repositories found in the project %s.\n", project)
	}
	var approvedRepos []model.Repo
	if len(repos) == 0 {
		approvedRepos = allRepos
	} else {
		if ok, name := logic.AllSliceElemsExistInReposSlice(repos, allRepos); !ok {
			log.Fatalf("the repository %s doesn't exist in the project %s.\n", name, project)
		}
		approvedRepos = logic.NamesToRepoSlice(repos, allRepos)
	}

	err = applyCommandToRepos(approvedRepos, args)
	if err != nil {
		log.Fatalf("Error applying command to repositories: %v\n", err)
	}
	fmt.Printf("Command '%s' executed successfully on repositories: %s\n", args, strings.Join(logic.ReposToNamesSlice(approvedRepos), ", "))
}
