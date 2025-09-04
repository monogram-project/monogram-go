package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sfkleach/regexptable"
)

// Version is set at build time via -ldflags
var Version = "unknown"

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

// BuildFormStartEndMappings analyzes all tokens and dynamically builds the classification tables
func (ce *ClassifierEngine) BuildFormStartEndMappings(tokens []string, config *ClassifierConfig) error {

	// Build a config-based StartTokenTable that maps start patterns to StartTokenInfo.
	configStartTableBuilder := regexptable.NewRegexpTableBuilder[*StartTokenInfo]()
	startTokenInfoList := make([]*StartTokenInfo, len(config.SurroundRegexp))
	for i, surroundConfig := range config.SurroundRegexp {
		if surroundConfig.Start != "" {
			// Create StartTokenInfo with serial number and endings
			startInfo := &StartTokenInfo{
				SerialNumber: i, // Use the index as the serial number
				Endings:      make(map[string]bool),
			}
			for _, ending := range surroundConfig.Endings {
				startInfo.Endings[ending] = true
			}

			startTokenInfoList[i] = startInfo
			configStartTableBuilder.AddPattern(surroundConfig.Start, startInfo)
		}
	}
	t, err := configStartTableBuilder.Build(true, true)
	if err != nil {
		return fmt.Errorf("failed to build start token table: %w", err)
	}
	ce.config.StartTokenTable = t

	// When Endings is not set, the startInfoTokens will be missing proper
	// endings. So we must infer the endings from the end patterns
	// applied to the list of tokens and backfill the startInfoTokens.
	count := 0
	inferEndingsTableBuilder := regexptable.NewRegexpTableBuilder[int]()
	for i, surroundConfig := range config.SurroundRegexp {
		if len(surroundConfig.Endings) == 0 && surroundConfig.End != "" {
			inferEndingsTableBuilder.AddPattern(surroundConfig.End, i)

		}
	}
	if count > 0 {
		it, err := inferEndingsTableBuilder.Build(true, true)
		if err != nil {
			return fmt.Errorf("failed to build inferred endings table: %w", err)
		}
		// If there are no explicit endings, we need to find all tokens that match the end pattern
		for _, token := range tokens {
			if serialNumber, _, ok := it.TryLookup(token); ok {
				startTokenInfoList[serialNumber].Endings[token] = true
			}
		}
	}

	// Now we create the ce.config.EndTokenTable - but a backfill obligation
	// may remain.
	backfillEnd := make(map[int]bool, 0)
	endTokenTableBuilder := regexptable.NewRegexpTableBuilder[bool]()
	for i, surroundConfig := range config.SurroundRegexp {
		if surroundConfig.End != "" {
			endTokenTableBuilder.AddPattern(surroundConfig.End, true)
		} else {
			// If there is no End then we must infer it from the Endings
			// pattern, if possible.
			for _, ending := range surroundConfig.Endings {
				// Does the pattern contain $0 or $N, N>1.
				hasDollarZero := strings.Contains(ending, "$0")
				hasDollarNonZero := nonZeroSubstRegex.MatchString(ending)
				if !hasDollarZero && !hasDollarNonZero {
					endTokenTableBuilder.AddPattern(regexp.QuoteMeta(ending), true)
				} else if hasDollarZero && !hasDollarNonZero {
					// Split at $0 and QuoteMeta the components then join
					// using the Start regexp.
					startPattern := regexp.QuoteMeta(surroundConfig.Start)
					parts := strings.Split(ending, "$0")
					for i, part := range parts {
						parts[i] = regexp.QuoteMeta(part)
					}
					endTokenTableBuilder.AddPattern(strings.Join(parts, startPattern), true)
				} else {
					// We will need to backfill this pattern by applying the
					// endings to actual tokens.
					backfillEnd[i] = true
				}
			}
		}
	}

	if len(backfillEnd) > 0 {
		// We need to backfill the end patterns for these tokens.
		for _, token := range tokens {
			if info, _, ok := ce.config.StartTokenTable.TryLookup(token); ok {
				if backfillEnd[info.SerialNumber] {
					// Backfill the end pattern for this token
					endTokenTableBuilder.AddPattern(regexp.QuoteMeta(token), true)
				}
			}
		}
	}

	// Now we can construct ce.config.EndTokenTable.
	ce.config.EndTokenTable, err = endTokenTableBuilder.Build(true, true)
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
			for endPattern := range startInfo.Endings {
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
	// Define command-line flags
	checkOnly := flag.Bool("check", false, "Validate configuration syntax only (don't process input)")
	version := flag.Bool("version", false, "Show version information")

	// Customize usage message
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <config.yaml>\n\n", os.Args[0])
		fmt.Println("re-classify is a token classification tool that uses regex patterns")
		fmt.Println("to classify identifiers and operators in monogram syntax.")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nThe tool reads tokens from stdin (one per line) and outputs")
		fmt.Println("classification results based on the regex patterns in config.yaml")
	}

	// Parse command-line flags
	flag.Parse()

	// Handle version flag
	if *version {
		fmt.Printf("re-classify version %s\n", Version)
		return
	}

	// Check for required config file argument
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Error: exactly one config file must be specified\n\n")
		flag.Usage()
		os.Exit(1)
	}

	configFile := args[0]

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

	// If check-only mode, just report success and exit
	if *checkOnly {
		fmt.Println("Configuration syntax is valid")
		return
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
