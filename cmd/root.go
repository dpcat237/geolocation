package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd defines generic description of app for CLI
var rootCmd = &cobra.Command{
	Use:   "geolocation",
	Short: "Geouser load data to database from CSV files and provide user data",
	Long: `Geouser behavior:
                - Load users sessions from CSV
                - Provide users sessions via RPC`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
