package sqlctidy

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config is a struct to hold the contents of the configuration file.
type Config struct {
	Version string `yaml:"version"`
	SQL     []struct {
		Schema  string `yaml:"schema"`
		Queries string `yaml:"queries"`
		Engine  string `yaml:"engine"`
		Gen     struct {
			Go struct {
				Package string `yaml:"package"`
				Out     string `yaml:"out"`
			} `yaml:"go"`
		} `yaml:"gen"`
	} `yaml:"sql"`
}

// ReadConfig reads the configuration file from the specified path and stores it in the Config struct.
func ReadConfig(dir string) (Config, error) {
	var config Config
	path, err := ConfigPath(dir)
	if err != nil {
		return config, fmt.Errorf("cannot get config path: %w", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("cannot open file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse YAML: %w", err)
	}
	return config, nil
}

// Returns the path of the configuration file.
// The configuration file could be sqlc.yml, .yaml, or .json,
// so it returns the path by specifying the file name.
func ConfigPath(dir string) (string, error) {
	// Check if the current directory has sqlc.yml
	if _, err := os.Stat(filepath.Join(dir, "sqlc.yml")); err == nil {
		return filepath.Join(dir, "sqlc.yml"), nil
	}
	// Check if the current directory has sqlc.yaml
	if _, err := os.Stat(filepath.Join(dir, "sqlc.yaml")); err == nil {
		return filepath.Join(dir, "sqlc.yaml"), nil
	}
	// Check if the current directory has sqlc.json
	if _, err := os.Stat(filepath.Join(dir, "sqlc.json")); err == nil {
		return filepath.Join(dir, "sqlc.json"), nil
	}
	// If no configuration file is found, return an error
	return "", fmt.Errorf("no config file found")
}
