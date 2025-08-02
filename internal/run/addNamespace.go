package run

import (
	"fmt"
	"log"
	"mgit/internal/logic"
	"mgit/internal/store"
)

func AddNamespace(name string) {
	if name == " " {
		log.Fatalf("Namespace name cannot be empty.")
	}
	if !logic.IsValidNamespace(name) {
		log.Fatalf("Invalid namespace name '%s'. Namespace names must consist of lowercase letters only.", name)
	}
	name2, _ := store.GetNamespace(name)
	if name2 != "" {
		log.Fatalf("Namespace '%s' already exists.\n", name)
	}
	// if err != nil {
	// 	log.Fatalf("Error checking namespace existence: %v\n", err)
	// }
	//error handling was commented out cause we also get an error if the namespace doesn't exist lol
	// like we may also get an error for other things like tables doesn't exist if the db file was tampered with
	//i might in the future make it so that a table doesn't exist error and such would inittiate a database reset or smth

	err := store.AddNamespaceToDB(name)
	if err != nil {
		log.Fatalf("Error adding namespace '%s': %v\n", name, err)
	}
	fmt.Printf("Namespace '%s' added successfully.\n", name)
}
