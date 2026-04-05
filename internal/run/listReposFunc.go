package run

import (
	"log"
	"mgit/internal/model"
	"mgit/internal/store"
	"os"

	"github.com/aquasecurity/table"
)

func ListReposInProjectFunc(project string, allRepos bool) {
	var repos []model.Repo
	var err error
	if allRepos {
		repos, err = store.GetAllRepos()
	} else {
		repos, err = store.GetAllReposInProject(project)
	}
	if err != nil {
		log.Fatalf("Error fetching repositories in project %s: %v\n", project, err)
	}
	if len(repos) == 0 {
		log.Fatalf("No repositories found in project: %s\n", project)
	}

	table := table.New(os.Stdout)
	table.SetHeaders("Name", "Path", "Project", "Current Branch")
	table.SetRowLines(false)

	for _, repo := range repos {
		command := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}
		branch, err := applyCommandToRepo(repo, command)
		if err != nil {
			log.Printf("Error getting current branch for repo %s: %v\n", repo.Name, err)
			branch = []byte("unknown")
		}
		table.AddRow(
			repo.Name,
			repo.Path,
			repo.Project,
			string(branch),
		)
	}

	table.Render()
}
