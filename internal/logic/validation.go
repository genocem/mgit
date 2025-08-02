package logic

import (
	"regexp"
)

func IsValidNamespace(name string) bool {
	return regexp.MustCompile(`^[a-z]+$`).MatchString(name)
}
