// contract.go - This file contains all the exported members of the config
// package intended for public use. Other variables in the package may be
// exported, but are not intended for external consumption.

// Package config initializes and manages the application configuration.
package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

// PipelineConfigStruct contains the configuration of a single pipeline.
type PipelineConfigStruct struct {
	SleepDuration int    `yaml:"SleepDuration"`
	Param1        string `yaml:"param1"`
	Param2        string `yaml:"param2"`
	Param3        string `yaml:"param3"`
}

var (
	// Dir holds the path to the configuration directory.
	// This variable can be set via command line flag "-config".
	Dir string
	// AppConfig loaded from {Dir}/config.yml
	AppConfig AppConfigStruct
	// PipelineConfigs loaded from {Dir}/*.yml
	PipelineConfigs []PipelineConfigStruct
)

// Init uses Cobra to get command-line arguments.
// Returns an error if there was a problem, else nil.
func Init() error {
	// Initialize the Cobra command
	var rootCmd = &cobra.Command{
		Use:   "rss-transform",
		Short: "ETL pipeline for RSS feeds.",
		Long:  `Consume, transform and republish RSS feeds.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// This function will be called when the root command runs
			return InitConfig()
		},
	}

	// Define flags
	rootCmd.PersistentFlags().StringVar(&Dir, "config", "/etc/rss-transform", "config directory")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("failed to execute root command: %w", err)
	}

	// Get Dir/config.yaml and process it.
	appConfigPtr, err := loadAppConfig(Dir)
	if err != nil {
		return fmt.Errorf("config::Init(): %w", err)
	}
	appConfig := *appConfigPtr

	pipelineConfigs, err := loadPipelineConfigs(Dir)
	if err != nil {
		return fmt.Errorf("config::Init(): %w", err)
	}

	// Assign to package-level variables
	AppConfig = appConfig
	PipelineConfigs = pipelineConfigs
	return nil
}
