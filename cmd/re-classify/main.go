package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sfkleach/regexptable"
)

// Pre-compiled regex for detecting non-zero substitution variables
var nonZeroSubstRegex = regexp.MustCompile(`\$[1-9]`)

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

// hasOnlyDollarZeroSubstitution checks if a pattern contains only $0 substitutions (no $1, $2, etc.)
func hasOnlyDollarZeroSubstitution(pattern string) bool {
	// Check if pattern contains $0
	hasDollarZero := strings.Contains(pattern, "$0")

	// Check if pattern contains any $[1-9] (which we want to avoid)
	hasNonZeroSubst := nonZeroSubstRegex.MatchString(pattern)

	// Return true only if we have $0 and no $[1-9]
	return hasDollarZero && !hasNonZeroSubst
}

// BuildFormStartEndMappings analyzes all tokens and dynamically builds the classification tables
func (ce *ClassifierEngine) BuildFormStartEndMappings(tokens []string, config *ClassifierConfig) error {
	// Step 1: Build a config-based StartTokenTable that maps start patterns to StartTokenInfo
	// This will be used for both token analysis AND final classification
	configStartTableBuilder := regexptable.NewRegexpTableBuilder[StartTokenInfo]()

	for i, surroundConfig := range config.SurroundRegexp {
		if surroundConfig.Start != "" {
			// Make a copy of the slice to avoid reference issues
			endPatternsCopy := make([]string, len(surroundConfig.Endings))
			copy(endPatternsCopy, surroundConfig.Endings)

			// Create StartTokenInfo with serial number and endings
			startInfo := StartTokenInfo{
				SerialNumber: i, // Use the index as the serial number
				Endings:      endPatternsCopy,
			}
			configStartTableBuilder.AddPattern(surroundConfig.Start, startInfo)
		}
	}

	var err error
	if len(config.SurroundRegexp) > 0 {
		ce.config.StartTokenTable, err = configStartTableBuilder.Build(true, true)
		if err != nil {
			return fmt.Errorf("failed to build start token table: %w", err)
		}
	}

	// Step 2: Initialize end-tokens map with serial numbers instead of just boolean
	endTokensMap := make(map[string]int) // Maps end token to serial number

	// Step 3: Build EndTokenTable builder and populate it from `end` patterns first
	endTableBuilder := regexptable.NewRegexpTableBuilder[int]()

	// Add patterns from the `end` field (if present) before scanning tokens
	for i, surroundConfig := range config.SurroundRegexp {
		serialNumber := i // Use index as serial number
		if surroundConfig.End != "" {
			// Use -1 if endings is not empty (we don't need reverse map for these)
			// Use actual serial number if endings is empty (we need reverse map for these)
			valueToStore := -1
			if len(surroundConfig.Endings) == 0 {
				valueToStore = serialNumber
			}
			// Add the end pattern to the builder
			endTableBuilder.AddPattern(surroundConfig.End, valueToStore)
		}
	}

	// Add constant end-tokens and patterns with single $0 substitution
	for i, surroundConfig := range config.SurroundRegexp {
		serialNumber := i // Use index as serial number
		for _, endPattern := range surroundConfig.Endings {
			// Check if this is a constant (no $\d+ substitutions)
			if !hasSubstitutionVariable(endPattern) {
				endTokensMap[endPattern] = serialNumber
			} else if hasOnlyDollarZeroSubstitution(endPattern) {
				// This pattern has only $0 substitutions (can be multiple occurrences)
				// We can create a regex pattern by replacing $0 with the start pattern
				if surroundConfig.Start != "" {
					// Replace $0 with the start regex pattern
					regexPattern := strings.ReplaceAll(endPattern, "$0", "("+surroundConfig.Start+")")
					// Add this directly as a regex pattern to the EndTokenTable
					endTableBuilder.AddPattern(regexPattern, serialNumber)
				}
			}
		}
	}

	// Step 3.5: Build initial EndTokenTable from `end` patterns to scan for present end tokens
	initialEndTokenTable, err := endTableBuilder.Build(true, true)
	if err != nil {
		return fmt.Errorf("failed to build initial end token table: %w", err)
	}

	// Step 3.6: Scan tokens to find which end tokens are present and build reverse map
	endTokensBySerial := make(map[int]map[string]bool) // Maps serial number to set of end tokens
	for _, token := range tokens {
		if serialNumber, _, ok := initialEndTokenTable.TryLookup(token); ok {
			// Only store in reverse map if serial number is not -1 (i.e., endings array is empty)
			if serialNumber != -1 {
				// This token matches an end pattern, add it to the set for this serial number
				if endTokensBySerial[serialNumber] == nil {
					endTokensBySerial[serialNumber] = make(map[string]bool)
				}
				endTokensBySerial[serialNumber][token] = true
			}
		}
	}

	// Step 4: Analyze tokens to generate dynamic end tokens
	for _, token := range tokens {
		if ce.config.StartTokenTable != nil {
			startInfo, captureGroups, ok := ce.config.StartTokenTable.TryLookup(token)
			if ok {
				if len(startInfo.Endings) > 0 {
					// Use the endings array to generate end tokens with full substitution
					for _, endPattern := range startInfo.Endings {
						endToken := substitutePattern(endPattern, captureGroups)
						endTokensMap[endToken] = startInfo.SerialNumber
					}
				} else {
					// Endings array is empty, use the reverse map to find tokens from the `end` regexp
					if endTokensSet, exists := endTokensBySerial[startInfo.SerialNumber]; exists {
						// Add all end tokens found for this serial number
						for endToken := range endTokensSet {
							endTokensMap[endToken] = startInfo.SerialNumber
						}
					}
				}
			}
		}
	}

	// Step 5: Add all other end patterns to the same builder with -1 (no reverse map needed)

	// Add constant end-tokens and patterns with single $0 substitution
	for i, surroundConfig := range config.SurroundRegexp {
		serialNumber := i // Use index as serial number
		for _, endPattern := range surroundConfig.Endings {
			// Check if this is a constant (no $\d+ substitutions)
			if !hasSubstitutionVariable(endPattern) {
				endTokensMap[endPattern] = serialNumber
			} else if hasOnlyDollarZeroSubstitution(endPattern) {
				// This pattern has only $0 substitutions (can be multiple occurrences)
				// We can create a regex pattern by replacing $0 with the start pattern
				if surroundConfig.Start != "" {
					// Replace $0 with the start regex pattern
					regexPattern := strings.ReplaceAll(endPattern, "$0", "("+surroundConfig.Start+")")
					// Add this directly as a regex pattern to the EndTokenTable with -1 (no reverse map needed)
					endTableBuilder.AddPattern(regexPattern, -1)
				}
			}
		}
	}

	// Add all literal end tokens to the builder (with escaping) using -1
	for endToken := range endTokensMap {
		// Escape the end token since it should be matched literally, not as a regex
		escapedEndToken := regexp.QuoteMeta(endToken)
		endTableBuilder.AddPattern(escapedEndToken, -1)
	}

	// Step 6: Build the final EndTokenTable using the same builder
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
		startInfo, captureGroups, ok := ce.config.StartTokenTable.TryLookup(token)
		if ok {
			// Generate the possible end tokens for display
			endTokens := make([]string, 0, len(startInfo.Endings))
			for _, endPattern := range startInfo.Endings {
				endToken := substitutePattern(endPattern, captureGroups)
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
		serialNumber, _, ok := ce.config.EndTokenTable.TryLookup(token)
		if ok {
			// For now, just return "E" - later we can use the serialNumber for more sophisticated matching
			_ = serialNumber // Acknowledge we have the serial number for future use
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
