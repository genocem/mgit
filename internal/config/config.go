package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	CurrentProject string `json:"current_project"`
}

var mgitpath string

var configPath = filepath.Join(GetPath(), "config.json")

func GetPath() string {
	if mgitpath == "" {
		if runtime.GOOS == "windows" {
			d := os.Getenv("LOCALAPPDATA")
			if d == "" {
				log.Fatal("LOCALAPPDATA not set")
			}
			return filepath.Join(d, "mgit")
		} else {
			userpath, err := os.UserHomeDir()
			if err != nil {
				log.Fatal("no home directory path")
			}
			mgitpath = filepath.Join(userpath, ".mgit")
		}
	}
	return mgitpath
}

func ensureConfigExists() error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{CurrentProject: "default"}
		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return err
		}
	}
	return nil
}

func GetCurrentProject() string {
	if err := ensureConfigExists(); err != nil {
		return ""
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return ""
	}
	return cfg.CurrentProject
}

func SetCurrentProject(ns string) error {
	cfg := Config{CurrentProject: ns}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
