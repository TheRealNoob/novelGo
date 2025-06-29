package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealnoob/novelGo/config"
)

// Config command
func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
	}
	cmd.AddCommand(configPrintCmd())
	return cmd
}

// Config print subcommand
func configPrintCmd() *cobra.Command {
	// arguments
	var configFile string

	cmd := &cobra.Command{
		Use:   "print",
		Short: "Print the configuration (full by default, user-defined with --user-defined)",

		RunE: func(cmd *cobra.Command, args []string) error {
			// parse config
			cfg, err := config.NewConfig("config.yaml")
			if err != nil {
				return err
			}

			fmt.Println("Effective Configuration:")
			fmt.Printf("%v\n", cfg)
			return nil
		},
	}

	cmd.Flags().StringVar(&configFile, "config", "config.yaml", "path to config file")

	return cmd
}
