// Package helpers provides utility types and functions that are designed to be
// self-contained and potentially extracted into separate libraries.
package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

// RegexClassifier provides efficient multi-pattern regex classification using Go's built-in regexp.
// It compiles multiple regex patterns into a single automaton for optimal performance.
type RegexClassifier[T any] struct {
	compiled     *regexp.Regexp
	values       map[string]T
	patternNames []string
}

// NewRegexClassifier creates a new empty RegexClassifier.
func NewRegexClassifier[T any]() *RegexClassifier[T] {
	return &RegexClassifier[T]{
		values:       make(map[string]T),
		patternNames: make([]string, 0),
	}
}

// AddPattern adds a new regex pattern with its associated value to the classifier.
// The name must be unique and will be used to identify which pattern matched.
// Returns an error if the name is already in use or if recompilation fails.
func (rc *RegexClassifier[T]) AddPattern(name, pattern string, value T) error {
	groupName := fmt.Sprintf("__REGEXCLASSIFIER_%s__", name)

	if _, exists := rc.values[groupName]; exists {
		return fmt.Errorf("pattern name '%s' already exists", name)
	}

	// Create a unique capture group name with reserved prefix
	namedPattern := fmt.Sprintf("(?P<%s>%s)", groupName, pattern)

	rc.patternNames = append(rc.patternNames, namedPattern)
	rc.values[groupName] = value

	return rc.recompile()
}

// RemovePattern removes a pattern from the classifier by name.
// Returns an error if the pattern doesn't exist or if recompilation fails.
func (rc *RegexClassifier[T]) RemovePattern(name string) error {
	groupName := fmt.Sprintf("__REGEXCLASSIFIER_%s__", name)

	if _, exists := rc.values[groupName]; !exists {
		return fmt.Errorf("pattern name '%s' does not exist", name)
	}

	delete(rc.values, groupName)

	// Remove from pattern names slice
	for i, pName := range rc.patternNames {
		if strings.Contains(pName, groupName) {
			rc.patternNames = append(rc.patternNames[:i], rc.patternNames[i+1:]...)
			break
		}
	}

	return rc.recompile()
}

// HasPatterns returns true if the classifier has any patterns configured.
func (rc *RegexClassifier[T]) HasPatterns() bool {
	return len(rc.patternNames) > 0
}

// recompile rebuilds the union regex from all registered patterns.
func (rc *RegexClassifier[T]) recompile() error {
	if len(rc.patternNames) == 0 {
		rc.compiled = nil
		return nil
	}

	// Create union pattern: (?P<group1>pattern1)|(?P<group2>pattern2)|...
	unionPattern := "^(?:" + strings.Join(rc.patternNames, "|") + ")"

	var err error
	rc.compiled, err = regexp.Compile(unionPattern)
	if err != nil {
		return fmt.Errorf("failed to compile union regex: %w", err)
	}

	return nil
}

// Classify attempts to match the input string against all registered patterns.
// Returns the value, submatch slice, and error. If no patterns match, returns zero value, nil, error.
func (rc *RegexClassifier[T]) Classify(input string) (T, []string, error) {
	var zero T

	if rc.compiled == nil {
		return zero, nil, fmt.Errorf("no patterns configured")
	}

	matches := rc.compiled.FindStringSubmatch(input)
	if matches == nil {
		return zero, nil, fmt.Errorf("no pattern matched")
	}

	// Find which named group matched by checking submatches
	subexpNames := rc.compiled.SubexpNames()
	for i, name := range subexpNames {
		if name != "" && i < len(matches) && matches[i] != "" {
			if value, exists := rc.values[name]; exists {
				return value, matches, nil
			}
		}
	}

	return zero, nil, fmt.Errorf("internal error: match found but no capture group matched")
}

// TryClassify is like Classify but returns a boolean success indicator instead of an error.
// This is useful when you want to check if something matches without handling errors.
func (rc *RegexClassifier[T]) TryClassify(input string) (T, []string, bool) {
	value, matches, err := rc.Classify(input)
	return value, matches, err == nil
}
