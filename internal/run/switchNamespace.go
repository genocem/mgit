package run

import (
	"fmt"
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
)

func SwitchNamespace(name string) {
	if name == "" {
		log.Fatalf("Please provide a namespace name.")
	}
	if _, err := store.GetNamespace(name); err != nil {
		log.Fatalf("Namespace '%s' does not exist.", name)
	}
	if err := config.SetCurrentNamespace(name); err != nil {
		log.Fatalf("Error changing namespace: %v\n", err)
	}
	fmt.Printf("Namespace changed to: %s\n", name)
}
