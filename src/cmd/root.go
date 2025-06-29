package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd sets up the Cobra command structure.
func Execute() error {
	// Root command (no direct execution)
	rootCmd := &cobra.Command{
		Use:   "novelGo",
		Short: "A web scraper for novels",
	}

	rootCmd.AddCommand(scrapeCmd(), configCmd())

	return rootCmd.Execute()
}
