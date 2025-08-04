package main

// This is the entry point of the program. It processes command-line arguments
// and performs file format translation based on user-specified flags. The
// program supports built-in formats (e.g., XML and JSON) as well as delegating
// to external subprograms for custom formats.
//
// Flags:
// - --help: Displays help information for the program and available flags.
// - --format (-f): Specifies the output format. Required for both built-in and external formats.
// - --input (-i): Specifies the input file. If omitted, standard input (stdin) is used.
// - --output (-o): Specifies the output file. If omitted, standard output (stdout) is used.
//
// Built-in Formats:
// - xml: The program processes input and outputs in XML format.
// - json: The program processes input and outputs in JSON format.
// Additional built-in formats can be added by updating the global formatHandlers map.
//
// For non-built-in formats, the program delegates processing to a subprogram named "monogram-to-{format}".
//
// Usage Example:
// To translate a file to JSON format:
//
//	monogram --format json --input input.txt --output output.json
//
// To delegate to a custom subprogram:
//
//	monogram --format custom --input input.txt --output output.custom

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/monogram-project/monogram-go/mg"
	"github.com/spf13/pflag"
)

// Controlled by ldflags.
var IsBuiltForDocker = "false"

// setupFlags initializes a flag set with the common flag definitions.
func setupFlags(fs *pflag.FlagSet, options *mg.FormatOptions, configFile *string, showHelp *bool, classifyTokens *bool, showVersion *bool, testPort *string, openBrowserFlag *bool) {
	fs.StringVarP(&options.Format, "format", "f", options.Format, "Output format xml|json|yaml|mermaid|dot")
	fs.StringVarP(&options.Input, "input", "i", options.Input, "Input file (optional, defaults to stdin)")
	fs.StringVarP(&options.Output, "output", "o", options.Output, "Output file (optional, defaults to stdout)")
	fs.IntVar(&options.Indent, "indent", options.Indent, "Number of spaces for indentation (0 for no formatting)")
	fs.BoolVar(&options.Limit, "one", options.Limit, "Process only one monogram value and do not wrap in a unit node")
	fs.StringVarP(&options.DefaultLabel, "default-label", "b", options.DefaultLabel, "Default labels")
	fs.BoolVar(&options.IncludeSpans, "include-spans", options.IncludeSpans, "Include start/end of expressions in the output")
	fs.BoolVar(&options.Decimal, "decimal", options.Decimal, "Decode numbers integers and floats in base 10")
	fs.BoolVar(&options.CheckLiterals, "check-literals", options.CheckLiterals, "Check regexs and other literal strings for validity")
	if configFile != nil {
		fs.StringVarP(configFile, "config", "c", "", "Configuration file (YAML format)")
	}
	if showHelp != nil {
		fs.BoolVarP(showHelp, "help", "h", false, "Display help information")
	}
	if classifyTokens != nil {
		fs.BoolVar(classifyTokens, "classify-tokens", false, "Classify tokens")
	}
	if showVersion != nil {
		fs.BoolVar(showVersion, "version", false, "Display the version information")
	}
	if WithWeb {
		if openBrowserFlag != nil {
			fs.BoolVar(openBrowserFlag, "open-browser", true, "Specify whether --test automatically opens a browser")
		}
		if testPort != nil {
			pflag.StringVar(testPort, "test", "", "Start HTTP test server on specified port (optional, e.g., 3000)")
			pflag.Lookup("test").NoOptDefVal = "8080"
		}
	}
}

// Define a type for the translation function
// type translationFunc func(io.Reader, io.Writer, *mg.FormatOptions)
type translateFunc func(root *mg.Node, indentDelta string, output io.Writer)

type formatHandler struct {
	Format string
	Name   string
	Fn     translateFunc
}

// Global map for format-to-function associations
var formatHandlerList = []formatHandler{
	{Format: "xml", Name: "XML", Fn: mg.PrintASTXML},
	{Format: "json", Name: "JSON", Fn: mg.PrintASTJSON},
	{Format: "yaml", Name: "YAML", Fn: mg.PrintASTYAML},
	{Format: "mermaid", Name: "Mermaid", Fn: mg.PrintASTMermaid},
	{Format: "dot", Name: "Dot", Fn: mg.PrintASTDOT},
	{Format: "asciitree", Name: "asciitree", Fn: mg.PrintASTAsciiTree},
}

var formatToFormatHandler = func() map[string]formatHandler {
	handler := make(map[string]formatHandler)
	for _, f := range formatHandlerList {
		handler[f.Format] = f
	}
	return handler
}()

var nameToFormatHandler = func() map[string]formatHandler {
	handler := make(map[string]formatHandler)
	for _, f := range formatHandlerList {
		handler[f.Name] = f
	}
	return handler
}()

// Compute availableFormatNames from the keys in formatHandlers. Maintain
// the order of the formatHandlerList for consistent output.
var availableFormatNames = func() []string {
	var formats []string
	for _, f := range formatHandlerList {
		formats = append(formats, f.Name)
	}
	return formats
}()

func parseToAST(input string, foptions *mg.FormatOptions, classifiers *mg.TokenClassifiers) (*mg.Node, error) {
	p_opts := &mg.ParserOptions{
		CoreFormatOptions: foptions.CoreFormatOptions,
		TokenClassifiers:  classifiers,
	}
	return p_opts.ParseToAST(input, foptions.Input, foptions.Limit)
}

func main() {
	// Initialize the options struct
	options := mg.FormatOptions{
		Input:  "",
		Output: "",
		Limit:  false,
		CoreFormatOptions: mg.CoreFormatOptions{
			Format:       "",
			Indent:       2,
			DefaultLabel: "_",
			IncludeSpans: false,
			Decimal:      false,
		},
	}

	var configFile string
	var showHelp bool
	var classifyTokens bool
	var showVersion bool // New variable for the --version flag
	var testPort string
	openBrowserFlag := true

	// Set up the main command-line flag set
	setupFlags(pflag.CommandLine, &options, &configFile, &showHelp, &classifyTokens, &showVersion, &testPort, &openBrowserFlag)

	// Parse command-line flags
	pflag.Parse()

	// Load configuration file if specified
	var config *mg.Config
	var err error
	if configFile != "" {
		config, err = mg.LoadConfig(configFile)
		if err != nil {
			log.Fatalf("Error loading config file: %v", err)
		}

		// Track which flags were explicitly set to avoid overriding them with config defaults
		flagsExplicitlySet := make(map[string]bool)
		pflag.Visit(func(flag *pflag.Flag) {
			flagsExplicitlySet[flag.Name] = true
		})

		// Apply config defaults only for flags that weren't explicitly set
		config.ApplyConfigDefaults(&options, flagsExplicitlySet)
	}

	if testPort != "" {
		startTestServer(testPort, openBrowserFlag, &options, config)
		os.Exit(0) // Exit after printing the version, cannot be reached at present.
	}

	// Check for the version flag
	if showVersion {
		fmt.Printf("Monogram version: v%s\n", mg.Version)
		os.Exit(0) // Exit after printing the version
	}

	// Check for help flag
	if showHelp {
		fmt.Println("Monogram: converts program-like text in monogram notation to various other formats.")
		fmt.Println("\nUsage:")
		fmt.Println("  monogram [OPTIONS] < STDIN > STDOUT")
		fmt.Println("\nOptions:")
		pflag.PrintDefaults()
		os.Exit(0) // Exit after displaying the help message
	}

	// Check if the format is built-in
	translator, isBuiltInFormat := formatToFormatHandler[options.Format]

	// Open input (default to stdin if input is not provided)
	var inputReader io.Reader
	if options.Input == "" {
		inputReader = os.Stdin
	} else {
		file, err := os.Open(options.Input)
		if err != nil {
			log.Fatalf("Error: Cannot open input file '%s': %v", options.Input, err)
		}
		defer file.Close()
		inputReader = file
	}

	// Open output (default to stdout if output is not provided)
	var outputWriter io.Writer
	if options.Output == "" {
		outputWriter = os.Stdout
	} else {
		file, err := os.Create(options.Output)
		if err != nil {
			log.Fatalf("Error: Cannot create output file '%s': %v", options.Output, err)
		}
		defer file.Close()
		outputWriter = file
	}

	// Handle built-in formats
	if isBuiltInFormat {
		err := translator.translate(inputReader, outputWriter, &options, config)
		if err != nil {
			log.Fatalf("Error: Failed to translate input: %v", err)
		}
		return
	} else if classifyTokens {
		mg.VSCodeClassifyTokens(inputReader, outputWriter)
		return
	}

	// For non-built-in formats, exec into a subprogram
	if options.Format == "" {
		log.Fatalf("Error: Format was not specified.")
	}

	execName := "monogram-to-" + options.Format
	newArgs := make([]string, len(os.Args))
	newArgs[0] = execName
	copy(newArgs[1:], os.Args[1:])

	err = syscall.Exec(execName, newArgs, os.Environ())
	if err != nil {
		log.Fatalf("Failed to execute %s: %v", execName, err)
	}
}

func (printAST *formatHandler) translate(input io.Reader, output io.Writer, options *mg.FormatOptions, config *mg.Config) error {
	return translate(input, output, printAST.Fn, options, config)
}

func translate(input io.Reader, output io.Writer, printAST func(*mg.Node, string, io.Writer), options *mg.FormatOptions, config *mg.Config) error {
	// Read the entire input as a string
	data, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("failed to read input: %v", err)
	}

	// Get TokenClassifiers from config, or use empty if no config
	var classifiers *mg.TokenClassifiers
	if config != nil {
		classifiers = &config.TokenClassifiers
	} else {
		classifiers = &mg.TokenClassifiers{}
	}

	// Convert the input string into an AST
	ast, err := parseToAST(string(data), options, classifiers)
	if err != nil {
		return err
	}

	// Determine the indentation string (spaces or none)
	indent := ""
	if options.Indent > 0 {
		indent = strings.Repeat(" ", options.Indent)
	}

	// Use the provided print function to recursively print the AST
	printAST(ast, indent, output)

	return nil
}
