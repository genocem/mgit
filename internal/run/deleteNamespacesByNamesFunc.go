package run

import (
	"log"
	"mgit/internal/store"
)

func DeleteNamespacesByNamesFunc(names []string) {
	for _, name := range names {
		if err := store.DeleteNamespace(name); err != nil {
			log.Printf("failed to delete namespace with name %s : %s", name, err)
		}
	}
}
