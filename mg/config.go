package mg

import (
	"fmt"
	"os"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/sfkleach/regexptable"
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
	// Regex patterns for identifier classification (arrays of strings from YAML)
	FormStartRegex     []string `yaml:"form-start-regex,omitempty"`
	FormEndRegex       []string `yaml:"form-end-regex,omitempty"`
	FormPrefixRegex    []string `yaml:"form-prefix-regex,omitempty"`
	SimpleLabelRegex   []string `yaml:"simple-label-regex,omitempty"`
	CompoundLabelRegex []string `yaml:"compound-label-regex,omitempty"`
	FormSurroundMatch  []string `yaml:"form-surround-match,omitempty"`
}

type TokenClassifiersCompiled struct {
	// Compiled regex tables using regexptable
	FormStartRegexTable       *regexptable.RegexpTable[bool]
	FormEndRegexTable         *regexptable.RegexpTable[bool]
	FormPrefixRegexTable      *regexptable.RegexpTable[bool]
	SimpleLabelRegexTable     *regexptable.RegexpTable[bool]
	CompoundLabelRegexTable   *regexptable.RegexpTable[bool]
	FormSurroundMatchCompiled *regexp2.Regexp
}

// MatchesFormSurroundPattern checks if the given text matches any of the form-surround-match patterns
func (tcc *TokenClassifiersCompiled) MatchesFormSurroundPattern(text string) bool {
	if tcc.FormSurroundMatchCompiled == nil {
		return false
	}
	matched, err := tcc.FormSurroundMatchCompiled.MatchString(text)
	return err == nil && matched
}

// CheckFormStartEndMatch checks if a form-start and form-end token pair match
// by constructing "start end" and checking against the form-surround-match patterns
func (tcc *TokenClassifiersCompiled) CheckFormStartEndMatch(startToken, endToken string) bool {
	testString := startToken + " " + endToken
	return tcc.MatchesFormSurroundPattern(testString)
}

// MatchesFormStart checks if the identifier matches any form-start patterns
func (tcc *TokenClassifiersCompiled) MatchesFormStart(identifier string) bool {
	if tcc.FormStartRegexTable == nil {
		return false
	}
	_, _, ok := tcc.FormStartRegexTable.TryLookup(identifier)
	return ok
}

// MatchesFormEnd checks if the identifier matches any form-end patterns
func (tcc *TokenClassifiersCompiled) MatchesFormEnd(identifier string) bool {
	if tcc.FormEndRegexTable == nil {
		return false
	}
	_, _, ok := tcc.FormEndRegexTable.TryLookup(identifier)
	return ok
}

// MatchesFormPrefix checks if the identifier matches any form-prefix patterns
func (tcc *TokenClassifiersCompiled) MatchesFormPrefix(identifier string) bool {
	if tcc.FormPrefixRegexTable == nil {
		return false
	}
	_, _, ok := tcc.FormPrefixRegexTable.TryLookup(identifier)
	return ok
}

// MatchesSimpleLabel checks if the identifier matches any simple-label patterns
func (tcc *TokenClassifiersCompiled) MatchesSimpleLabel(identifier string) bool {
	if tcc.SimpleLabelRegexTable == nil {
		return false
	}
	_, _, ok := tcc.SimpleLabelRegexTable.TryLookup(identifier)
	return ok
}

// MatchesCompoundLabel checks if the identifier matches any compound-label patterns
func (tcc *TokenClassifiersCompiled) MatchesCompoundLabel(identifier string) bool {
	if tcc.CompoundLabelRegexTable == nil {
		return false
	}
	_, _, ok := tcc.CompoundLabelRegexTable.TryLookup(identifier)
	return ok
}

// CompileRegexes converts TokenClassifiers to TokenClassifiersCompiled using RegexpTableBuilder
func (tc *TokenClassifiers) CompileRegexes() (*TokenClassifiersCompiled, error) {
	compiled := &TokenClassifiersCompiled{}
	var err error

	// Build FormStartRegexTable
	if len(tc.FormStartRegex) > 0 {
		builder := regexptable.NewRegexpTableBuilder[bool]()
		for _, pattern := range tc.FormStartRegex {
			if pattern != "" {
				builder.AddPattern(pattern, true)
			}
		}
		compiled.FormStartRegexTable, err = builder.Build(true, true) // Exact matching
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-start-regex patterns: %w", err)
		}
	}

	// Build FormEndRegexTable
	if len(tc.FormEndRegex) > 0 {
		builder := regexptable.NewRegexpTableBuilder[bool]()
		for _, pattern := range tc.FormEndRegex {
			if pattern != "" {
				builder.AddPattern(pattern, true)
			}
		}
		compiled.FormEndRegexTable, err = builder.Build(true, true) // Exact matching
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-end-regex patterns: %w", err)
		}
	}

	// Build FormPrefixRegexTable
	if len(tc.FormPrefixRegex) > 0 {
		builder := regexptable.NewRegexpTableBuilder[bool]()
		for _, pattern := range tc.FormPrefixRegex {
			if pattern != "" {
				builder.AddPattern(pattern, true)
			}
		}
		compiled.FormPrefixRegexTable, err = builder.Build(true, true) // Exact matching
		if err != nil {
			return nil, fmt.Errorf("failed to compile form-prefix-regex patterns: %w", err)
		}
	}

	// Build SimpleLabelRegexTable
	if len(tc.SimpleLabelRegex) > 0 {
		builder := regexptable.NewRegexpTableBuilder[bool]()
		for _, pattern := range tc.SimpleLabelRegex {
			if pattern != "" {
				builder.AddPattern(pattern, true)
			}
		}
		compiled.SimpleLabelRegexTable, err = builder.Build(true, true) // Exact matching
		if err != nil {
			return nil, fmt.Errorf("failed to compile simple-label-regex patterns: %w", err)
		}
	}

	// Build CompoundLabelRegexTable
	if len(tc.CompoundLabelRegex) > 0 {
		builder := regexptable.NewRegexpTableBuilder[bool]()
		for _, pattern := range tc.CompoundLabelRegex {
			if pattern != "" {
				builder.AddPattern(pattern, true)
			}
		}
		compiled.CompoundLabelRegexTable, err = builder.Build(true, true) // Exact matching
		if err != nil {
			return nil, fmt.Errorf("failed to compile compound-label-regex patterns: %w", err)
		}
	}

	// Handle FormSurroundMatch (unchanged for now)
	if len(tc.FormSurroundMatch) > 0 {
		// Build a single alternation pattern from all the start/end pairs
		var alternatives []string
		for i, pattern := range tc.FormSurroundMatch {
			if pattern == "" {
				continue
			}

			// Validate that the pattern has the expected format "start end"
			parts := strings.SplitN(pattern, " ", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("form-surround-match pattern %d '%s' must have format 'start_regex end_regex'", i, pattern)
			}

			startPattern := strings.TrimSpace(parts[0])
			endPattern := strings.TrimSpace(parts[1])

			if startPattern == "" || endPattern == "" {
				return nil, fmt.Errorf("form-surround-match pattern %d '%s' has empty start or end regex", i, pattern)
			}

			// Use the pattern directly as it already has the format "start end"
			alternatives = append(alternatives, pattern)
		}

		if len(alternatives) > 0 {
			// Combine all alternatives into a single pattern and wrap for exact matching
			combinedPattern := "^(?:" + strings.Join(alternatives, "|") + ")$"
			compiled.FormSurroundMatchCompiled, err = regexp2.Compile(combinedPattern, regexp2.None)
			if err != nil {
				return nil, fmt.Errorf("failed to compile form-surround-match patterns '%s': %w", combinedPattern, err)
			}
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
