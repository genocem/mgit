package logic

import (
	"regexp"
)

func IsValidProject(name string) bool {
	return regexp.MustCompile(`^[a-z]+$`).MatchString(name)
}
