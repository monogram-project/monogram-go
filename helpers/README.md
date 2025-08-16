# RegexClassifier

A high-performance multi-pattern regex classifier for Go using the built-in `regexp` package. This helper compiles multiple regex patterns into a single automaton for efficient pattern matching.

## Features

- **High Performance**: Uses a single compiled regex with named capture groups for O(n) matching regardless of pattern count
- **Type Safe**: Generic implementation supports any value type `T`
- **Reserved Namespace**: Uses `__REGEXCLASSIFIER_` prefix to avoid conflicts with user-defined capture groups
- **Built-in Regexp**: Uses Go's standard `regexp` package - no external dependencies
- **Full Match Access**: Returns both the classified value and complete submatch details
- **Self-Contained**: Designed to be extracted into a separate library

## Basic Usage

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
    // Create classifier
    classifier := helpers.NewRegexClassifier[TokenType]()
    
    // Add patterns
    classifier.AddPattern("keyword", `if|else|while|for`, TokenKeyword)
    classifier.AddPattern("identifier", `[a-zA-Z_][a-zA-Z0-9_]*`, TokenIdentifier) 
    classifier.AddPattern("number", `\d+`, TokenNumber)
    
    // Classify input
    if value, matches, err := classifier.Classify("if"); err == nil {
        fmt.Printf("Matched: %v\n", value) // TokenKeyword
        fmt.Printf("Text: %s\n", matches[0]) // "if"
    }
}
```

## API Reference

### Core Types

```go
type RegexClassifier[T any] struct {
    // Private fields
}
```

### Constructor

```go
func NewRegexClassifier[T any]() *RegexClassifier[T]
```

Creates a new empty classifier for the specified type `T`.

### Methods

#### AddPattern

```go
func (rc *RegexClassifier[T]) AddPattern(name, pattern string, value T) error
```

Adds a regex pattern with an associated value. The `name` must be unique.

**Parameters:**
- `name`: Unique identifier for the pattern
- `pattern`: Regular expression (regexp2 syntax)
- `value`: Value to return when this pattern matches

**Returns:** Error if name already exists or regex compilation fails.

#### RemovePattern

```go
func (rc *RegexClassifier[T]) RemovePattern(name string) error
```

Removes a pattern by name and recompiles the classifier.

#### Classify

```go
func (rc *RegexClassifier[T]) Classify(input string) (T, []string, error)
```

Attempts to match input against all patterns.

**Returns:**
- `T`: The value associated with the matched pattern
- `[]string`: Submatch slice with full match at index 0 and capture groups at subsequent indices
- `error`: Error if no patterns match or other issues

#### TryClassify

```go
func (rc *RegexClassifier[T]) TryClassify(input string) (T, []string, bool)
```

Like `Classify` but returns a boolean instead of an error.

**Returns:**
- `T`: The value (zero value if no match)
- `[]string`: Submatch slice (nil if no match)  
- `bool`: True if a pattern matched

#### HasPatterns

```go
func (rc *RegexClassifier[T]) HasPatterns() bool
```

Returns true if any patterns are configured.

## Advanced Usage

### Using Match Details

```go
classifier := helpers.NewRegexClassifier[string]()
classifier.AddPattern("capture", `(\w+):(\d+)`, "key-value")

if value, matches, err := classifier.Classify("name:42"); err == nil {
    fmt.Printf("Type: %s\n", value)     // "key-value"
    fmt.Printf("Full: %s\n", matches[0]) // "name:42"
    
    // Access capture groups
    fmt.Printf("Key: %s\n", matches[1])   // "name"
    fmt.Printf("Value: %s\n", matches[2]) // "42"
}
```

### Error Handling

```go
// Method 1: Using Classify with error handling
if value, match, err := classifier.Classify(input); err != nil {
    switch {
    case strings.Contains(err.Error(), "no patterns configured"):
        // Handle empty classifier
    case strings.Contains(err.Error(), "no pattern matched"):
        // Handle no match
    default:
        // Handle other errors
    }
}

// Method 2: Using TryClassify for simple success/failure
if value, matches, ok := classifier.TryClassify(input); ok {
    // Handle successful match
} else {
    // Handle no match
}
```

## Performance Characteristics

- **Compilation**: O(P) where P = number of patterns
- **Matching**: O(N) where N = input length (independent of pattern count)
- **Memory**: O(P) for pattern storage + compiled automaton size

The key advantage is that matching time is independent of the number of patterns, making it ideal for scenarios with many patterns.

## Thread Safety

`RegexClassifier` is **not** thread-safe for modifications (`AddPattern`, `RemovePattern`) but is safe for concurrent reads (`Classify`, `TryClassify`) once patterns are configured.

For concurrent modification, use external synchronization:

```go
var mu sync.RWMutex
var classifier = helpers.NewRegexClassifier[TokenType]()

// Modify (exclusive)
mu.Lock()
classifier.AddPattern("new", `pattern`, value)
mu.Unlock()

// Read (shared)
mu.RLock()
value, matches, err := classifier.Classify(input)
mu.RUnlock()
```
