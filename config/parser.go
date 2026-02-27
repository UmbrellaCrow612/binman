package config

import (
	"os"

	"github.com/UmbrellaCrow612/binman/t"
	"github.com/goccy/go-yaml"
)

// parses the file path and trys to convert it to config struct
func Parse(filePath string) (*t.Config, error) {
	var config = &t.Config{}

	_, err := os.Stat(filePath)
	if err != nil {
		return config, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return config, err
	}

	return config, nil
}
