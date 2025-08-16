package helpers

import (
	"fmt"
)

// RegexTableBuilder provides a convenient builder pattern for creating RegexTable instances.
// It accumulates patterns and builds the final RegexTable with a single compilation step.
type RegexTableBuilder[T any] struct {
	patterns []patternEntry[T]
	engine   RegexEngine
}

// patternEntry holds a pattern and its associated value during building
type patternEntry[T any] struct {
	pattern string
	value   T
}

// NewRegexTableBuilder creates a new builder for RegexTable[T] using the standard regex engine.
func NewRegexTableBuilder[T any]() *RegexTableBuilder[T] {
	return NewRegexTableBuilderWithEngine[T](NewStandardRegexEngine())
}

// NewRegexTableBuilderWithEngine creates a new builder for RegexTable[T] with a custom regex engine.
func NewRegexTableBuilderWithEngine[T any](engine RegexEngine) *RegexTableBuilder[T] {
	return &RegexTableBuilder[T]{
		patterns: make([]patternEntry[T], 0),
		engine:   engine,
	}
}

// AddPattern adds a pattern to be included in the final RegexTable.
// This method never fails - validation happens during Build().
func (b *RegexTableBuilder[T]) AddPattern(pattern string, value T) *RegexTableBuilder[T] {
	b.patterns = append(b.patterns, patternEntry[T]{
		pattern: pattern,
		value:   value,
	})
	return b
}

// Build creates the final RegexTable with all accumulated patterns.
// This is when compilation and validation occur.
func (b *RegexTableBuilder[T]) Build() (*RegexTable[T], error) {
	table := NewRegexTableWithEngine[T](b.engine)

	// Add all patterns to the table (using lazy compilation)
	for _, entry := range b.patterns {
		_, err := table.AddPattern(entry.pattern, entry.value)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern '%s': %w", entry.pattern, err)
		}
	}

	// Trigger compilation once at the end
	err := table.Recompile()
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex table: %w", err)
	}

	return table, nil
} // MustBuild is like Build but panics on error. Useful for static configurations
// where patterns are known to be valid.
func (b *RegexTableBuilder[T]) MustBuild() *RegexTable[T] {
	table, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("RegexTableBuilder.MustBuild failed: %v", err))
	}
	return table
}

// HasPatterns returns true if any patterns have been added to the builder.
func (b *RegexTableBuilder[T]) HasPatterns() bool {
	return len(b.patterns) > 0
}

// PatternCount returns the number of patterns that have been added.
func (b *RegexTableBuilder[T]) PatternCount() int {
	return len(b.patterns)
}

// Clear removes all patterns from the builder, allowing it to be reused.
func (b *RegexTableBuilder[T]) Clear() *RegexTableBuilder[T] {
	b.patterns = b.patterns[:0] // Reset slice but keep capacity
	return b
}

// Clone creates a copy of the builder with the same patterns and engine.
func (b *RegexTableBuilder[T]) Clone() *RegexTableBuilder[T] {
	clone := NewRegexTableBuilderWithEngine[T](b.engine)
	clone.patterns = make([]patternEntry[T], len(b.patterns))
	copy(clone.patterns, b.patterns)
	return clone
}
