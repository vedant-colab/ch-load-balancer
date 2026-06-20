package config

import (
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port        int    `yaml:"port"`
	Environment string `yaml:"environment"`
}

func LoadConfig(log *zerolog.Logger) *Config {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal().Msgf("Error loading config file: %v", err)
		return nil
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal().Msgf("Failed to unmarshal config file: %v", err)
	}
	return &config
}
