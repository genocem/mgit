package run

import (
	"fmt"
	"mgit/internal/model"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

type result struct {
	err error
}

func applyCommandToRepos(repos []model.Repo, args []string) error {
	ch := make(chan result)
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo model.Repo, ch chan<- result) {
			defer wg.Done()
			output, err := applyCommandToRepo(repo, args)
			if err == nil {
				fmt.Printf("Output for %s:\n%s\n", repo.Name, output)
			} else {
				ch <- result{err: err}
			}
		}(repo, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var finalErr string
	for res := range ch {
		finalErr += fmt.Sprintf("\n%s\n", res.err)
	}
	if finalErr != "" {
		return fmt.Errorf(finalErr)
	}

	return nil
}
