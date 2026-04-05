package store

import (
	"fmt"
	"mgit/internal/model"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var CREATE_PROJECT_TABLE_QUERY string = `
CREATE TABLE IF NOT EXISTS projects (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE
	)`

var ADD_PROJECT_QUERY string = `INSERT INTO projects (name) VALUES (?) ON CONFLICT(name) DO NOTHING`
var ADD_DEFAULT_PROJECT_QUERY string = `INSERT INTO projects (name) VALUES ('default');`
var GET_ALL_PROJECTS_QUERY string = `SELECT name FROM projects`
var GET_PROJECT_QUERY string = `SELECT name FROM projects WHERE name = ?`
var GET_ALL_REPOS_IN_PROJECT_QUERY string = `SELECT id, name, path, project FROM repos WHERE project = ?`
var DELETE_PROJECT_QUERY string = `DELETE FROM projects WHERE name = ?`

func GetAllProjects() ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(GET_ALL_PROJECTS_QUERY)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan project rows: %w", err)
		}
		projects = append(projects, name)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	return projects, nil
}
func GetProject(name string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	var project string
	row := db.QueryRow(GET_PROJECT_QUERY, name)
	if err := row.Scan(&project); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("project with name '%s' does not exist", name)
		}
		return "", fmt.Errorf("failed to scan project: %w", err)
	}
	return project, nil
}

func AddProjectToDB(name string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if name == "" {
		log.Fatal("Project name cannot be empty.")
	}

	if name, _ := GetProject(name); name != "" {
		return fmt.Errorf("project '%s' already exists", name)
	}

	_, err = db.Exec(ADD_PROJECT_QUERY, name)
	if err != nil {
		return fmt.Errorf("failed to add project: %w", err)
	}

	log.Printf("Project '%s' added successfully. in addProjectToDB func", name)
	return nil
}

func DeleteProject(name string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if name == "" {
		log.Fatal("Project name cannot be empty.")
	}
	if _, err := GetProject(name); err != nil {
		return fmt.Errorf("project '%s' does not exist", name)
	}

	_, err = db.Exec(DELETE_PROJECT_QUERY, name)
	if err != nil {
		return fmt.Errorf("failed to delete project '%s': %v", name, err)
	}

	log.Printf("Project '%s' deleted successfully.", name)
	return nil
}

func GetAllReposInProject(project string) ([]model.Repo, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(GET_ALL_REPOS_IN_PROJECT_QUERY, project)
	if err != nil {
		return nil, fmt.Errorf("failed to query repos: %w", err)
	}
	defer rows.Close()

	var repos []model.Repo
	for rows.Next() {
		var r model.Repo
		if err := rows.Scan(&r.ID, &r.Name, &r.Path, &r.Project); err != nil {
			return nil, fmt.Errorf("failed to scan repo: %w", err)
		}
		repos = append(repos, r)
	}
	return repos, nil
}
