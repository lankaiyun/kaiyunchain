package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kayo",
	Long:  "All software has versions. This is kayo's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kayo")
		fmt.Println("Version: 1.0.0-stable")
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("Operating System: %s\n", runtime.GOOS)
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Println()
	},
}
