package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ClassifierEngine implements the token classification logic
type ClassifierEngine struct {
	config          *CompiledClassifierConfig
	tokens          []string
	formStartEndMap map[string][]string // maps form-start tokens to their possible end tokens
}

// NewClassifierEngine creates a new classifier engine with the given configuration
func NewClassifierEngine(config *CompiledClassifierConfig) *ClassifierEngine {
	return &ClassifierEngine{
		config:          config,
		formStartEndMap: make(map[string][]string),
	}
}

// BuildFormStartEndMappings analyzes all tokens and builds the mapping from form-start to form-end tokens
func (ce *ClassifierEngine) BuildFormStartEndMappings(tokens []string) {
	// Collect all potential form-start and form-end tokens from the input
	potentialStarts := make(map[string]bool)
	potentialEnds := make(map[string]bool)

	for _, token := range tokens {
		if MatchesAny(token, ce.config.FormStartRegexp) {
			potentialStarts[token] = true
		}
		if MatchesAny(token, ce.config.FormEndRegexp) {
			potentialEnds[token] = true
		}
	}

	// Use form-surround-match patterns to build explicit start->end mappings
	for startToken := range potentialStarts {
		var endTokens []string
		for endToken := range potentialEnds {
			if ce.config.MatchesFormSurroundPattern(startToken, endToken) {
				endTokens = append(endTokens, endToken)
			}
		}
		if len(endTokens) > 0 {
			ce.formStartEndMap[startToken] = endTokens
		}
	}
}

// ClassifyToken classifies a single token and returns the classification string
func (ce *ClassifierEngine) ClassifyToken(token string) string {
	// Check compound label first (highest priority)
	if MatchesAny(token, ce.config.CompoundLabelRegexp) {
		return "C"
	}

	// Check simple label
	if MatchesAny(token, ce.config.SimpleLabelRegexp) {
		return "L"
	}

	// Check form prefix
	if MatchesAny(token, ce.config.FormPrefixRegexp) {
		return "P"
	}

	// Check form end
	if MatchesAny(token, ce.config.FormEndRegexp) {
		return "E"
	}

	// Check form start (and include explicit end tokens)
	if MatchesAny(token, ce.config.FormStartRegexp) {
		if endTokens, exists := ce.formStartEndMap[token]; exists {
			return "S " + strings.Join(endTokens, " ")
		}
		return "S" // No matching end tokens found
	}

	// Check operator
	if opConfig := ce.config.FindOperatorConfig(token); opConfig != nil {
		return fmt.Sprintf("O %d %d %d", opConfig.PrefixPrec, opConfig.InfixPrec, opConfig.PostfixPrec)
	}

	// Default to variable
	return "V"
}

// ProcessTokens processes all tokens and outputs classifications
func (ce *ClassifierEngine) ProcessTokens(tokens []string) {
	for _, token := range tokens {
		classification := ce.ClassifyToken(token)
		fmt.Println(classification)
	}
}

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config.yaml>\n", os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]

	// Load configuration
	config, err := LoadClassifierConfig(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Compile regex patterns
	compiledConfig, err := config.CompileRegexes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling regexes: %v\n", err)
		os.Exit(1)
	}

	// Create classifier engine
	engine := NewClassifierEngine(compiledConfig)

	// Read tokens from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var tokens []string

	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		if token != "" {
			tokens = append(tokens, token)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	// Build form-start to form-end mappings by analyzing all tokens
	engine.BuildFormStartEndMappings(tokens)

	// Process tokens and output classifications
	engine.ProcessTokens(tokens)
}
