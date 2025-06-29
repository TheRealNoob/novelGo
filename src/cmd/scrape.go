package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealnoob/novelGo/config"
)

// NewRunCmd creates the `run` command
func scrapeCmd() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:   "scrape",
		Short: "Run the web scraper",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("novelGo: Running scraper")
			cfg, err := config.NewConfig(configFile)
			if err != nil {
				return err
			}

			cmd.Printf("URL: %s", cfg.URL)
			return nil
		},
	}

	cmd.Flags().StringVar(&configFile, "config", "config.yaml", "path to config file")
	return cmd
}
