package logic

import "mgit/internal/model"

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
