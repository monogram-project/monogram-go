package main

import (
	"fmt"
	"log"

	"github.com/monogram-project/monogram-go/helpers"
)

// TokenType represents different types of tokens
type TokenType int

const (
	TokenKeyword TokenType = iota
	TokenIdentifier
	TokenNumber
	TokenOperator
	TokenString
)

func (t TokenType) String() string {
	switch t {
	case TokenKeyword:
		return "Keyword"
	case TokenIdentifier:
		return "Identifier"
	case TokenNumber:
		return "Number"
	case TokenOperator:
		return "Operator"
	case TokenString:
		return "String"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("=== RegexTableBuilder Demo ===")

	// Using the builder pattern - no need to think about compilation!
	table, err := helpers.NewRegexTableBuilder[TokenType]().
		AddPattern(`\b(if|else|while|for|return|function)\b`, TokenKeyword).
		AddPattern(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`, TokenIdentifier).
		AddPattern(`\b\d+(\.\d+)?\b`, TokenNumber).
		AddPattern(`[+\-*/=<>!]+`, TokenOperator).
		AddPattern(`"[^"]*"`, TokenString).
		Build()

	if err != nil {
		log.Fatalf("Failed to build regex table: %v", err)
	}

	fmt.Println("✓ Table built successfully with single compilation!")

	// Test various code snippets
	testInputs := []string{
		"if",         // keyword
		"variable",   // identifier
		"42",         // number
		"3.14",       // decimal number
		"++",         // operator
		`"hello"`,    // string
		"unknown123", // should not match (has numbers in identifier)
	}

	fmt.Println("\n=== Classification Results ===")
	for _, input := range testInputs {
		if tokenType, matches, ok := table.TryClassify(input); ok {
			fmt.Printf("'%s' -> %s (matched: '%s')\n", input, tokenType, matches[0])
		} else {
			fmt.Printf("'%s' -> No match\n", input)
		}
	}

	fmt.Println("\n=== Builder Chaining Demo ===")

	// Show how builder can be reused and chained
	baseBuilder := helpers.NewRegexTableBuilder[string]().
		AddPattern(`form\w*`, "form_start").
		AddPattern(`end\w*`, "form_end")

	// Clone and extend for different use cases
	webBuilder := baseBuilder.Clone().
		AddPattern(`button\w*`, "button").
		AddPattern(`input\w*`, "input")

	codeBuilder := baseBuilder.Clone().
		AddPattern(`class\w*`, "class_def").
		AddPattern(`method\w*`, "method_def")

	// Build specialized tables
	webTable, err := webBuilder.Build()
	if err != nil {
		log.Fatalf("Failed to build web table: %v", err)
	}

	codeTable, err := codeBuilder.Build()
	if err != nil {
		log.Fatalf("Failed to build code table: %v", err)
	}

	// Test specialized tables
	fmt.Println("\nWeb table:")
	webInputs := []string{"form", "button", "input", "class"}
	for _, input := range webInputs {
		if value, _, ok := webTable.TryClassify(input); ok {
			fmt.Printf("  '%s' -> %s\n", input, value)
		} else {
			fmt.Printf("  '%s' -> No match\n", input)
		}
	}

	fmt.Println("\nCode table:")
	codeInputs := []string{"form", "class", "method", "button"}
	for _, input := range codeInputs {
		if value, _, ok := codeTable.TryClassify(input); ok {
			fmt.Printf("  '%s' -> %s\n", input, value)
		} else {
			fmt.Printf("  '%s' -> No match\n", input)
		}
	}

	fmt.Println("\n=== MustBuild for Static Configs ===")

	// For static configurations where patterns are known to be valid
	staticTable := helpers.NewRegexTableBuilder[string]().
		AddPattern(`config_\w+`, "config").
		AddPattern(`static_\w+`, "static").
		MustBuild() // Will panic if patterns are invalid

	fmt.Println("✓ Static table built with MustBuild")

	if value, _, ok := staticTable.TryClassify("config_file"); ok {
		fmt.Printf("'config_file' -> %s\n", value)
	}
}
