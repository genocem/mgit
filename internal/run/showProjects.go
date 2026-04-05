package run

import (
	"fmt"
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
)

func ShowProjects() {
	current := config.GetCurrentProject()
	allProjects, err := store.GetAllProjects()
	if err != nil {
		log.Fatalf("%d", err)
	}
	for _, project := range allProjects {
		if project == current {
			fmt.Printf("%s (Current Project)\n", project)
		} else {
			fmt.Println(project)
		}
	}
}
