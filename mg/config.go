package mg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigurableOptions struct {
	Format            string `yaml:"option-format,omitempty"`
	Indent            int    `yaml:"option-indent,omitempty"`
	DefaultLabel      string `yaml:"option-default-label,omitempty"`
	IncludeSpans      bool   `yaml:"option-include-spans,omitempty"`
	Decimal           bool   `yaml:"option-decimal,omitempty"`
	CheckLiterals     bool   `yaml:"option-check-literals,omitempty"`
	UseClassifier     string `yaml:"option-use-classifier,omitempty"`
	TrimTokenOnOutput int    `yaml:"option-trim-token-on-output,omitempty"`
}

// FormatOptions represents the formatting options
type FormatOptions struct {
	Input  string
	Output string
	Limit  bool
	ConfigurableOptions
}

// Config represents the configuration structure that can be loaded from YAML.
// At present this is the same as the CoreFormatOptions but may in the future
// include additional fields.
type Config struct {
	// Default options that can be overridden by command line
	ConfigurableOptions `yaml:",inline"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	if filename == "" {
		return &Config{}, nil // Return empty config if no file specified
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", filename, err)
	}

	return &config, nil
}

// ApplyConfigDefaults applies configuration defaults to FormatOptions, but only
// for fields that haven't been explicitly set via command line flags
func (c *Config) ApplyConfigDefaults(options *FormatOptions, flagsExplicitlySet map[string]bool) {
	if !flagsExplicitlySet["format"] && c.Format != "" {
		options.Format = c.Format
	}
	if !flagsExplicitlySet["indent"] && c.Indent > 0 {
		options.Indent = c.Indent
	}
	if !flagsExplicitlySet["default-label"] && c.DefaultLabel != "" {
		options.DefaultLabel = c.DefaultLabel
	}
	if !flagsExplicitlySet["include-spans"] && c.IncludeSpans {
		options.IncludeSpans = c.IncludeSpans
	}
	if !flagsExplicitlySet["decimal"] && c.Decimal {
		options.Decimal = c.Decimal
	}
	if !flagsExplicitlySet["check-literals"] && c.CheckLiterals {
		options.CheckLiterals = c.CheckLiterals
	}
	if !flagsExplicitlySet["use-classifier"] && c.UseClassifier != "" {
		options.UseClassifier = c.UseClassifier
	}
}
