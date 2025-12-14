package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadConfig reads and parses the octohook.yml file from the given path.
// Returns the parsed Config struct or an error if the file cannot be read
// or contains invalid YAML syntax.
//
// Example:
//
//	cfg, err := config.LoadConfig("octohook.yml")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(cfg.PreCommit[0].Command)

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
