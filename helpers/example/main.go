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
	// Create a new table for TokenType
	table := helpers.NewRegexTable[TokenType]()

	// Add patterns for different token types
	if _, err := table.AddPattern(`form\w*`, TokenFormStart); err != nil {
		log.Fatal(err)
	}

	if _, err := table.AddPattern(`end\w*`, TokenFormEnd); err != nil {
		log.Fatal(err)
	}

	if _, err := table.AddPattern(`[a-z]+:`, TokenSimpleLabel); err != nil {
		log.Fatal(err)
	}

	// Test various inputs
	testInputs := []string{
		"form",
		"formData",
		"endform",
		"endif",
		"else:",
		"unknown",
	}

	fmt.Println("=== Classification Results ===")
	for _, input := range testInputs {
		if tokenType, matches, err := table.Classify(input); err == nil {
			fmt.Printf("'%s' -> %s (full match: '%s')\n",
				input, tokenType, matches[0])
		} else {
			fmt.Printf("'%s' -> No match\n", input)
		}
	}

	fmt.Println("\n=== Using TryClassify ===")
	for _, input := range testInputs {
		if tokenType, matches, ok := table.TryClassify(input); ok {
			fmt.Printf("'%s' -> %s (full match: '%s')\n",
				input, tokenType, matches[0])
		} else {
			fmt.Printf("'%s' -> No match\n", input)
		}
	}
}
