package run

import (
	"fmt"
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
)

func ShowNamespaces() {
	current := config.GetCurrentNamespace()
	allNamespaces, err := store.GetAllNamespaces()
	if err != nil {
		log.Fatalf("%d", err)
	}
	for _, namespace := range allNamespaces {
		if namespace == current {
			fmt.Printf("%s (Current Namespace)\n", namespace)
		} else {
			fmt.Println(namespace)
		}
	}
}
