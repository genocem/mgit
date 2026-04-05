package run

import (
	"log"
	"mgit/internal/config"
	"mgit/internal/store"
	"os"
	"path/filepath"
)

func AddRepoFunc(path, name, project string) {

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

	path, err = filepath.Abs(path)

	if name == "" {
		name = filepath.Base(path)
	}

	if exists, _ := store.GetRepo(name); exists.Name != "" && exists.Project == project {
		log.Fatalf("repository with name '%s' already exists", name)
	}

	if project != "" {
		if _, err := store.GetProject(project); err != nil {
			log.Fatalf("the specified project '%s' does not exist", project)
		}
	}
	if project == "" {
		project = config.GetCurrentProject()
	}

	if err != nil {
		log.Fatalf("problem getting absolute path of path '%s'", path)
	}
	if err := store.AddRepo(path, name, project); err != nil {
		log.Fatalf("add repo db function got fucked %s", err)
	}
	log.Printf("Repository '%s' added successfully.", name)

}
