package helpers

import (
	"fmt"
	"regexp"
)

// DotNetRegexEngine simulates how .NET regex engines work with different named group syntax.
type DotNetRegexEngine struct{}

// NewDotNetRegexEngine creates a new .NET-style regex engine.
func NewDotNetRegexEngine() *DotNetRegexEngine {
	return &DotNetRegexEngine{}
}

// Compile wraps Go's regex with .NET-compatible interface.
func (e *DotNetRegexEngine) Compile(pattern string) (CompiledRegex, error) {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &StandardCompiledRegex{compiled}, nil
}

// FormatNamedGroup uses .NET-style named capture group syntax.
func (e *DotNetRegexEngine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?<%s>%s)", groupName, pattern)
}

// JavaRegexEngine simulates how Java regex engines work.
type JavaRegexEngine struct{}

// NewJavaRegexEngine creates a new Java-style regex engine.
func NewJavaRegexEngine() *JavaRegexEngine {
	return &JavaRegexEngine{}
}

// Compile wraps Go's regex with Java-compatible interface.
func (e *JavaRegexEngine) Compile(pattern string) (CompiledRegex, error) {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &StandardCompiledRegex{compiled}, nil
}

// FormatNamedGroup uses Java-style named capture group syntax.
func (e *JavaRegexEngine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?<%s>%s)", groupName, pattern)
}

// ExampleDifferentEngines demonstrates how to use different regex engines.
func ExampleDifferentEngines() {
	// Create tables with different engines
	goTable := NewRegexTableBuilder[string]().
		AddPattern("if.*", "form_start").
		AddPattern("end.*", "form_end").
		AddPattern("else", "simple_label").
		MustBuild()

	dotNetTable := NewRegexTableBuilderWithEngine[string](NewDotNetRegexEngine()).
		AddPattern("if.*", "form_start").
		AddPattern("end.*", "form_end").
		AddPattern("else", "simple_label").
		MustBuild()

	javaTable := NewRegexTableBuilderWithEngine[string](NewJavaRegexEngine()).
		AddPattern("if.*", "form_start").
		AddPattern("end.*", "form_end").
		AddPattern("else", "simple_label").
		MustBuild()

	testInput := "else"

	// All should produce the same result despite different internal regex syntax
	goResult, _, goOk := goTable.TryLookup(testInput)
	dotNetResult, _, dotNetOk := dotNetTable.TryLookup(testInput)
	javaResult, _, javaOk := javaTable.TryLookup(testInput)

	fmt.Printf("Go engine:     %s (found: %t)\n", goResult, goOk)
	fmt.Printf(".NET engine:   %s (found: %t)\n", dotNetResult, dotNetOk)
	fmt.Printf("Java engine:   %s (found: %t)\n", javaResult, javaOk)

	// Show the internal regex patterns to demonstrate differences
	fmt.Printf("\nInternal patterns:\n")
	fmt.Printf("Go style:      %s\n", goTable.engine.FormatNamedGroup("test", "pattern"))
	fmt.Printf(".NET style:    %s\n", dotNetTable.engine.FormatNamedGroup("test", "pattern"))
	fmt.Printf("Java style:    %s\n", javaTable.engine.FormatNamedGroup("test", "pattern"))
}
