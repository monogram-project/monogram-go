package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sfkleach/regexptable"
	"gopkg.in/yaml.v3"
)

// SurroundRegexpConfig represents a start/end pair with regex substitution
type SurroundRegexpConfig struct {
	Start string   `yaml:"start"`
	End   []string `yaml:"end"`
}

// OperatorConfig represents operator configuration with three precedence values
type OperatorConfig struct {
	Pattern     string   `yaml:"pattern"`
	PrefixPrec  uint16   `yaml:"prefix-prec"`
	InfixPrec   uint16   `yaml:"infix-prec"`
	PostfixPrec uint16   `yaml:"postfix-prec"`
	EndTokens   []string `yaml:"end-tokens,omitempty"` // For form-start tokens
}

// ClassifierConfig represents the configuration structure for the re-classify tool
type ClassifierConfig struct {
	// New surround regex patterns with substitution
	SurroundRegexp []SurroundRegexpConfig `yaml:"surround-regexp,omitempty"`

	// Legacy regex patterns for identifier classification (for backward compatibility)
	FormStartRegexp     []string `yaml:"form-start-regexp,omitempty"`
	FormEndRegexp       []string `yaml:"form-end-regexp,omitempty"`
	FormPrefixRegexp    []string `yaml:"form-prefix-regexp,omitempty"`
	SimpleLabelRegexp   []string `yaml:"simple-label-regexp,omitempty"`
	CompoundLabelRegexp []string `yaml:"compound-label-regexp,omitempty"`
	FormSurroundMatch   []string `yaml:"form-surround-match,omitempty"`

	// Operator configurations with precedence values
	OperatorRegexp []OperatorConfig `yaml:"operator-regexp,omitempty"`
}

// CompiledSurroundRegexp holds a compiled surround regex configuration
type CompiledSurroundRegexp struct {
	StartPattern string   // Original pattern for reference
	EndSubsts    []string // End substitution patterns
}

// CompiledClassifierConfig holds compiled regex patterns
type CompiledClassifierConfig struct {
	// New efficient start token recognizer - maps start patterns to end substitution lists
	StartTokenTable   *regexptable.RegexpTable[[]string] // For quick lookup of valid end substitutions
	EndTokenTable     *regexptable.RegexpTable[bool]     // For quick lookup of valid end tokens
	FormSurroundMatch []*regexptable.RegexpTable[bool]   // Compiled form-surround-match patterns

	// Legacy compiled patterns (for backward compatibility)
	FormPrefixRegexp    []*regexp.Regexp
	SimpleLabelRegexp   []*regexp.Regexp
	CompoundLabelRegexp []*regexp.Regexp
	OperatorConfigs     []CompiledOperatorConfig
}

// CompiledOperatorConfig holds a compiled operator configuration
type CompiledOperatorConfig struct {
	Pattern     *regexp.Regexp
	PrefixPrec  uint16
	InfixPrec   uint16
	PostfixPrec uint16
	EndTokens   []string
}

// LoadClassifierConfig loads configuration from a YAML file
func LoadClassifierConfig(filename string) (*ClassifierConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", filename, err)
	}

	var config ClassifierConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", filename, err)
	}

	return &config, nil
}

// CompileRegexes compiles all regex patterns in the configuration
func (cc *ClassifierConfig) CompileRegexes() (*CompiledClassifierConfig, error) {
	compiled := &CompiledClassifierConfig{}
	var err error

	// Build start token recognizer using RegexpTableBuilder
	if len(cc.SurroundRegexp) > 0 {
		builder := regexptable.NewRegexpTableBuilder[[]string]()

		for _, surroundConfig := range cc.SurroundRegexp {
			if surroundConfig.Start != "" {
				// Add pattern to the table builder with the end substitutions as the value
				builder.AddPattern(surroundConfig.Start, surroundConfig.End)
			}
		}

		// Build the table with exact matching (anchored)
		compiled.StartTokenTable, err = builder.Build(true, true)
		if err != nil {
			return nil, fmt.Errorf("failed to build start token table: %w", err)
		}
	}

	// Compile form-prefix-regexp patterns
	compiled.FormPrefixRegexp, err = compileRegexpList(cc.FormPrefixRegexp, "form-prefix-regexp")
	if err != nil {
		return nil, err
	}

	// Compile simple-label-regexp patterns
	compiled.SimpleLabelRegexp, err = compileRegexpList(cc.SimpleLabelRegexp, "simple-label-regexp")
	if err != nil {
		return nil, err
	}

	// Compile compound-label-regexp patterns
	compiled.CompoundLabelRegexp, err = compileRegexpList(cc.CompoundLabelRegexp, "compound-label-regexp")
	if err != nil {
		return nil, err
	}

	// Note: FormSurroundMatch will be populated dynamically during token analysis

	// Compile operator-regexp patterns
	for i, opConfig := range cc.OperatorRegexp {
		compiledOp := CompiledOperatorConfig{
			PrefixPrec:  opConfig.PrefixPrec,
			InfixPrec:   opConfig.InfixPrec,
			PostfixPrec: opConfig.PostfixPrec,
			EndTokens:   opConfig.EndTokens,
		}

		if opConfig.Pattern != "" {
			compiledOp.Pattern, err = regexp.Compile("^" + opConfig.Pattern + "$")
			if err != nil {
				return nil, fmt.Errorf("failed to compile operator-regexp pattern %d '%s': %w", i, opConfig.Pattern, err)
			}
		}

		compiled.OperatorConfigs = append(compiled.OperatorConfigs, compiledOp)
	}

	return compiled, nil
}

// simpleSubstitute performs simple $0 substitution
func simpleSubstitute(pattern, matchText string) string {
	return strings.ReplaceAll(pattern, "$0", matchText)
}

// compileRegexpList compiles a list of regex patterns with anchors
func compileRegexpList(patterns []string, fieldName string) ([]*regexp.Regexp, error) {
	var compiled []*regexp.Regexp

	for i, pattern := range patterns {
		if pattern == "" {
			continue
		}

		// Add anchors to ensure exact matching
		anchoredPattern := "^" + pattern + "$"
		regex, err := regexp.Compile(anchoredPattern)
		if err != nil {
			return nil, fmt.Errorf("failed to compile %s pattern %d '%s': %w", fieldName, i, pattern, err)
		}
		compiled = append(compiled, regex)
	}

	return compiled, nil
}

// MatchesAny checks if the given text matches any of the compiled regexes
func MatchesAny(text string, regexes []*regexp.Regexp) bool {
	for _, regex := range regexes {
		if regex.MatchString(text) {
			return true
		}
	}
	return false
}

// MatchesFormSurroundPattern checks if the given text matches any form-surround-match pattern
func (ccc *CompiledClassifierConfig) MatchesFormSurroundPattern(startToken, endToken string) bool {
	testString := startToken + " " + endToken
	for _, table := range ccc.FormSurroundMatch {
		if table != nil {
			_, _, ok := table.TryLookup(testString)
			if ok {
				return true
			}
		}
	}
	return false
}

// FindOperatorConfig returns the first matching operator configuration for the given text
func (ccc *CompiledClassifierConfig) FindOperatorConfig(text string) *CompiledOperatorConfig {
	for _, opConfig := range ccc.OperatorConfigs {
		if opConfig.Pattern != nil {
			if opConfig.Pattern.MatchString(text) {
				return &opConfig
			}
		}
	}
	return nil
}
