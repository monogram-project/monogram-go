# RegexTable

A high-performance multi-pattern regex classifier for Go using the built-in `regexp` package. This helper compiles multiple regex patterns into a single automaton for efficient pattern matching.

## Features

- **High Performance**: Uses a single compiled regex with named capture groups for O(n) matching regardless of pattern count
- **Builder Pattern**: `RegexTableBuilder` provides a convenient API that hides compilation complexity
- **Lazy Compilation**: Defers regex compilation until classification for better performance when adding multiple patterns
- **Type Safe**: Generic implementation supports any value type `T`
- **Reserved Namespace**: Uses `__REGEXTABLE_` prefix to avoid conflicts with user-defined capture groups
- **Built-in Regexp**: Uses Go's standard `regexp` package - no external dependencies
- **Full Match Access**: Returns both the classified value and complete submatch details
- **Self-Contained**: Designed to be extracted into a separate library

## Quick Start (Recommended)

```go
package main

import (
    "fmt"
    "github.com/monogram-project/monogram-go/helpers"
)

type TokenType int

const (
    TokenKeyword TokenType = iota
    TokenIdentifier
    TokenNumber
)

func main() {
    // Use the builder pattern - no need to think about compilation!
    table, err := helpers.NewRegexTableBuilder[TokenType]().
        AddPattern(`\b(if|else|while|for)\b`, TokenKeyword).
        AddPattern(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`, TokenIdentifier).
        AddPattern(`\b\d+\b`, TokenNumber).
        Build()
    
    if err != nil {
        panic(err) // Only fails if patterns are invalid
    }
    
    // Use the table
    if value, matches, ok := table.TryClassify("if"); ok {
        fmt.Printf("Matched: %v (%s)\n", value, matches[0])
    }
}
```

## Builder Usage Patterns

### Basic Builder Usage

```go
// Simple fluent interface
table, err := helpers.NewRegexTableBuilder[string]().
    AddPattern("hello", "greeting").
    AddPattern("world", "place").
    Build()
```

### Builder Chaining and Reuse

```go
// Create a base builder
base := helpers.NewRegexTableBuilder[TokenType]().
    AddPattern(`form\w*`, FormStart).
    AddPattern(`end\w*`, FormEnd)

// Clone and extend for different contexts  
webBuilder := base.Clone().
    AddPattern(`button\w*`, Button).
    AddPattern(`input\w*`, Input)

codeBuilder := base.Clone().
    AddPattern(`class\w*`, ClassDef).
    AddPattern(`method\w*`, MethodDef)

// Build specialized tables
webTable, _ := webBuilder.Build()
codeTable, _ := codeBuilder.Build()
```

### Static Configuration with MustBuild

```go
// For static configs where patterns are known valid
var GlobalTokenTable = helpers.NewRegexTableBuilder[TokenType]().
    AddPattern(`\b(if|else|while|for)\b`, Keyword).
    AddPattern(`\b[a-zA-Z_]\w*\b`, Identifier).
    AddPattern(`\b\d+\b`, Number).
    MustBuild() // Panics if patterns are invalid

func classifyToken(input string) TokenType {
    value, _, _ := GlobalTokenTable.TryClassify(input)
    return value
}
```

### Builder State Management

```go
builder := helpers.NewRegexTableBuilder[string]()

// Add patterns
builder.AddPattern("test1", "value1")
builder.AddPattern("test2", "value2")

fmt.Printf("Patterns: %d\n", builder.PatternCount()) // 2
fmt.Printf("Has patterns: %t\n", builder.HasPatterns()) // true

// Clear and reuse
builder.Clear()
fmt.Printf("Patterns after clear: %d\n", builder.PatternCount()) // 0
```

## API Reference

### Builder Pattern (Recommended)

#### `NewRegexTableBuilder[T any]() *RegexTableBuilder[T]`
Creates a new builder for RegexTable[T].

#### `AddPattern(pattern string, value T) *RegexTableBuilder[T]`
Adds a pattern to the builder. Returns the builder for method chaining.

#### `Build() (*RegexTable[T], error)`
Creates the final RegexTable with all accumulated patterns. Compilation happens here.

#### `MustBuild() *RegexTable[T]`
Like Build but panics on error. Useful for static configurations.

#### `Clone() *RegexTableBuilder[T]`
Creates a copy of the builder with the same patterns.

#### `Clear() *RegexTableBuilder[T]`
Removes all patterns from the builder.

#### `HasPatterns() bool`
Returns true if any patterns have been added.

#### `PatternCount() int`
Returns the number of patterns added.

### Direct RegexTable API

#### `NewRegexTable[T any]() *RegexTable[T]`
Creates a new empty RegexTable for values of type T.

#### `AddPattern(pattern string, value T) (int, error)`
Adds a regex pattern with its associated value to the table. Returns the pattern ID for later removal.
**Note**: This method uses lazy compilation - the regex is not compiled until classification is performed.

#### `AddPatternThenRecompile(pattern string, value T) (int, error)`
Like AddPattern but immediately recompiles the regex. Use this when you need immediate validation 
of the pattern or when you're only adding one pattern.

#### `RemovePattern(patternID int) error`
Removes a pattern from the table by its ID.
**Note**: This method uses lazy compilation - the regex is not recompiled until classification is performed.

#### `RemovePatternThenRecompile(patternID int) error`
Like RemovePattern but immediately recompiles the regex. Use this when you need immediate validation
or when you're only removing one pattern.

#### `Recompile() error`
Manually rebuilds the union regex from all registered patterns. This is exposed to allow manual 
control over when recompilation occurs.

#### `Classify(input string) (T, []string, error)`
Attempts to match the input against all registered patterns. Returns the associated value, 
submatch slice, and error. Automatically recompiles if patterns have been added/removed.

#### `TryClassify(input string) (T, []string, bool)`
Like Classify but returns a boolean success indicator instead of an error.

#### `HasPatterns() bool`
Returns true if the table has any patterns configured.

## Pattern Management

### Adding Patterns

```go
table := helpers.NewRegexTable[string]()

// Returns pattern ID for later removal
id, err := table.AddPattern(`\d+`, "number")
if err != nil {
    // Handle regex compilation error
}
```

### Removing Patterns

```go
// Remove by pattern ID
err := table.RemovePattern(id)
if err != nil {
    // Handle removal error (pattern not found)
}
```

### Pattern IDs

Pattern IDs are auto-generated integers starting from 1. They're used internally with the format `__REGEXTABLE_N__` where N is the ID. This reserved namespace prevents conflicts with user-defined capture groups.

## Error Handling

```go
// Method 1: Using Classify with error handling
if value, matches, err := table.Classify(input); err != nil {
    switch {
    case strings.Contains(err.Error(), "no patterns configured"):
        // Handle empty table
    case strings.Contains(err.Error(), "no pattern matched"):
        // Handle no match
    default:
        // Handle other errors
    }
}

// Method 2: Using TryClassify for simple success/failure
if value, matches, ok := table.TryClassify(input); ok {
    // Handle successful match
} else {
    // Handle no match
}
```

## Performance Considerations

### Lazy vs Immediate Compilation

- **Lazy compilation** (default): Best when adding multiple patterns at once
- **Immediate compilation**: Best when you need immediate error feedback or adding single patterns
- **Manual compilation**: Best when you want precise control over compilation timing

### Union Regex Performance

RegexTable compiles all patterns into a single union regex like:
```
^(?:(?P<__REGEXTABLE_1__>pattern1)|(?P<__REGEXTABLE_2__>pattern2)|(?P<__REGEXTABLE_3__>pattern3))
```

This provides O(n) matching performance regardless of the number of patterns, as opposed to O(n*m) when testing patterns individually.

## Advanced Usage

### Complex Pattern Matching

```go
type TokenInfo struct {
    Type     string
    Category string
}

table := helpers.NewRegexTable[TokenInfo]()

// Add complex patterns with rich metadata
table.AddPattern(`\b(if|else|while|for)\b`, TokenInfo{
    Type:     "keyword",
    Category: "control",
})

table.AddPattern(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`, TokenInfo{
    Type:     "identifier", 
    Category: "symbol",
})

// Classification returns rich metadata
if info, matches, err := table.Classify("if"); err == nil {
    fmt.Printf("Type: %s, Category: %s\n", info.Type, info.Category)
}
```

### Submatch Access

```go
table := helpers.NewRegexTable[string]()

// Pattern with capture groups
table.AddPattern(`(\d{4})-(\d{2})-(\d{2})`, "date")

if value, matches, err := table.Classify("2023-12-25"); err == nil {
    fmt.Printf("Full match: %s\n", matches[0])  // "2023-12-25"
    fmt.Printf("Year: %s\n", matches[1])        // "2023" 
    fmt.Printf("Month: %s\n", matches[2])       // "12"
    fmt.Printf("Day: %s\n", matches[3])         // "25"
}
```

## Implementation Notes

- Uses Go's built-in `regexp` package with named capture groups
- Auto-generates unique pattern names with reserved `__REGEXTABLE_` prefix  
- Compiles all patterns into a single union regex for optimal performance
- Supports lazy compilation to minimize overhead when adding multiple patterns
- Thread-safe for concurrent reads after compilation (not thread-safe for writes)
- Designed to be self-contained for potential library extraction
