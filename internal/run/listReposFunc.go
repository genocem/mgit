package run

import (
	"log"
	"mgit/internal/model"
	"mgit/internal/store"
	"os"

	"github.com/aquasecurity/table"
)

func ListReposInNamespaceFunc(namespace string, allRepos bool) {
	var repos []model.Repo
	var err error
	if allRepos {
		repos, err = store.GetAllRepos()
	} else {
		repos, err = store.GetAllReposInNamespace(namespace)
	}
	if err != nil {
		log.Fatalf("Error fetching repositories in namespace %s: %v\n", namespace, err)
	}
	if len(repos) == 0 {
		log.Fatalf("No repositories found in namespace: %s\n", namespace)
	}

	table := table.New(os.Stdout)
	table.SetHeaders("Name", "Path", "Namespace", "Current Branch")
	table.SetRowLines(false)

	for _, repo := range repos {
		branch, err := applyCommandToRepo(repo, "git rev-parse --abbrev-ref HEAD")
		if err != nil {
			log.Printf("Error getting current branch for repo %s: %v\n", repo.Name, err)
			branch = []byte("unknown")
		}
		table.AddRow(
			repo.Name,
			repo.Path,
			repo.Namespace,
			string(branch),
		)
	}

	table.Render()
}
