package config

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Database struct {
		URI  string `yaml:"uri"`
		Name string `yaml:"name"`
	} `yaml:"database"`
	Server struct {
		Port           string `yaml:"port"`
		FrontendOrigin string `yaml:"frontend_origin"`
	} `yaml:"server"`
}

var AppConfig *Config

func LoadConfig() {
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = "config.yml"
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file (%s): %v", configPath, err)
	}

	AppConfig = &Config{}
	err = yaml.Unmarshal(file, AppConfig)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Override port with environment variable if present (for Cloud Run)
	if port := os.Getenv("PORT"); port != "" {
		if port[0] != ':' {
			AppConfig.Server.Port = ":" + port
		} else {
			AppConfig.Server.Port = port
		}
	}
}
