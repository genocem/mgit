package run

import (
	"log"
	"mgit/internal/store"
)

func DeleteProjectsByNamesFunc(names []string) {
	for _, name := range names {
		if err := store.DeleteProject(name); err != nil {
			log.Printf("failed to delete project with name %s : %s", name, err)
		}
	}
}
