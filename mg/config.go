package mg

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type CoreFormatOptions struct {
	Format        string `yaml:"option-format,omitempty"`
	Indent        int    `yaml:"option-indent,omitempty"`
	DefaultLabel  string `yaml:"option-default-label,omitempty"`
	IncludeSpans  bool   `yaml:"option-include-spans,omitempty"`
	Decimal       bool   `yaml:"option-decimal,omitempty"`
	CheckLiterals bool   `yaml:"option-check-literals,omitempty"`
}

// FormatOptions represents the formatting options
type FormatOptions struct {
	Input  string
	Output string
	Limit  bool
	CoreFormatOptions
}

type TokenClassifiers struct {
	// Regex patterns for identifier classification (raw strings from YAML)
	FormStartRegex       *string `yaml:"form-start-regex,omitempty"`
	FormEndRegex         *string `yaml:"form-end-regex,omitempty"`
	FormEndWildcardRegex *string `yaml:"form-end-wildcard-regex,omitempty"`
	FormPrefixRegex      *string `yaml:"form-prefix-regex,omitempty"`
	SimpleLabelRegex     *string `yaml:"simple-label-regex,omitempty"`
	CompoundLabelRegex   *string `yaml:"compound-label-regex,omitempty"`
}

type TokenClassifiersCompiled struct {
	// Compiled regex patterns
	FormStartRegexCompiled       *regexp.Regexp
	FormEndRegexCompiled         *regexp.Regexp
	FormEndWildcardRegexCompiled *regexp.Regexp
	FormPrefixRegexCompiled      *regexp.Regexp
	SimpleLabelRegexCompiled     *regexp.Regexp
	CompoundLabelRegexCompiled   *regexp.Regexp
}

// wrapForExactMatch wraps a regex pattern for exact matching
func wrapForExactMatch(pattern string) string {
	// Check if pattern is already anchored at both ends.
	if strings.HasPrefix(pattern, "^(?:") && strings.HasSuffix(pattern, ")$") {
		return pattern // Already anchored, don't modify
	}
	return "^(?:" + pattern + ")$"
}

// CompileRegexes converts TokenClassifiers to TokenClassifiersCompiled
func (tc *TokenClassifiers) CompileRegexes() (*TokenClassifiersCompiled, error) {
	compiled := &TokenClassifiersCompiled{}
	var err error

	if tc.FormStartRegex != nil && *tc.FormStartRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.FormStartRegex)
		compiled.FormStartRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-start-regex '%s': %w", *tc.FormStartRegex, err)
		}
	}

	if tc.FormEndRegex != nil && *tc.FormEndRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.FormEndRegex)
		compiled.FormEndRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-end-regex '%s': %w", *tc.FormEndRegex, err)
		}
	}

	if tc.FormEndWildcardRegex != nil && *tc.FormEndWildcardRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.FormEndWildcardRegex)
		compiled.FormEndWildcardRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-end-wildcard-regex '%s': %w", *tc.FormEndWildcardRegex, err)
		}
	}

	if tc.FormPrefixRegex != nil && *tc.FormPrefixRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.FormPrefixRegex)
		compiled.FormPrefixRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-prefix-regex '%s': %w", *tc.FormPrefixRegex, err)
		}
	}

	if tc.SimpleLabelRegex != nil && *tc.SimpleLabelRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.SimpleLabelRegex)
		compiled.SimpleLabelRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile simple-label-regex '%s': %w", *tc.SimpleLabelRegex, err)
		}
	}

	if tc.CompoundLabelRegex != nil && *tc.CompoundLabelRegex != "" {
		wrappedPattern := wrapForExactMatch(*tc.CompoundLabelRegex)
		compiled.CompoundLabelRegexCompiled, err = regexp.Compile(wrappedPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile compound-label-regex '%s': %w", *tc.CompoundLabelRegex, err)
		}
	}

	return compiled, nil
}

// Config represents the configuration structure that can be loaded from YAML
type Config struct {
	// Regex patterns for identifier classification
	TokenClassifiers `yaml:",inline"`

	// Compiled regex patterns (populated after loading)
	CompiledClassifiers *TokenClassifiersCompiled `yaml:"-"`

	// Default options that can be overridden by command line
	CoreFormatOptions `yaml:",inline"`
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

	// Compile regex patterns
	compiled, err := config.TokenClassifiers.CompileRegexes()
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex patterns: %w", err)
	}
	config.CompiledClassifiers = compiled

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
}
