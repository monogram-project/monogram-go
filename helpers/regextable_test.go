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

	// Add some test patterns
	err := table.AddPattern("form_start", `form\w*`, TokenFormStart)
	if err != nil {
		t.Fatalf("Failed to add form_start pattern: %v", err)
	}

	err = table.AddPattern("form_end", `end\w*`, TokenFormEnd)
	if err != nil {
		t.Fatalf("Failed to add form_end pattern: %v", err)
	}

	err = table.AddPattern("simple_label", `[a-z]+:`, TokenSimpleLabel)
	if err != nil {
		t.Fatalf("Failed to add simple_label pattern: %v", err)
	}

	// Test successful matches
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
			value, match, err := table.Classify(tc.input)

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

func TestRegexTable_TryClassify(t *testing.T) {
	table := NewRegexTable[string]()

	err := table.AddPattern("test", `hello`, "greeting")
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Test successful match
	value, match, ok := table.TryClassify("hello")
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
	value, match, ok = table.TryClassify("goodbye")
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

	err := table.AddPattern("test1", `foo`, 1)
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	err = table.AddPattern("test2", `bar`, 2)
	if err != nil {
		t.Fatalf("Failed to add pattern: %v", err)
	}

	// Verify both patterns work
	value, _, err := table.Classify("foo")
	if err != nil || value != 1 {
		t.Error("Pattern test1 should match")
	}

	value, _, err = table.Classify("bar")
	if err != nil || value != 2 {
		t.Error("Pattern test2 should match")
	}

	// Remove first pattern
	err = table.RemovePattern("test1")
	if err != nil {
		t.Fatalf("Failed to remove pattern: %v", err)
	}

	// Verify first pattern no longer works
	_, _, err = table.Classify("foo")
	if err == nil {
		t.Error("Pattern test1 should no longer match after removal")
	}

	// Verify second pattern still works
	value, _, err = table.Classify("bar")
	if err != nil || value != 2 {
		t.Error("Pattern test2 should still match after removing test1")
	}
}
