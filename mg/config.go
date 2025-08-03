package mg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// FormatOptions represents the formatting options
type FormatOptions struct {
	Format        string
	Input         string
	Output        string
	Indent        int
	Limit         bool
	DefaultLabel  string
	IncludeSpans  bool
	Decimal       bool
	CheckLiterals bool
}

// Config represents the configuration structure that can be loaded from YAML
type Config struct {
	// Regex patterns for identifier classification
	FormEndRegex       string `yaml:"form-end-regex,omitempty"`
	FormStartRegex     string `yaml:"form-start-regex,omitempty"`
	FormPrefixRegex    string `yaml:"form-prefix-regex,omitempty"`
	SimpleLabelRegex   string `yaml:"simple-label-regex,omitempty"`
	CompoundLabelRegex string `yaml:"compound-label-regex,omitempty"`

	// Default options that can be overridden by command line
	DefaultFormat string `yaml:"option-format,omitempty"`
	DefaultIndent int    `yaml:"option-indent,omitempty"`
	DefaultLabel  string `yaml:"option-default-label,omitempty"`
	IncludeSpans  bool   `yaml:"option-include-spans,omitempty"`
	Decimal       bool   `yaml:"option-decimal,omitempty"`
	CheckLiterals bool   `yaml:"option-check-literals,omitempty"`
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
	if !flagsExplicitlySet["format"] && c.DefaultFormat != "" {
		options.Format = c.DefaultFormat
	}
	if !flagsExplicitlySet["indent"] && c.DefaultIndent > 0 {
		options.Indent = c.DefaultIndent
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
}
