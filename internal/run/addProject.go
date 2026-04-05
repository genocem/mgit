package run

import (
	"fmt"
	"log"
	"mgit/internal/logic"
	"mgit/internal/store"
)

func AddProject(name string) {
	if name == " " {
		log.Fatalf("Project name cannot be empty.")
	}
	if !logic.IsValidProject(name) {
		log.Fatalf("Invalid project name '%s'. Project names must consist of lowercase letters only.", name)
	}
	name2, _ := store.GetProject(name)
	if name2 != "" {
		log.Fatalf("Project '%s' already exists.\n", name)
	}
	// if err != nil {
	// 	log.Fatalf("Error checking project existence: %v\n", err)
	// }
	//error handling was commented out cause we also get an error if the project doesn't exist lol
	// like we may also get an error for other things like tables doesn't exist if the db file was tampered with
	//i might in the future make it so that a table doesn't exist error and such would inittiate a database reset or smth

	err := store.AddProjectToDB(name)
	if err != nil {
		log.Fatalf("Error adding project '%s': %v\n", name, err)
	}
	fmt.Printf("Project '%s' added successfully.\n", name)
}
