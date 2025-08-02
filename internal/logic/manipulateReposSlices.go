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

// this function has quite the shit name ngl
func AllSliceElemsExistInReposSlice(subset []string, repos []model.Repo) (bool, string) {
	nameSet := make(map[string]struct{}, len(repos))
	for _, r := range repos {
		nameSet[r.Name] = struct{}{}
	}
	for _, s := range subset {
		if _, ok := nameSet[s]; !ok {
			return false, s
		}
	}
	return true, ""
}
func NamesToRepoSlice(names []string, repoSlice []model.Repo) []model.Repo { // i don't care if this is inefficient like who's gonna have 1000 repos to manage
	var filteredRepos []model.Repo
	for _, name := range names {
		for _, repo := range repoSlice {
			if repo.Name == name {
				filteredRepos = append(filteredRepos, repo)
			}
		}
	}
	return filteredRepos
}
