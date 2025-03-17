package config

import (
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type LoaderConfig struct {
	Path string
}

// NewConfigLoader создает загрузчик конфигурации с поиском пути:
// 1. Аргумент `-config` (если указан).z
// 2. Переменная окружения `CONFIG_PATH`.
// 3. `defaultPath`, если ничего не найдено.
func NewLoaderConfig(defaultPath string) *LoaderConfig {
	var path string
	flag.StringVar(&path, "config", "", "Path to the configuration file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	if path == "" {
		path = defaultPath
	}

	return &LoaderConfig{Path: path}
}

// Load загружает конфигурацию в переданную структуру
func (cl *LoaderConfig) Load(cfg interface{}) error {
	return LoadConfig(cl.Path, cfg)
}

// LoadConfig загружает конфигурацию из файла или переменной окружения
func LoadConfig(path string, cfg interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			envPath := os.Getenv("CONFIG_PATH")
			if envPath == "" {
				return fmt.Errorf("config not found: %q (CONFIG_PATH not set)", path)
			}
			data, err = os.ReadFile(envPath)
			if err != nil {
				return fmt.Errorf("config not found at %q or CONFIG_PATH (%q): %w", path, envPath, err)
			}
		} else {
			return fmt.Errorf("error reading config at %q: %w", path, err)
		}
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}
	return nil
}
