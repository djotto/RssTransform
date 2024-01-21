// contract.go - This file contains all the exported members of the config
// package intended for public use. Other variables in the package may be
// exported, but are not intended for external consumption.

// Package config initializes and manages the application configuration.
package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// Dir holds the path to the configuration directory.
	// This variable can be set via command line flag "-config".
	Dir string
	// AppConfig loaded from {Dir}/config.yml
	AppConfig AppConfigStruct
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
	loadAppConfig()

	return nil
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
