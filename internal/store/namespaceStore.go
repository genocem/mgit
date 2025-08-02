package store

import (
	"fmt"
	"mgit/internal/model"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var CREATE_NAMESPACE_TABLE_QUERY string = `
CREATE TABLE IF NOT EXISTS namespaces (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE
	)`

var ADD_NAMESPACE_QUERY string = `INSERT INTO namespaces (name) VALUES (?) ON CONFLICT(name) DO NOTHING`
var ADD_DEFAULT_NAMESPACE_QUERY string = `INSERT INTO namespaces (name) VALUES ('default');`
var GET_ALL_NAMESPACES_QUERY string = `SELECT name FROM namespaces`
var GET_NAMESPACE_QUERY string = `SELECT name FROM namespaces WHERE name = ?`
var GET_ALL_REPOS_IN_NAMESPACE_QUERY string = `SELECT id, name, path, namespace FROM repos WHERE namespace = ?`
var DELETE_NAMESPACE_QUERY string = `DELETE FROM namespaces WHERE name = ?`

func GetAllNamespaces() ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(GET_ALL_NAMESPACES_QUERY)
	if err != nil {
		return nil, fmt.Errorf("failed to query namespaces: %w", err)
	}
	defer rows.Close()

	var namespaces []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan namespace rows: %w", err)
		}
		namespaces = append(namespaces, name)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	return namespaces, nil
}
func GetNamespace(name string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	var namespace string
	row := db.QueryRow(GET_NAMESPACE_QUERY, name)
	if err := row.Scan(&namespace); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("namespace with name '%s' does not exist", name)
		}
		return "", fmt.Errorf("failed to scan namespace: %w", err)
	}
	return namespace, nil
}

func AddNamespaceToDB(name string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if name == "" {
		log.Fatal("Namespace name cannot be empty.")
	}

	if name, _ := GetNamespace(name); name != "" {
		return fmt.Errorf("namespace '%s' already exists", name)
	}

	_, err = db.Exec(ADD_NAMESPACE_QUERY, name)
	if err != nil {
		return fmt.Errorf("failed to add namespace: %w", err)
	}

	log.Printf("Namespace '%s' added successfully. in addNamespaceToDB func", name)
	return nil
}

func DeleteNamespace(name string) error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if name == "" {
		log.Fatal("Namespace name cannot be empty.")
	}
	if _, err := GetNamespace(name); err != nil {
		return fmt.Errorf("namespace '%s' does not exist", name)
	}

	_, err = db.Exec(DELETE_NAMESPACE_QUERY, name)
	if err != nil {
		return fmt.Errorf("failed to delete namespace '%s': %v", name, err)
	}

	log.Printf("Namespace '%s' deleted successfully.", name)
	return nil
}

func GetAllReposInNamespace(namespace string) ([]model.Repo, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(GET_ALL_REPOS_IN_NAMESPACE_QUERY, namespace)
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
	return repos, nil
}
