// yaml.go functions for handling yaml configuration.

// Package config initializes and manages the application and pipeline configuration.
package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// AppConfigStruct config.yml is read into one of these on startup.
type AppConfigStruct struct {
	Param1 string `yaml:"param1"`
	Param2 string `yaml:"param2"`
	Param3 string `yaml:"param3"`
}

// loadAppConfig loads config.yml and returns nil, else returns an error.
func loadAppConfig() error {
	configFilePath := filepath.Join(Dir, "config.yml")

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the YAML file into a Config struct
	if err := yaml.Unmarshal(configFile, &AppConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
