package store

import (
	"fmt"
	"os"
	"path/filepath"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "/home/ahmed/.mgit/mgit.db"

// i added these for future ideas
var CREATE_REPO_TAGS_TABLE_QUERY string = `
	CREATE TABLE IF NOT EXISTS repo_tags (
		repo_id TEXT NOT NULL,
		tag TEXT NOT NULL,
		PRIMARY KEY (repo_id, tag),
		FOREIGN KEY (repo_id) REFERENCES repos(id) ON DELETE CASCADE
		)`
var CREATE_REPO_METADATA_TABLE_QUERY string = `
		CREATE TABLE IF NOT EXISTS repo_metadata (
			repo_id TEXT NOT NULL,
			key TEXT NOT NULL,
			value TEXT NOT NULL,
			PRIMARY KEY (repo_id, key),
			FOREIGN KEY (repo_id) REFERENCES repos(id) ON DELETE CASCADE
			)`

func initDB() error {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	queries := []string{
		CREATE_NAMESPACE_TABLE_QUERY,
		CREATE_REPO_TABLE_QUERY,
		CREATE_REPO_TAGS_TABLE_QUERY,
		CREATE_REPO_METADATA_TABLE_QUERY,
		ADD_DEFAULT_NAMESPACE_QUERY,
	}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			db.Close()
			return fmt.Errorf("failed to execute query '%s': %w", query, err)
		}
	}
	return nil
}

func getDB() (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := initDB()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize database: %w", err)
		}
		log.Println("Database initialized successfully.")
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return db, nil
}
