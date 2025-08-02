package store

import (
	"fmt"
	"mgit/internal/model"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var CREATE_REPO_TABLE_QUERY string = `
CREATE TABLE repos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	path TEXT NOT NULL,
	namespace TEXT NOT NULL,
	config TEXT,
	FOREIGN KEY (namespace) REFERENCES namespaces(name) ON DELETE CASCADE
	UNIQUE (path, namespace)
	)`

var ADD_REPO_QUERY string = `
	INSERT INTO repos (name, path, namespace) VALUES (?, ?, ?)
	ON CONFLICT(path, namespace) DO UPDATE SET
	name = excluded.name
	`
var DELETE_REPO_QUERY string = "DELETE FROM repos WHERE name = ?"
var REPOS_EXIST_QUERY string = `SELECT * FROM repos WHERE name IN (?)`
var GET_REPO_QUERY string = `SELECT id, name, path, namespace FROM repos WHERE name = ?`
var GET_ALL_REPOS_QUERY string = `SELECT id, name, path, namespace FROM repos`

func AddRepo(path, name, namespace string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer db.Close()

	_, err = db.Exec(ADD_REPO_QUERY, name, path, namespace)
	if err != nil {
		return fmt.Errorf("failed to add repository: %w", err)
	}
	return nil
}
func DeleteRepoByName(name string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	// execOrError(db, DELETE_REPO_QUERY, repoID)
	_, err = db.Exec(DELETE_REPO_QUERY, name)
	if err != nil {
		return fmt.Errorf("failed to delete repository with ID %s: %v", name, err)
	}
	log.Printf("Repository with ID %s deleted successfully.", name)
	return nil
}

func GetRepo(name string) (model.Repo, error) {
	db, err := getDB()
	if err != nil {
		return model.Repo{}, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	var r model.Repo
	row := db.QueryRow(GET_REPO_QUERY, name)
	if err := row.Scan(&r.ID, &r.Name, &r.Path, &r.Namespace); err != nil {
		if err == sql.ErrNoRows {
			return model.Repo{}, fmt.Errorf("repository with name '%s' does not exist", name)
		}
		return model.Repo{}, fmt.Errorf("failed to scan repo: %w", err)
	}
	return r, nil
}

// this getAllRepos func is for the future if i ever add a -A flag that would get repos from all namespaces
func GetAllRepos() ([]model.Repo, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(GET_ALL_REPOS_QUERY)
	if err != nil {
		return nil, fmt.Errorf("failed to query repos: %w", err)
	}
	defer rows.Close()

	var repos []model.Repo
	for rows.Next() {
		var r model.Repo
		if err := rows.Scan(&r.ID, &r.Name, &r.Path, &r.Namespace); err != nil {
			return nil, fmt.Errorf("failed to scan repo: %w", err)
		}
		repos = append(repos, r)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return repos, nil
}
