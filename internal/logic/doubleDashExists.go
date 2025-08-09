package logic

import (
	"os"
	"slices"
)

func DoubleDashExists() bool {
	return slices.Contains(os.Args, "--")
}
