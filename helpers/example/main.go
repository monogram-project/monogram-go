package main

import (
	"fmt"
	"log"

	"github.com/monogram-project/monogram-go/helpers"
)

// TokenType represents different types of tokens
type TokenType int

const (
	TokenFormStart TokenType = iota
	TokenFormEnd
	TokenSimpleLabel
	TokenVariable
)

func (t TokenType) String() string {
	switch t {
	case TokenFormStart:
		return "FormStart"
	case TokenFormEnd:
		return "FormEnd"
	case TokenSimpleLabel:
		return "SimpleLabel"
	case TokenVariable:
		return "Variable"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("=== Lazy Compilation Demo ===")

	// Create a new table for TokenType
	table := helpers.NewRegexTable[TokenType]()

	// Add patterns - these are deferred until classification
	fmt.Println("Adding patterns (no compilation yet)...")
	if _, err := table.AddPattern(`form\w*`, TokenFormStart); err != nil {
		log.Fatal(err)
	}

	if _, err := table.AddPattern(`end\w*`, TokenFormEnd); err != nil {
		log.Fatal(err)
	}

	if _, err := table.AddPattern(`[a-z]+:`, TokenSimpleLabel); err != nil {
		log.Fatal(err)
	}

	// Test various inputs - compilation happens on first classify
	testInputs := []string{
		"form",
		"formData",
		"endform",
		"endif",
		"else:",
		"unknown",
	}

	fmt.Println("\n=== Classification Tests (triggers compilation) ===")
	for _, input := range testInputs {
		if tokenType, matches, err := table.Classify(input); err == nil {
			fmt.Printf("'%s' -> %s (full match: '%s')\n",
				input, tokenType, matches[0])
		} else {
			fmt.Printf("'%s' -> No match\n", input)
		}
	}

	fmt.Println("\n=== Immediate Compilation Demo ===")

	// Create another table for immediate compilation
	immediateTable := helpers.NewRegexTable[TokenType]()

	fmt.Println("Adding pattern with immediate compilation...")
	if _, err := immediateTable.AddPatternThenRecompile(`test\w*`, TokenVariable); err != nil {
		log.Fatal(err)
	}

	// This will succeed because compilation already happened
	if tokenType, matches, ok := immediateTable.TryClassify("testing"); ok {
		fmt.Printf("'testing' -> %s (full match: '%s')\n", tokenType, matches[0])
	}

	fmt.Println("\n=== Manual Recompile Demo ===")

	// Create another table and add patterns
	manualTable := helpers.NewRegexTable[TokenType]()
	if _, err := manualTable.AddPattern(`manual\w*`, TokenVariable); err != nil {
		log.Fatal(err)
	}

	// Manually trigger compilation
	fmt.Println("Manually triggering compilation...")
	if err := manualTable.Recompile(); err != nil {
		log.Fatal(err)
	}

	// Now classification works
	if tokenType, matches, ok := manualTable.TryClassify("manually"); ok {
		fmt.Printf("'manually' -> %s (full match: '%s')\n", tokenType, matches[0])
	}
}
