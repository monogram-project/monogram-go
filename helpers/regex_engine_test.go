package helpers

import (
	"fmt"
	"testing"
)

// MockRegexEngine implements RegexEngine for testing with different naming conventions.
type MockRegexEngine struct {
	compiledRegexes map[string]*MockCompiledRegex
	groupSyntax     string // e.g., "(?P<%s>%s)" for Go, "(?<%s>%s)" for .NET
}

// NewMockRegexEngine creates a new mock engine with the specified group syntax.
func NewMockRegexEngine(groupSyntax string) *MockRegexEngine {
	return &MockRegexEngine{
		compiledRegexes: make(map[string]*MockCompiledRegex),
		groupSyntax:     groupSyntax,
	}
}

// Compile returns a pre-configured mock or creates a simple one.
func (e *MockRegexEngine) Compile(pattern string) (CompiledRegex, error) {
	if compiled, exists := e.compiledRegexes[pattern]; exists {
		return compiled, nil
	}
	// Return a simple mock that doesn't actually match anything
	return &MockCompiledRegex{pattern: pattern}, nil
}

// FormatNamedGroup uses the configured group syntax.
func (e *MockRegexEngine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf(e.groupSyntax, groupName, pattern)
}

// SetCompiledRegex allows tests to configure what a pattern should return.
func (e *MockRegexEngine) SetCompiledRegex(pattern string, compiled *MockCompiledRegex) {
	e.compiledRegexes[pattern] = compiled
}

// MockCompiledRegex implements CompiledRegex for testing.
type MockCompiledRegex struct {
	pattern     string
	matchResult []string
	subexpNames []string
	shouldMatch bool
}

// SetMatchResult configures what FindStringSubmatch should return.
func (r *MockCompiledRegex) SetMatchResult(matches []string, subexpNames []string) {
	r.matchResult = matches
	r.subexpNames = subexpNames
	r.shouldMatch = matches != nil
}

// FindStringSubmatch returns the configured match result.
func (r *MockCompiledRegex) FindStringSubmatch(s string) []string {
	if r.shouldMatch {
		return r.matchResult
	}
	return nil
}

// SubexpNames returns the configured subexpression names.
func (r *MockCompiledRegex) SubexpNames() []string {
	return r.subexpNames
}

func TestRegexEngine_DifferentNamingConventions(t *testing.T) {
	testCases := []struct {
		name        string
		groupSyntax string
		expected    string
	}{
		{
			name:        "Go/PCRE syntax",
			groupSyntax: "(?P<%s>%s)",
			expected:    "(?P<test_group>hello)",
		},
		{
			name:        ".NET syntax",
			groupSyntax: "(?<%s>%s)",
			expected:    "(?<test_group>hello)",
		},
		{
			name:        "Custom syntax",
			groupSyntax: "(?'%s'%s)",
			expected:    "(?'test_group'hello)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			engine := NewMockRegexEngine(tc.groupSyntax)
			result := engine.FormatNamedGroup("test_group", "hello")
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestRegexTable_WithCustomEngine(t *testing.T) {
	// Create a mock engine with .NET-style syntax
	mockEngine := NewMockRegexEngine("(?<%s>%s)")

	// Create a mock compiled regex that will match "hello"
	mockCompiled := &MockCompiledRegex{}
	mockCompiled.SetMatchResult(
		[]string{"hello", "hello"},       // full match and capture group
		[]string{"", "__REGEXTABLE_1__"}, // subexp names
	)

	// Configure the engine to return our mock when compiling the union pattern
	// This is a bit artificial but demonstrates the concept
	table := NewRegexTableWithEngine[string](mockEngine)

	// Add a pattern - this should use the .NET syntax
	_, err := table.AddPattern("hello", "greeting")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Verify the pattern was formatted with .NET syntax
	expectedPattern := "(?<__REGEXTABLE_1__>hello)"
	if len(table.patternNames) != 1 || table.patternNames[0] != expectedPattern {
		t.Errorf("Expected pattern %q, got %q", expectedPattern, table.patternNames[0])
	}
}

func TestStandardRegexEngine_FormatNamedGroup(t *testing.T) {
	engine := NewStandardRegexEngine()
	result := engine.FormatNamedGroup("testgroup", "pattern")
	expected := "(?P<testgroup>pattern)"

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
