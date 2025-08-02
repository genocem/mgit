package run

import (
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
	"os"
	"path/filepath"
)

func AddRepoFunc(path, name, namespace string) {

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("path '%s' does not exist", path)
	}
	if !info.IsDir() {
		log.Fatalf("path '%s' is not a directory", path)
	}
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {
		log.Fatalf("path '%s' is not a git repository", path)
	}

	if name == "" {
		name = filepath.Base(path)
	}

	if exists, _ := store.GetRepo(name); exists.Name != "" && exists.Namespace == namespace {
		log.Fatalf("repository with name '%s' already exists", name)
	}

	if namespace != "" {
		if _, err := store.GetNamespace(namespace); err != nil {
			log.Fatalf("the specified namespace '%s' does not exist", namespace)
		}
	}
	if namespace == "" {
		namespace = config.GetCurrentNamespace()
	}

	if err := store.AddRepo(path, name, namespace); err != nil {
		log.Fatalf("add repo db function got fucked %s", err)
	}
	log.Printf("Repository '%s' added successfully.", name)

}
