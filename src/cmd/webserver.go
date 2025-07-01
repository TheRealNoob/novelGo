package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealnoob/novelGo/config"
)

// NewRunCmd creates the `run` command
func webserverCmd() *cobra.Command {
	var configFile string

	cmd := &cobra.Command{
		Use:   "webserver",
		Short: "Starts a webserver for interacting with novelGo.  Useful for long-term scraping.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.NewConfig(configFile)
			if err != nil {
				return err
			}

			fmt.Println("starting webserver...")
			fmt.Printf("URL: %s\n", cfg.URL)
			return nil
		},
	}

	cmd.Flags().StringVar(&configFile, "config", "config.yaml", "path to config file")
	return cmd
}
