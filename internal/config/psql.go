package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string `yaml:"name"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func NewConfig(path string) (*Config, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	return &config, yaml.NewDecoder(f).Decode(&config)
}
