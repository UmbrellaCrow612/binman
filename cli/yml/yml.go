package yml

import (
	"os"

	"github.com/UmbrellaCrow612/binman/cli/args"
	"github.com/UmbrellaCrow612/binman/cli/printer"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Binaries []Binary `yaml:"binaries"`
}

type Binary struct {
	Name  string            `yaml:"name"`
	URL   map[string]string `yaml:"url"`
	Extra map[string]string `yaml:",inline"`
}

// parse reads the YAML file from opts.PathToFile and returns the parsed Config
func Parse(opts *args.Options) *Config {
	if opts.PathToFile == "" {
		printer.ExitError("PathToFile is empty")
	}

	data, err := os.ReadFile(opts.PathToFile)
	if err != nil {
		printer.ExitError("Failed to read file: " + err.Error())
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		printer.ExitError("Failed to parse YAML: " + err.Error())
	}

	printer.PrintSuccess("YAML file parsed successfully")
	return &cfg
}
