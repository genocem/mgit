package run

import (
	"fmt"
	"mgit/internal/model"
	"os/exec"
	"strings"
)

func applyCommandToRepo(repo model.Repo, command string) ([]byte, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = repo.Path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error executing command '%s' on repo: '%s': %v\nOutput: %s", command, repo.Name, err, output)
	}
	return output, nil
}
func applyCommandToRepos(repos []model.Repo, command string) error {
	for _, repo := range repos {
		output, err := applyCommandToRepo(repo, command)
		if err != nil {
			return fmt.Errorf("error applying command '%s' to repo '%s': %v", command, repo.Name, err)
		}
		fmt.Printf("Output for %s:\n%s\n", repo.Name, output)
	}
	return nil
}
