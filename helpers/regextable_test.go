package helpers

import (
	"testing"
)

// TokenType represents different types of tokens for testing
type TokenType int

const (
	TokenFormStart TokenType = iota
	TokenFormEnd
	TokenSimpleLabel
	TokenVariable
)

func TestRegexTable_Basic(t *testing.T) {
	table := NewRegexTable[TokenType]()

	// Add some test patterns using deferred compilation
	_, err := table.AddPattern(`form\w*`, TokenFormStart)
	if err != nil {
		t.Fatalf("Failed to add form_start pattern: %v", err)
	}

	_, err = table.AddPattern(`end\w*`, TokenFormEnd)
	if err != nil {
		t.Fatalf("Failed to add form_end pattern: %v", err)
	}

	_, err = table.AddPattern(`[a-z]+:`, TokenSimpleLabel)
	if err != nil {
		t.Fatalf("Failed to add simple_label pattern: %v", err)
	}

	// Test successful matches (compilation happens here)
	testCases := []struct {
		input       string
		expected    TokenType
		shouldMatch bool
	}{
		{"form", TokenFormStart, true},
		{"formData", TokenFormStart, true},
		{"endform", TokenFormEnd, true},
		{"endif", TokenFormEnd, true},
		{"else:", TokenSimpleLabel, true},
		{"nomatch", TokenVariable, false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			value, match, err := table.Lookup(tc.input)

			if tc.shouldMatch {
				if err != nil {
					t.Errorf("Expected match for '%s', but got error: %v", tc.input, err)
					return
				}
				if value != tc.expected {
					t.Errorf("Expected %v for '%s', got %v", tc.expected, tc.input, value)
				}
				if match == nil {
					t.Errorf("Expected non-nil match for '%s'", tc.input)
				}
			} else {
				if err == nil {
					t.Errorf("Expected no match for '%s', but got value: %v", tc.input, value)
				}
				if match != nil {
					t.Errorf("Expected nil match for '%s', but got: %v", tc.input, match)
				}
			}
		})
	}
}

func TestRegexTable_TryLookup(t *testing.T) {
	table := NewRegexTable[string]()

	_, err := table.AddPattern(`hello`, "greeting")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Test successful match
	value, match, ok := table.TryLookup("hello")
	if !ok {
		t.Error("Expected successful match for 'hello'")
	}
	if value != "greeting" {
		t.Errorf("Expected 'greeting', got '%s'", value)
	}
	if match == nil {
		t.Error("Expected non-nil match")
	}

	// Test unsuccessful match
	value, match, ok = table.TryLookup("goodbye")
	if ok {
		t.Error("Expected unsuccessful match for 'goodbye'")
	}
	if value != "" {
		t.Errorf("Expected empty string for no match, got '%s'", value)
	}
	if match != nil {
		t.Error("Expected nil match for no match")
	}
}

func TestRegexTable_RemovePattern(t *testing.T) {
	table := NewRegexTable[int]()

	id1, err := table.AddPattern(`foo`, 1)
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	id2, err := table.AddPattern(`bar`, 2)
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}
	_ = id2 // We keep this pattern for testing

	// Verify both patterns work
	value, _, err := table.Lookup("foo")
	if err != nil || value != 1 {
		t.Error("Pattern test1 should match")
	}

	value, _, err = table.Lookup("bar")
	if err != nil || value != 2 {
		t.Error("Pattern test2 should match")
	}

	// Remove first pattern
	err = table.RemovePattern(id1)
	if err != nil {
		t.Fatalf("Failed to remove pattern: %v", err)
	}

	// Verify first pattern no longer works
	_, _, err = table.Lookup("foo")
	if err == nil {
		t.Error("Pattern test1 should no longer match after removal")
	}

	// Verify second pattern still works
	value, _, err = table.Lookup("bar")
	if err != nil || value != 2 {
		t.Error("Pattern test2 should still match after removing test1")
	}
}

func TestRegexTable_LazyVsImmediateCompilation(t *testing.T) {
	// Test lazy compilation
	lazy := NewRegexTable[string]()

	// These should succeed without compilation
	_, err := lazy.AddPattern("valid", "value1")
	if err != nil {
		t.Errorf("Lazy AddPattern should succeed: %v", err)
	}

	// This should fail at classification time, not add time
	_, err = lazy.AddPattern("[invalid", "value2") // Invalid regex
	if err != nil {
		t.Errorf("Lazy AddPattern should defer validation: %v", err)
	}

	// Classification should fail due to invalid regex
	_, _, err = lazy.Lookup("test")
	if err == nil {
		t.Error("Expected lookup to fail due to invalid regex")
	}

	// Test immediate compilation
	immediate := NewRegexTable[string]()

	// Valid pattern should succeed
	_, err = immediate.AddPatternThenRecompile("valid", "value1")
	if err != nil {
		t.Errorf("Immediate AddPattern should succeed: %v", err)
	}

	// Invalid pattern should fail immediately
	_, err = immediate.AddPatternThenRecompile("[invalid", "value2") // Invalid regex
	if err == nil {
		t.Error("Expected immediate AddPattern to fail with invalid regex")
	}
}

func TestRegexTable_ManualRecompile(t *testing.T) {
	table := NewRegexTable[string]()

	// Add patterns without compilation
	_, err := table.AddPattern("hello", "greeting")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	_, err = table.AddPattern("world", "place")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Manually trigger compilation
	err = table.Recompile()
	if err != nil {
		t.Fatalf("Manual recompile failed: %v", err)
	}

	// Now lookup should work
	value, _, err := table.Lookup("hello")
	if err != nil {
		t.Fatalf("Lookup failed: %v", err)
	}
	if value != "greeting" {
		t.Errorf("Expected 'greeting', got '%s'", value)
	}
}

func TestRegexTable_RemovePatternThenRecompile(t *testing.T) {
	table := NewRegexTable[string]()

	// Add patterns using immediate compilation
	id1, err := table.AddPatternThenRecompile("hello", "greeting")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	_, err = table.AddPatternThenRecompile("world", "place")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Remove pattern with immediate recompilation
	err = table.RemovePatternThenRecompile(id1)
	if err != nil {
		t.Fatalf("Failed to remove pattern: %v", err)
	}

	// Verify the pattern is gone
	_, _, err = table.Lookup("hello")
	if err == nil {
		t.Error("Expected hello pattern to be removed")
	}

	// Verify other pattern still works
	value, _, err := table.Lookup("world")
	if err != nil {
		t.Fatalf("World pattern should still work: %v", err)
	}
	if value != "place" {
		t.Errorf("Expected 'place', got '%s'", value)
	}
}
