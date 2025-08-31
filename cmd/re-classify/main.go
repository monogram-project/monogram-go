package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sfkleach/regexptable"
)

// ClassifierEngine implements the token classification logic
type ClassifierEngine struct {
	config *CompiledClassifierConfig
}

// NewClassifierEngine creates a new classifier engine with the given configuration
func NewClassifierEngine(config *CompiledClassifierConfig) *ClassifierEngine {
	return &ClassifierEngine{
		config: config,
	}
}

// hasSubstitutionVariable checks if a pattern contains $\d+ substitution variables
func hasSubstitutionVariable(pattern string) bool {
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '$' && i+1 < len(pattern) {
			// Found $, check if followed by digits
			j := i + 1
			for j < len(pattern) && pattern[j] >= '0' && pattern[j] <= '9' {
				j++
			}
			// If we found at least one digit after $, it's a substitution variable
			if j > i+1 {
				return true
			}
		}
	}
	return false
}

// BuildFormStartEndMappings analyzes all tokens and dynamically builds the classification tables
func (ce *ClassifierEngine) BuildFormStartEndMappings(tokens []string, config *ClassifierConfig) error {
	// Step 1: Build a config-based StartTokenTable that maps start patterns to end substitutions
	// This will be used for both token analysis AND final classification
	configStartTableBuilder := regexptable.NewRegexpTableBuilder[[]string]()

	for _, surroundConfig := range config.SurroundRegexp {
		if surroundConfig.Start != "" {
			// Make a copy of the slice to avoid reference issues
			endPatternsCopy := make([]string, len(surroundConfig.End))
			copy(endPatternsCopy, surroundConfig.End)
			configStartTableBuilder.AddPattern(surroundConfig.Start, endPatternsCopy)
		}
	}

	var err error
	if len(config.SurroundRegexp) > 0 {
		ce.config.StartTokenTable, err = configStartTableBuilder.Build(true, true)
		if err != nil {
			return fmt.Errorf("failed to build start token table: %w", err)
		}
	}

	// Step 2: Initialize end-tokens set with constants (no substitution variables)
	endTokensSet := make(map[string]bool)

	// Step 3: Add constant end-tokens (those without $\d+ substitutions) from config
	for _, surroundConfig := range config.SurroundRegexp {
		for _, endPattern := range surroundConfig.End {
			// Check if this is a constant (no $\d+ substitutions)
			if !hasSubstitutionVariable(endPattern) {
				endTokensSet[endPattern] = true
			}
		}
	}

	// Step 4: Analyze tokens to generate dynamic end tokens
	for _, token := range tokens {
		if ce.config.StartTokenTable != nil {
			endSubsts, _, ok := ce.config.StartTokenTable.TryLookup(token)
			if ok {
				// Generate all substituted end tokens and add to set
				for _, endPattern := range endSubsts {
					endToken := simpleSubstitute(endPattern, token)
					endTokensSet[endToken] = true
				}
			}
		}
	}

	// Step 5: Build EndTokenTable from all collected end tokens
	endTableBuilder := regexptable.NewRegexpTableBuilder[bool]()
	for endToken := range endTokensSet {
		// Escape the end token since it should be matched literally, not as a regex
		escapedEndToken := regexp.QuoteMeta(endToken)
		endTableBuilder.AddPattern(escapedEndToken, true)
	}
	ce.config.EndTokenTable, err = endTableBuilder.Build(true, true)
	if err != nil {
		return fmt.Errorf("failed to build end token table: %w", err)
	}

	return nil
}

// ClassifyToken classifies a single token and returns the classification string
func (ce *ClassifierEngine) ClassifyToken(token string) string {
	// Check compound label first (highest priority)
	if ce.config.CompoundLabelRegexpTable != nil {
		_, _, ok := ce.config.CompoundLabelRegexpTable.TryLookup(token)
		if ok {
			return "C"
		}
	}

	// Check simple label
	if ce.config.SimpleLabelRegexpTable != nil {
		_, _, ok := ce.config.SimpleLabelRegexpTable.TryLookup(token)
		if ok {
			return "L"
		}
	}

	// Check form prefix
	if ce.config.FormPrefixRegexpTable != nil {
		_, _, ok := ce.config.FormPrefixRegexpTable.TryLookup(token)
		if ok {
			return "P"
		}
	}

	// Check form start using StartTokenTable BEFORE checking end tokens
	if ce.config.StartTokenTable != nil {
		endSubsts, _, ok := ce.config.StartTokenTable.TryLookup(token)
		if ok {
			// Generate the possible end tokens for display
			endTokens := make([]string, 0, len(endSubsts))
			for _, endPattern := range endSubsts {
				endToken := simpleSubstitute(endPattern, token)
				endTokens = append(endTokens, endToken)
			}
			if len(endTokens) > 0 {
				return "S " + strings.Join(endTokens, " ")
			}
			return "S"
		}
	}

	// Check if this token is an end token using EndTokenTable
	if ce.config.EndTokenTable != nil {
		_, _, ok := ce.config.EndTokenTable.TryLookup(token)
		if ok {
			return "E"
		}
	}

	// Check operator using OperatorRegexpTable
	// TODO: Fix type issue with OperatorRegexpTable
	// if ce.config.OperatorRegexpTable != nil {
	//     operatorTable := ce.config.OperatorRegexpTable
	//     _, operatorConfig, ok := operatorTable.TryLookup(token)
	//     if ok {
	//         return fmt.Sprintf("O %d %d %d", operatorConfig.PrefixPrec, operatorConfig.InfixPrec, operatorConfig.PostfixPrec)
	//     }
	// }

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
	err = engine.BuildFormStartEndMappings(tokens, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building form mappings: %v\n", err)
		os.Exit(1)
	}

	// Process tokens and output classifications
	engine.ProcessTokens(tokens)
}
