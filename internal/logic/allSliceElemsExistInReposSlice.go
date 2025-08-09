package logic

import "mgit/internal/model"

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
