package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"server"`
	Database struct {
		DSN string `mapstructure:"DSN"`
	} `mapstructure:"database"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// чтение .env
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
