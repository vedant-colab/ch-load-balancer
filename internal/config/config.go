package config

import (
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server       ServerConfig    `yaml:"server"`
	Backends     []BackEndConfig `yaml:"backends"`
	VirtualNodes VirtualConfig   `yaml:"virtual_nodes"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
	// Environment string `yaml:"environment"`
}

type VirtualConfig struct {
	Total int `yaml:"total"`
}

type BackEndConfig struct {
	Id   string `yaml:"id"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
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
