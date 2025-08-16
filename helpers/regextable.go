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
	compiled       *regexp.Regexp
	values         map[string]T
	patternNames   []string
	nextID         int
	needsRecompile bool
}

// NewRegexTable creates a new empty RegexTable.
func NewRegexTable[T any]() *RegexTable[T] {
	return &RegexTable[T]{
		values:         make(map[string]T),
		patternNames:   make([]string, 0),
		nextID:         1,
		needsRecompile: false,
	}
}

// AddPattern adds a new regex pattern with its associated value to the table.
// Returns the pattern ID (for use with RemovePattern) and an error if regex compilation fails.
// This method defers recompilation until Lookup is called for better performance.
func (rt *RegexTable[T]) AddPattern(pattern string, value T) (int, error) {
	// Auto-generate a unique internal name
	patternID := rt.nextID
	groupName := fmt.Sprintf("__REGEXTABLE_%d__", patternID)
	rt.nextID++

	// Create a unique capture group name with reserved prefix
	namedPattern := fmt.Sprintf("(?P<%s>%s)", groupName, pattern)

	rt.patternNames = append(rt.patternNames, namedPattern)
	rt.values[groupName] = value
	rt.needsRecompile = true

	return patternID, nil
}

// AddPatternThenRecompile is like AddPattern but immediately recompiles the regex.
// Use this when you need immediate validation of the pattern or when you're only adding one pattern.
func (rt *RegexTable[T]) AddPatternThenRecompile(pattern string, value T) (int, error) {
	patternID, err := rt.AddPattern(pattern, value)
	if err != nil {
		return 0, err
	}

	err = rt.Recompile()
	if err != nil {
		// Rollback on error
		rt.RemovePattern(patternID)
		return 0, err
	}

	return patternID, nil
}

// RemovePattern removes a pattern from the table by its ID.
// This method defers recompilation until Lookup is called for better performance.
func (rt *RegexTable[T]) RemovePattern(patternID int) error {
	groupName := fmt.Sprintf("__REGEXTABLE_%d__", patternID)

	if _, exists := rt.values[groupName]; !exists {
		return fmt.Errorf("pattern ID %d does not exist", patternID)
	}

	delete(rt.values, groupName)

	// Remove from pattern names slice
	for i, pName := range rt.patternNames {
		if strings.Contains(pName, groupName) {
			rt.patternNames = append(rt.patternNames[:i], rt.patternNames[i+1:]...)
			break
		}
	}

	rt.needsRecompile = true
	return nil
}

// RemovePatternThenRecompile is like RemovePattern but immediately recompiles the regex.
// Use this when you need immediate validation or when you're only removing one pattern.
func (rt *RegexTable[T]) RemovePatternThenRecompile(patternID int) error {
	err := rt.RemovePattern(patternID)
	if err != nil {
		return err
	}

	return rt.Recompile()
}

// HasPatterns returns true if the table has any patterns configured.
func (rt *RegexTable[T]) HasPatterns() bool {
	return len(rt.patternNames) > 0
}

// Recompile rebuilds the union regex from all registered patterns.
// This is exposed to allow manual control over when recompilation occurs.
func (rt *RegexTable[T]) Recompile() error {
	if len(rt.patternNames) == 0 {
		rt.compiled = nil
		rt.needsRecompile = false
		return nil
	}

	// Create union pattern: (?P<group1>pattern1)|(?P<group2>pattern2)|...
	unionPattern := "^(?:" + strings.Join(rt.patternNames, "|") + ")"

	var err error
	rt.compiled, err = regexp.Compile(unionPattern)
	if err != nil {
		return fmt.Errorf("failed to compile union regex: %w", err)
	}

	rt.needsRecompile = false
	return nil
}

// ensureCompiled ensures the regex is compiled before use, recompiling if necessary.
func (rt *RegexTable[T]) ensureCompiled() error {
	if rt.needsRecompile || rt.compiled == nil {
		return rt.Recompile()
	}
	return nil
}

// Lookup attempts to match the input string against all registered patterns.
// Returns the value, submatch slice, and error. If no patterns match, returns zero value, nil, error.
// This method automatically recompiles the regex if patterns have been added/removed since last compilation.
func (rt *RegexTable[T]) Lookup(input string) (T, []string, error) {
	var zero T

	err := rt.ensureCompiled()
	if err != nil {
		return zero, nil, err
	}

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

// TryLookup is like Lookup but returns a boolean success indicator instead of an error.
// This is useful when you want to check if something matches without handling errors.
// This method automatically recompiles the regex if patterns have been added/removed since last compilation.
func (rt *RegexTable[T]) TryLookup(input string) (T, []string, bool) {
	value, matches, err := rt.Lookup(input)
	return value, matches, err == nil
}
