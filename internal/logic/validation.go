package logic

import "strings"

func IsValidProject(name string) bool {
	return strings.TrimSpace(name) != ""
}
