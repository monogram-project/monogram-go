package helpers

import (
	"strings"
	"testing"
)

func TestRegexTableBuilder_Basic(t *testing.T) {
	builder := NewRegexTableBuilder[string]()

	// Chain pattern additions
	table, err := builder.
		AddPattern("hello", "greeting").
		AddPattern("world", "place").
		AddPattern(`\d+`, "number").
		Build()

	if err != nil {
		t.Fatalf("Failed to build table: %v", err)
	}

	// Test the built table
	testCases := []struct {
		input       string
		expected    string
		shouldMatch bool
	}{
		{"hello", "greeting", true},
		{"world", "place", true},
		{"123", "number", true},
		{"unknown", "", false},
	}

	for _, tc := range testCases {
		value, _, ok := table.TryLookup(tc.input)
		if tc.shouldMatch {
			if !ok {
				t.Errorf("Expected match for '%s'", tc.input)
			} else if value != tc.expected {
				t.Errorf("Expected '%s' for '%s', got '%s'", tc.expected, tc.input, value)
			}
		} else {
			if ok {
				t.Errorf("Expected no match for '%s', got '%s'", tc.input, value)
			}
		}
	}
}

func TestRegexTableBuilder_InvalidPattern(t *testing.T) {
	builder := NewRegexTableBuilder[string]()

	// Add valid and invalid patterns
	table, err := builder.
		AddPattern("valid", "good").
		AddPattern("[invalid", "bad"). // Invalid regex
		Build()

	if err == nil {
		t.Error("Expected build to fail with invalid pattern")
	}

	if table != nil {
		t.Error("Expected nil table when build fails")
	}

	// Error should mention the invalid pattern or compilation failure
	if !strings.Contains(err.Error(), "invalid pattern") && !strings.Contains(err.Error(), "failed to compile") {
		t.Errorf("Expected error to mention invalid pattern or compilation failure, got: %v", err)
	}
}

func TestRegexTableBuilder_MustBuild(t *testing.T) {
	// Test successful MustBuild
	builder := NewRegexTableBuilder[int]()
	table := builder.
		AddPattern("test", 42).
		MustBuild()

	value, _, ok := table.TryLookup("test")
	if !ok || value != 42 {
		t.Error("MustBuild should create working table")
	}

	// Test panic on invalid pattern
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected MustBuild to panic with invalid pattern")
		}
	}()

	invalidBuilder := NewRegexTableBuilder[int]()
	invalidBuilder.AddPattern("[invalid", 1).MustBuild() // Should panic
}

func TestRegexTableBuilder_Utilities(t *testing.T) {
	builder := NewRegexTableBuilder[string]()

	// Test empty builder
	if builder.HasPatterns() {
		t.Error("New builder should have no patterns")
	}
	if builder.PatternCount() != 0 {
		t.Error("New builder should have 0 pattern count")
	}

	// Add patterns
	builder.AddPattern("test1", "value1")
	builder.AddPattern("test2", "value2")

	if !builder.HasPatterns() {
		t.Error("Builder should have patterns after adding")
	}
	if builder.PatternCount() != 2 {
		t.Errorf("Expected 2 patterns, got %d", builder.PatternCount())
	}

	// Test clear
	builder.Clear()
	if builder.HasPatterns() {
		t.Error("Builder should have no patterns after clear")
	}
	if builder.PatternCount() != 0 {
		t.Error("Builder should have 0 patterns after clear")
	}
}

func TestRegexTableBuilder_Clone(t *testing.T) {
	original := NewRegexTableBuilder[string]()
	original.AddPattern("test1", "value1")
	original.AddPattern("test2", "value2")

	clone := original.Clone()

	// Verify clone has same patterns
	if clone.PatternCount() != original.PatternCount() {
		t.Error("Clone should have same pattern count as original")
	}

	// Add to clone and verify original is unchanged
	clone.AddPattern("test3", "value3")

	if original.PatternCount() != 2 {
		t.Error("Original should be unchanged after modifying clone")
	}
	if clone.PatternCount() != 3 {
		t.Error("Clone should have additional pattern")
	}

	// Build both and verify they work correctly
	originalTable, err := original.Build()
	if err != nil {
		t.Fatalf("Failed to build original: %v", err)
	}

	cloneTable, err := clone.Build()
	if err != nil {
		t.Fatalf("Failed to build clone: %v", err)
	}

	// Original should not have test3
	_, _, ok := originalTable.TryLookup("test3")
	if ok {
		t.Error("Original table should not have test3 pattern")
	}

	// Clone should have test3
	value, _, ok := cloneTable.TryLookup("test3")
	if !ok || value != "value3" {
		t.Error("Clone table should have test3 pattern")
	}
}

func TestRegexTableBuilder_EmptyBuild(t *testing.T) {
	builder := NewRegexTableBuilder[string]()

	// Build empty table
	table, err := builder.Build()
	if err != nil {
		t.Fatalf("Building empty table should not fail: %v", err)
	}

	// Empty table should return error on lookup
	_, _, err = table.Lookup("anything")
	if err == nil {
		t.Error("Empty table should return error on lookup")
	}
}

func TestRegexTableBuilder_ReuseAfterBuild(t *testing.T) {
	builder := NewRegexTableBuilder[string]()

	// Build first table
	table1, err := builder.
		AddPattern("test1", "value1").
		Build()
	if err != nil {
		t.Fatalf("Failed to build first table: %v", err)
	}

	// Add more patterns and build second table
	table2, err := builder.
		AddPattern("test2", "value2").
		Build()
	if err != nil {
		t.Fatalf("Failed to build second table: %v", err)
	}

	// Both tables should work
	value1, _, ok := table1.TryLookup("test1")
	if !ok || value1 != "value1" {
		t.Error("First table should work")
	}

	// Second table should have both patterns
	value1, _, ok = table2.TryLookup("test1")
	if !ok || value1 != "value1" {
		t.Error("Second table should have first pattern")
	}

	value2, _, ok := table2.TryLookup("test2")
	if !ok || value2 != "value2" {
		t.Error("Second table should have second pattern")
	}
}
