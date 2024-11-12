package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Loader struct {
	Path string
}

func NewLoader(envVar, defaultPath string) *Loader {
	var path string

	flag.StringVar(&path, "config", defaultPath, "Path to the configuration file")
	flag.Parse()

	if path == "" {
		path = os.Getenv(envVar)
	}

	if path == "" {
		path = defaultPath
	}

	return &Loader{Path: path}
}

func (l *Loader) Load(config interface{}) error {
	data, err := os.ReadFile(l.Path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return nil
}
