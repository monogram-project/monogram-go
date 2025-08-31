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
	tokens []string
}

// NewClassifierEngine creates a new classifier engine with the given configuration
func NewClassifierEngine(config *CompiledClassifierConfig) *ClassifierEngine {
	return &ClassifierEngine{
		config: config,
	}
}

// BuildFormStartEndMappings analyzes all tokens and dynamically builds the classification tables
func (ce *ClassifierEngine) BuildFormStartEndMappings(tokens []string, config *ClassifierConfig) error {
	// First debug line - use stdout to ensure it's visible
	fmt.Println("DEBUG_VISIBLE: BuildFormStartEndMappings called")
	fmt.Fprintf(os.Stderr, "DEBUG: Starting token analysis - this should be visible")

	// Step 1: Create empty StartTokenTable and initial set of end-tokens
	startTableBuilder := regexptable.NewRegexpTableBuilder[[]string]()
	endTokensSet := make(map[string]bool)
	formSurroundPatterns := make(map[string]bool)

	fmt.Fprintf(os.Stderr, "DEBUG: Starting token analysis with %d tokens...\n", len(tokens))
	fmt.Fprintf(os.Stderr, "DEBUG: Config has %d surround-regexp entries\n", len(config.SurroundRegexp))

	// Step 2: Populate initial constant end-tokens (those without $\d+ substitutions)
	for i, surroundConfig := range config.SurroundRegexp {
		fmt.Fprintf(os.Stderr, "DEBUG: Surround config %d: start='%s', end=%v\n", i, surroundConfig.Start, surroundConfig.End)
		for _, endPattern := range surroundConfig.End {
			// Check if this is a constant (no $\d+ substitutions)
			if !regexp.MustCompile(`\$\d+`).MatchString(endPattern) {
				endTokensSet[endPattern] = true
				fmt.Fprintf(os.Stderr, "DEBUG: Added constant end token: '%s'\n", endPattern)
			}
		}
	}

	// Step 3: First pass - populate StartTokenTable and collect all possible end tokens
	fmt.Fprintf(os.Stderr, "DEBUG: Starting first pass over tokens\n")
	for _, token := range tokens {
		fmt.Fprintf(os.Stderr, "DEBUG: Processing token: '%s'\n", token)
		for _, surroundConfig := range config.SurroundRegexp {
			if surroundConfig.Start != "" {
				// Check if this token matches the start pattern (already anchored by regexp.MatchString)
				pattern := "^" + surroundConfig.Start + "$"
				matched, err := regexp.MatchString(pattern, token)
				if err != nil {
					return fmt.Errorf("invalid start pattern '%s': %w", surroundConfig.Start, err)
				}

				fmt.Fprintf(os.Stderr, "DEBUG: Testing token '%s' against pattern '%s' -> %v\n", token, pattern, matched)

				if matched {
					fmt.Fprintf(os.Stderr, "DEBUG: Token '%s' matches start pattern '%s'\n", token, surroundConfig.Start)
					// Add exact token to StartTokenTable (no manual anchoring, Build() will handle it)
					startTableBuilder.AddPattern(token, surroundConfig.End)
					fmt.Fprintf(os.Stderr, "DEBUG: Added token '%s' to StartTokenTable with end substs: %v\n", token, surroundConfig.End)

					// Generate all substituted end tokens and add to set
					for _, endPattern := range surroundConfig.End {
						endToken := simpleSubstitute(endPattern, token)
						endTokensSet[endToken] = true
						fmt.Fprintf(os.Stderr, "DEBUG: Generated end token '%s' from pattern '%s' + token '%s'\n", endToken, endPattern, token)

						// Add to FormSurroundMatch patterns
						formSurroundPattern := token + " " + endToken
						formSurroundPatterns[formSurroundPattern] = true
						fmt.Fprintf(os.Stderr, "DEBUG: Added form surround pattern: '%s'\n", formSurroundPattern)
					}
				}
			}
		}
	}

	// Step 4: Build the final tables with anchoring
	fmt.Fprintf(os.Stderr, "DEBUG: Building StartTokenTable...\n")
	var err error
	ce.config.StartTokenTable, err = startTableBuilder.Build(true, true)
	if err != nil {
		return fmt.Errorf("failed to build start token table: %w", err)
	}

	// Build EndTokenTable (no manual anchoring, Build() will handle it)
	fmt.Fprintf(os.Stderr, "DEBUG: Building EndTokenTable with %d end tokens\n", len(endTokensSet))
	endTableBuilder := regexptable.NewRegexpTableBuilder[bool]()
	for endToken := range endTokensSet {
		endTableBuilder.AddPattern(endToken, true)
		fmt.Fprintf(os.Stderr, "DEBUG: Added end token to table: '%s'\n", endToken)
	}
	ce.config.EndTokenTable, err = endTableBuilder.Build(true, true)
	if err != nil {
		return fmt.Errorf("failed to build end token table: %w", err)
	}

	// Build FormSurroundMatch tables (no manual anchoring, Build() will handle it)
	if len(formSurroundPatterns) > 0 {
		fmt.Fprintf(os.Stderr, "DEBUG: Building FormSurroundMatch table with %d patterns\n", len(formSurroundPatterns))
		surroundTableBuilder := regexptable.NewRegexpTableBuilder[bool]()
		for pattern := range formSurroundPatterns {
			surroundTableBuilder.AddPattern(pattern, true)
			fmt.Fprintf(os.Stderr, "DEBUG: Added surround pattern: '%s'\n", pattern)
		}
		surroundTable, err := surroundTableBuilder.Build(true, true)
		if err != nil {
			return fmt.Errorf("failed to build form surround table: %w", err)
		}
		ce.config.FormSurroundMatch = []*regexptable.RegexpTable[bool]{surroundTable}
	}

	fmt.Fprintf(os.Stderr, "DEBUG: Token analysis complete\n")
	return nil
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

	// Check if this token is an end token using EndTokenTable
	if ce.config.EndTokenTable != nil {
		_, _, ok := ce.config.EndTokenTable.TryLookup(token)
		if ok {
			return "E"
		}
	}

	// Check form start using StartTokenTable
	if ce.config.StartTokenTable != nil {
		_, endSubsts, ok := ce.config.StartTokenTable.TryLookup(token)
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
	fmt.Fprintf(os.Stderr, "DEBUG: Program started\n")

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

	fmt.Printf("DEBUG: Read %d tokens: %v\n", len(tokens), tokens)

	// Build form-start to form-end mappings by analyzing all tokens
	fmt.Fprintf(os.Stderr, "DEBUG: About to call BuildFormStartEndMappings\n")
	err = engine.BuildFormStartEndMappings(tokens, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building form mappings: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "DEBUG: BuildFormStartEndMappings completed successfully\n")

	// Process tokens and output classifications
	fmt.Fprintf(os.Stderr, "DEBUG: About to process tokens\n")
	engine.ProcessTokens(tokens)

	// Debug output after processing
	fmt.Fprintf(os.Stderr, "DEBUG: Processed %d tokens\n", len(tokens))
}
