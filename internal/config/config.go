package config

import (
	"os"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

type Config struct {
	Service Service `yaml:"service"`
	Redis   Redis   `yaml:"redis"`
}

func New(configPath string) (*Config, error) {
	b, err := readFile(configPath)
	if err != nil {
		return nil, err
	}

	var out Config
	err = yaml.Unmarshal(b, &out)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal bytes to config")
	}

	return &out, nil
}

func readFile(configPath string) ([]byte, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read file with configuration")
	}

	return b, nil
}
