package run

import (
	"log"
	"mgit/internal/store"
)

func DeleteReposByNamesFunc(names []string) {
	for _, name := range names {
		if err := store.DeleteRepoByName(name); err != nil {
			log.Printf("failed to delete repo with name %s : %s", name, err)
		}
	}
}
