package logic

import (
	"mgit/internal/model"
)

func ReposToNamesSlice(repos []model.Repo) []string { // probably useless might delete in a bit
	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, repo.Name)
	}
	return repoNames
}
