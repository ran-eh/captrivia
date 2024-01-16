package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string        `mapstructure:"SERVER_ADDRESS"`
	ReadTimeout   time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout  time.Duration `mapstructure:"WRITE_TIMEOUT"`
	DatabaseDSN   string        `mapstructure:"DATABASE_DSN"`
}

// LoadConfig reads configuration from environment variables.
func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set defaults, which can be overridden by env
	v.SetDefault("SERVER_ADDRESS", ":8080")
	v.SetDefault("READ_TIMEOUT", 10*time.Second)
	v.SetDefault("WRITE_TIMEOUT", 10*time.Second)

	// Bind environment variables
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}