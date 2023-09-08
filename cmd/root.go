package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP Tracker CLI SHORT",
		Long:  `CIP Tracker CLI LOng...`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
