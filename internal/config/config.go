package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	CurrentNamespace string `json:"current_namespace"`
}

const configPath = "/home/ahmed/.mgit/config.json"

func ensureConfigExists() error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{CurrentNamespace: "default"}
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

func GetCurrentNamespace() string {
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
	return cfg.CurrentNamespace
}

func SetCurrentNamespace(ns string) error {
	cfg := Config{CurrentNamespace: ns}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
