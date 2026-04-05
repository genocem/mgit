package run

import (
	"fmt"
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
)

func SwitchProject(name string) {
	if name == "" {
		log.Fatalf("Please provide a project name.")
	}
	if _, err := store.GetProject(name); err != nil {
		log.Fatalf("Project '%s' does not exist.", name)
	}
	if err := config.SetCurrentProject(name); err != nil {
		log.Fatalf("Error changing project: %v\n", err)
	}
	fmt.Printf("Project changed to: %s\n", name)
}
