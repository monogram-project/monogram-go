package main

import (
	"fmt"
	"regexp"

	"github.com/monogram-project/monogram-go/helpers"
)

// DotNetRegexEngine simulates .NET regex syntax
type DotNetRegexEngine struct{}

func (e *DotNetRegexEngine) Compile(pattern string) (helpers.CompiledRegex, error) {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return helpers.NewStandardCompiledRegex(compiled), nil
}

func (e *DotNetRegexEngine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?<%s>%s)", groupName, pattern)
}

func main() {
	fmt.Println("=== RegexTable Engine Abstraction Demo ===\n")

	// Create tables with different engines
	goTable := helpers.NewRegexTableBuilder[string]().
		AddPattern("if.*", "form_start").
		AddPattern("end.*", "form_end").
		AddPattern("else", "simple_label").
		MustBuild()

	dotNetTable := helpers.NewRegexTableBuilderWithEngine[string](&DotNetRegexEngine{}).
		AddPattern("if.*", "form_start").
		AddPattern("end.*", "form_end").
		AddPattern("else", "simple_label").
		MustBuild()

	testInput := "else"

	// Both should produce the same result despite different internal regex syntax
	goResult, _, goOk := goTable.TryLookup(testInput)
	dotNetResult, _, dotNetOk := dotNetTable.TryLookup(testInput)

	fmt.Printf("Testing input: %q\n", testInput)
	fmt.Printf("Go engine:     %s (found: %t)\n", goResult, goOk)
	fmt.Printf(".NET engine:   %s (found: %t)\n", dotNetResult, dotNetOk)

	// Show the internal regex patterns to demonstrate differences
	fmt.Printf("\nInternal named group formats:\n")
	fmt.Printf("Go style:      %s\n", helpers.NewStandardRegexEngine().FormatNamedGroup("test", "pattern"))
	fmt.Printf(".NET style:    %s\n", (&DotNetRegexEngine{}).FormatNamedGroup("test", "pattern"))

	// Test more inputs
	fmt.Printf("\nTesting various inputs:\n")
	testCases := []string{"if", "ifThen", "endfor", "endif", "else", "nomatch"}
	for _, test := range testCases {
		goResult, _, goOk := goTable.TryLookup(test)
		dotNetResult, _, dotNetOk := dotNetTable.TryLookup(test)

		fmt.Printf("%-8s -> Go: %-12s .NET: %-12s", test,
			formatResult(goResult, goOk),
			formatResult(dotNetResult, dotNetOk))

		if goOk == dotNetOk && goResult == dotNetResult {
			fmt.Printf(" ✓\n")
		} else {
			fmt.Printf(" ✗ MISMATCH\n")
		}
	}
}

func formatResult(result string, found bool) string {
	if found {
		return result
	}
	return "no match"
}
