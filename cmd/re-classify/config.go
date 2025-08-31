package main

import (
	"fmt"
	"os"

	"github.com/dlclark/regexp2"
	"gopkg.in/yaml.v3"
)

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
	// Regex patterns for identifier classification
	FormStartRegexp     []string `yaml:"form-start-regexp,omitempty"`
	FormEndRegexp       []string `yaml:"form-end-regexp,omitempty"`
	FormPrefixRegexp    []string `yaml:"form-prefix-regexp,omitempty"`
	SimpleLabelRegexp   []string `yaml:"simple-label-regexp,omitempty"`
	CompoundLabelRegexp []string `yaml:"compound-label-regexp,omitempty"`
	FormSurroundMatch   []string `yaml:"form-surround-match,omitempty"`

	// Operator configurations with precedence values
	OperatorRegexp []OperatorConfig `yaml:"operator-regexp,omitempty"`
}

// CompiledClassifierConfig holds compiled regex patterns
type CompiledClassifierConfig struct {
	FormStartRegexp     []*regexp2.Regexp
	FormEndRegexp       []*regexp2.Regexp
	FormPrefixRegexp    []*regexp2.Regexp
	SimpleLabelRegexp   []*regexp2.Regexp
	CompoundLabelRegexp []*regexp2.Regexp
	FormSurroundMatch   []*regexp2.Regexp
	OperatorConfigs     []CompiledOperatorConfig
}

// CompiledOperatorConfig holds a compiled operator configuration
type CompiledOperatorConfig struct {
	Pattern     *regexp2.Regexp
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

	// Compile form-start-regexp patterns
	compiled.FormStartRegexp, err = compileRegexpList(cc.FormStartRegexp, "form-start-regexp")
	if err != nil {
		return nil, err
	}

	// Compile form-end-regexp patterns
	compiled.FormEndRegexp, err = compileRegexpList(cc.FormEndRegexp, "form-end-regexp")
	if err != nil {
		return nil, err
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

	// Compile form-surround-match patterns
	compiled.FormSurroundMatch, err = compileRegexpList(cc.FormSurroundMatch, "form-surround-match")
	if err != nil {
		return nil, err
	}

	// Compile operator-regexp patterns
	for i, opConfig := range cc.OperatorRegexp {
		compiledOp := CompiledOperatorConfig{
			PrefixPrec:  opConfig.PrefixPrec,
			InfixPrec:   opConfig.InfixPrec,
			PostfixPrec: opConfig.PostfixPrec,
			EndTokens:   opConfig.EndTokens,
		}

		if opConfig.Pattern != "" {
			compiledOp.Pattern, err = regexp2.Compile("^"+opConfig.Pattern+"$", 0)
			if err != nil {
				return nil, fmt.Errorf("failed to compile operator-regexp pattern %d '%s': %w", i, opConfig.Pattern, err)
			}
		}

		compiled.OperatorConfigs = append(compiled.OperatorConfigs, compiledOp)
	}

	return compiled, nil
}

// compileRegexpList compiles a list of regex patterns with anchors
func compileRegexpList(patterns []string, fieldName string) ([]*regexp2.Regexp, error) {
	var compiled []*regexp2.Regexp

	for i, pattern := range patterns {
		if pattern == "" {
			continue
		}

		// Add anchors to ensure exact matching
		anchoredPattern := "^" + pattern + "$"
		regex, err := regexp2.Compile(anchoredPattern, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to compile %s pattern %d '%s': %w", fieldName, i, pattern, err)
		}
		compiled = append(compiled, regex)
	}

	return compiled, nil
}

// MatchesAny checks if the given text matches any of the compiled regexes
func MatchesAny(text string, regexes []*regexp2.Regexp) bool {
	for _, regex := range regexes {
		match, err := regex.MatchString(text)
		if err == nil && match {
			return true
		}
	}
	return false
}

// MatchesFormSurroundPattern checks if the given text matches any form-surround-match pattern
func (ccc *CompiledClassifierConfig) MatchesFormSurroundPattern(startToken, endToken string) bool {
	testString := startToken + " " + endToken
	return MatchesAny(testString, ccc.FormSurroundMatch)
}

// FindOperatorConfig returns the first matching operator configuration for the given text
func (ccc *CompiledClassifierConfig) FindOperatorConfig(text string) *CompiledOperatorConfig {
	for _, opConfig := range ccc.OperatorConfigs {
		if opConfig.Pattern != nil {
			match, err := opConfig.Pattern.MatchString(text)
			if err == nil && match {
				return &opConfig
			}
		}
	}
	return nil
}
