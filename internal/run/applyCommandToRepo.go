package run

import (
	"fmt"
	"mgit/internal/model"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func applyCommandToRepo(repo model.Repo, args []string) ([]byte, error) {
	parts := make([]string, len(args))
	for i, a := range args {
		parts[i] = strconv.Quote(a)
	}
	command := strings.Join(parts, " ")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	cmd.Dir = repo.Path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error executing command '%s' on repo: '%s': %v\nOutput: %s", command, repo.Name, err, output)
	}
	return output, nil
}
func applyCommandToRepos(repos []model.Repo, args []string) error {
	for _, repo := range repos {
		output, err := applyCommandToRepo(repo, args)
		if err != nil {
			return fmt.Errorf("error applying command '%s' to repo '%s': %v", args, repo.Name, err)
		}
		fmt.Printf("Output for %s:\n%s\n", repo.Name, output)
	}
	return nil
}
