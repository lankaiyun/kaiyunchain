package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(consoleCmd)
	rootCmd.AddCommand(runCmd)
}

var rootCmd = &cobra.Command{
	Use:   "kayo",
	Short: "The kaiyunchain command line interface",
	Long:  "Kayo is a command-line tool implementation of kaiyunchain.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
