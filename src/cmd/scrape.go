package cmd

import (
	"github.com/spf13/cobra"
	"github.com/therealnoob/novelGo/config"
	"github.com/therealnoob/novelGo/scraper"
)

// NewRunCmd creates the `run` command
func scrapeCmd() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:   "scrape",
		Short: "One-shot scrape a novel",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.NewConfig(configFile)
			if err != nil {
				return err
			}

			if err := scraper.Scrape(cfg); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "path to config file")
	return cmd
}
