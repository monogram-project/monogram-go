// Package helpers provides utility types and functions that are designed to be
// self-contained and potentially extracted into separate libraries.
package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// RegexTable provides efficient multi-pattern regex classification using Go's built-in regexp.
// It compiles multiple regex patterns into a single automaton for optimal performance.
type RegexTable[T any] struct {
	compiled     *regexp.Regexp
	values       map[string]T
	patternNames []string
}

// NewRegexTable creates a new empty RegexTable.
func NewRegexTable[T any]() *RegexTable[T] {
	return &RegexTable[T]{
		values:       make(map[string]T),
		patternNames: make([]string, 0),
	}
}

// AddPattern adds a new regex pattern with its associated value to the table.
// The name must be unique and will be used to identify which pattern matched.
// Returns an error if the name is already in use or if recompilation fails.
func (rt *RegexTable[T]) AddPattern(name, pattern string, value T) error {
	groupName := fmt.Sprintf("__REGEXTABLE_%s__", name)

	if _, exists := rt.values[groupName]; exists {
		return fmt.Errorf("pattern name '%s' already exists", name)
	}

	// Create a unique capture group name with reserved prefix
	namedPattern := fmt.Sprintf("(?P<%s>%s)", groupName, pattern)

	rt.patternNames = append(rt.patternNames, namedPattern)
	rt.values[groupName] = value

	return rt.recompile()
}

// RemovePattern removes a pattern from the table by name.
// Returns an error if the pattern doesn't exist or if recompilation fails.
func (rt *RegexTable[T]) RemovePattern(name string) error {
	groupName := fmt.Sprintf("__REGEXTABLE_%s__", name)

	if _, exists := rt.values[groupName]; !exists {
		return fmt.Errorf("pattern name '%s' does not exist", name)
	}

	delete(rt.values, groupName)

	// Remove from pattern names slice
	for i, pName := range rt.patternNames {
		if strings.Contains(pName, groupName) {
			rt.patternNames = append(rt.patternNames[:i], rt.patternNames[i+1:]...)
			break
		}
	}

	return rt.recompile()
}

// HasPatterns returns true if the table has any patterns configured.
func (rt *RegexTable[T]) HasPatterns() bool {
	return len(rt.patternNames) > 0
}

// recompile rebuilds the union regex from all registered patterns.
func (rt *RegexTable[T]) recompile() error {
	if len(rt.patternNames) == 0 {
		rt.compiled = nil
		return nil
	}

	// Create union pattern: (?P<group1>pattern1)|(?P<group2>pattern2)|...
	unionPattern := "^(?:" + strings.Join(rt.patternNames, "|") + ")"

	var err error
	rt.compiled, err = regexp.Compile(unionPattern)
	if err != nil {
		return fmt.Errorf("failed to compile union regex: %w", err)
	}

	return nil
}

// Classify attempts to match the input string against all registered patterns.
// Returns the value, submatch slice, and error. If no patterns match, returns zero value, nil, error.
func (rt *RegexTable[T]) Classify(input string) (T, []string, error) {
	var zero T

	if rt.compiled == nil {
		return zero, nil, fmt.Errorf("no patterns configured")
	}

	matches := rt.compiled.FindStringSubmatch(input)
	if matches == nil {
		return zero, nil, fmt.Errorf("no pattern matched")
	}

	// Find which named group matched by checking submatches
	subexpNames := rt.compiled.SubexpNames()
	for i, name := range subexpNames {
		if name != "" && i < len(matches) && matches[i] != "" {
			if value, exists := rt.values[name]; exists {
				return value, matches, nil
			}
		}
	}

	return zero, nil, fmt.Errorf("internal error: match found but no capture group matched")
}

// TryClassify is like Classify but returns a boolean success indicator instead of an error.
// This is useful when you want to check if something matches without handling errors.
func (rt *RegexTable[T]) TryClassify(input string) (T, []string, bool) {
	value, matches, err := rt.Classify(input)
	return value, matches, err == nil
}
