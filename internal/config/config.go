package config

import (
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

const configFilePath = ".config/gitlab-auto-mr/config.yml"

type Config struct {
	GitConstraints struct {
		ExpectedOriginHost string `yaml:"expectedOriginHost" validate:"required"`
	} `yaml:"gitConstraints"`
}

func Load() (*Config, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig() (*Config, error) {
	homeDir := os.Getenv("HOME")

	file, err := os.Open(filepath.Join(homeDir, configFilePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	validate := validator.New()
	return validate.Struct(config)
}
