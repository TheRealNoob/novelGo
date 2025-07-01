package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	URL string `mapstructure:"url"`
}

func NewConfig(configFile string) (*ConfigStruct, error) {
	v := viper.New()

	// Set default values
	v.SetDefault("url", "")

	// Bind environment variables with prefix "NOVELGO"
	v.SetEnvPrefix("NOVELGO")
	v.AutomaticEnv()

	// Check if config file exists before reading
	if _, err := os.Stat(configFile); err == nil {
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			return nil, errors.Join(errors.New("failed to read config file"), err)
		}
	}

	var cfg ConfigStruct
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, errors.Join(errors.New("failed to parse config file"), err)
	}

	// Input validation
	if cfg.URL == "" {
		return nil, errors.New("url is required")
	}

	return &cfg, nil
}
