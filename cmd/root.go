package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd defines generic description of app for CLI
var rootCmd = &cobra.Command{
	Use:   "geolocation",
	Short: "Geolocation loads data to database from CSV files and provides geo location data",
	Long: `Geolocation behaviors:
                - Load geo locations from CSV
                - Provide geo location via RPC`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
