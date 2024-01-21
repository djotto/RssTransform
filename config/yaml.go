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

type PipelineConfigStruct struct {
	Param1 string `yaml:"param1"`
	Param2 string `yaml:"param2"`
	Param3 string `yaml:"param3"`
}

// loadAppConfig loads config.yml into AppConfig and returns nil, else returns an error.
func loadAppConfig(dir string) (*AppConfigStruct, error) {
	configFilePath := filepath.Join(Dir, "config.yml")

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("loadAppConfig(): %w", err)
	}

	var config AppConfigStruct

	// Unmarshal the YAML file into a Config struct
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("loadAppConfig(): %w", err)
	}

	return &config, nil
}

// loadPipelineConfigs loads *.yml into PipelineConfigs and returns nil, else returns an error.
func loadPipelineConfigs(dir string) ([]PipelineConfigStruct, error) {
	var configs []PipelineConfigStruct

	filePaths, err := getPipelineFilenames(dir)
	if err != nil {
		return nil, err
	}

	for _, filePath := range filePaths {
		configPtr, err := loadPipelineConfig(filePath)
		if err != nil {
			return nil, fmt.Errorf("loadPipelineConfigs(): %w", err)
		}

		// Append the loaded config to the slice
		configs = append(configs, *configPtr)
	}

	return configs, nil
}

// loadPipelineConfig loads a *.yml file into a PipelineConfigStruct and returns it, else error
func loadPipelineConfig(filePath string) (*PipelineConfigStruct, error) {
	var pipelineConfig PipelineConfigStruct

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Unmarshal the YAML data into the struct
	if err := yaml.Unmarshal(data, &pipelineConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config in file %s: %w", filePath, err)
	}

	return &pipelineConfig, nil
}

// getPipelineFilenames returns a slice of paths to all *.yml files in the specified directory,
// excluding config.yml.
func getPipelineFilenames(dir string) ([]string, error) {
	// Find all .yml files in the directory
	files, err := filepath.Glob(filepath.Join(dir, "*.yml"))
	if err != nil {
		return nil, err
	}

	// Filter out config.yml
	var result []string
	for _, file := range files {
		if filepath.Base(file) != "config.yml" {
			result = append(result, file)
		}
	}

	return result, nil
}

// InitConfig checks the availability of the configuration directory.
func InitConfig() error {
	// Check if the config directory is set and exists
	if Dir == "" {
		Dir = "/etc/rss-transform" // Set the default
	}

	if _, err := os.Stat(Dir); os.IsNotExist(err) {
		// Handle the case where the directory does not exist
		_, _ = fmt.Fprintf(os.Stderr, "Config directory '%s' does not exist\n", Dir)
		os.Exit(1)
	}

	return nil
}
