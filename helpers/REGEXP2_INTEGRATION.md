# Adding regexp2 Engine Support

This guide explains how to integrate Microsoft's `regexp2` library as an alternative regex engine for RegexTable. The `regexp2` library provides .NET-compatible regular expressions with advanced features like lookbehind, named groups, and other functionality not available in Go's standard `regexp` package.

## Important: External Integration Only

**⚠️ CRITICAL DESIGN DECISION**: The `regexp2` library should **NOT** be added as a dependency to the helpers package itself. This keeps the helpers package lightweight and free of external dependencies, making it suitable for extraction into its own repository.

Instead, `regexp2` integration should be implemented in consuming applications or as a separate companion package when the helpers are moved to their own repository.

## Prerequisites

In your **consuming application** (not in the helpers package), add the regexp2 dependency:

```bash
go get github.com/dlclark/regexp2
```

The helpers package itself should remain dependency-free except for Go's standard library.

## Implementation Overview

The regexp2 engine integration involves:
1. Implementing the `RegexEngine` interface for regexp2 **in your application code**
2. Implementing the `CompiledRegex` interface for regexp2 compiled patterns
3. Handling differences in API and feature set
4. **Keeping the helpers package dependency-free**

## Integration Strategies

### Strategy 1: Application-Level Integration (Recommended)

Create the regexp2 engine implementation in your main application:

```
your-app/
├── main.go
├── regexp2_engine.go    // Your regexp2 implementation
├── go.mod               // Contains regexp2 dependency
└── vendor/ or go.mod    // External dependencies here
```

### Strategy 2: Separate Companion Package

When helpers are moved to their own repo, create a companion package:

```
github.com/your-org/regextable-helpers     // Core package (no deps)
github.com/your-org/regextable-regexp2     // Regexp2 integration
```

## Step 1: Create the Regexp2Engine in Your Application

Create `regexp2_engine.go` **in your application** (not in helpers/):

```go
// File: your-app/regexp2_engine.go
package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/monogram-project/monogram-go/helpers" // Import helpers
)

// Regexp2Engine implements helpers.RegexEngine using Microsoft's regexp2 library.
// This provides .NET-compatible regex functionality with advanced features
// like lookbehind, balancing groups, and conditional expressions.
type Regexp2Engine struct {
	options regexp2.RegexOptions
}

// NewRegexp2Engine creates a new regexp2 engine with default options.
func NewRegexp2Engine() *Regexp2Engine {
	return &Regexp2Engine{
		options: regexp2.None,
	}
}

// NewRegexp2EngineWithOptions creates a new regexp2 engine with custom options.
func NewRegexp2EngineWithOptions(options regexp2.RegexOptions) *Regexp2Engine {
	return &Regexp2Engine{
		options: options,
	}
}

// Compile compiles a regex pattern using regexp2.
func (e *Regexp2Engine) Compile(pattern string) (helpers.CompiledRegex, error) {
	compiled, err := regexp2.Compile(pattern, e.options)
	if err != nil {
		return nil, err
	}
	return NewRegexp2CompiledRegex(compiled), nil
}

// FormatNamedGroup formats a named capture group using .NET syntax (?<name>pattern).
func (e *Regexp2Engine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?<%s>%s)", groupName, pattern)
}
```

## Step 2: Create the Compiled Regex Wrapper

Add to the same application file:

```go
// Regexp2CompiledRegex wraps a regexp2.Regexp to implement helpers.CompiledRegex.
type Regexp2CompiledRegex struct {
	regexp *regexp2.Regexp
}

// NewRegexp2CompiledRegex creates a new Regexp2CompiledRegex.
func NewRegexp2CompiledRegex(regexp *regexp2.Regexp) *Regexp2CompiledRegex {
	return &Regexp2CompiledRegex{regexp: regexp}
}

// FindStringSubmatch finds the first match and returns all submatches.
// This adapts regexp2's API to match Go's standard regexp interface.
func (r *Regexp2CompiledRegex) FindStringSubmatch(s string) []string {
	match, err := r.regexp.FindStringMatch(s)
	if err != nil || match == nil {
		return nil
	}

	// Convert regexp2.Match to []string format expected by RegexTable
	groups := match.Groups()
	result := make([]string, len(groups))
	
	for i, group := range groups {
		if group.Length > 0 {
			result[i] = group.String()
		} else {
			result[i] = ""
		}
	}
	
	return result
}

// SubexpNames returns the names of capturing groups.
// Note: regexp2 doesn't provide a direct equivalent to Go's SubexpNames,
// so we need to extract this information differently.
func (r *Regexp2CompiledRegex) SubexpNames() []string {
	// regexp2 doesn't expose group names the same way as Go's regexp.
	// We'd need to track them separately or parse them from the pattern.
	// For RegexTable's use case, we can return a slice with empty strings
	// since RegexTable generates its own group names.
	groupCount := r.regexp.GetGroupNumbers()
	names := make([]string, len(groupCount))
	// First element is always empty (represents the entire match)
	for i := range names {
		names[i] = ""
	}
	return names
}
```

## Step 3: Handle API Differences

The main challenges when adapting regexp2 are:

### 3.1 Different Match API
- Go's `regexp`: `FindStringSubmatch(s string) []string`
- regexp2: `FindStringMatch(s string) (*Match, error)`

### 3.2 Group Name Handling
- Go's `regexp`: Provides `SubexpNames()` method
- regexp2: Doesn't expose group names in the same way

### 3.3 Error Handling
- Go's `regexp`: Panics on invalid patterns during compilation
- regexp2: Returns errors during both compilation and matching

## Step 4: Usage Examples

### Basic Usage

```go
// File: your-app/main.go
package main

import (
    "fmt"
    "github.com/monogram-project/monogram-go/helpers"
)

func main() {
    // Create RegexTable with regexp2 engine (defined in same package)
    engine := NewRegexp2Engine()
    table := helpers.NewRegexTableBuilderWithEngine[string](engine).
        AddPattern(`\b(?:if|else|while|for)\b`, "keyword").
        AddPattern(`\b[a-zA-Z_]\w*\b`, "identifier").
        AddPattern(`\d+`, "number").
        MustBuild()

    // Test with advanced regexp2 features
    testCases := []string{"if", "myVariable", "123", "unknown"}
    for _, test := range testCases {
        if value, _, found := table.TryLookup(test); found {
            fmt.Printf("%s -> %s\n", test, value)
        } else {
            fmt.Printf("%s -> no match\n", test)
        }
    }
}
```

### Using Advanced regexp2 Features

```go
// Using lookbehind (not available in Go's standard regexp)
engine := NewRegexp2Engine() // Your local implementation
table := helpers.NewRegexTableBuilderWithEngine[string](engine).
    AddPattern(`(?<=\bclass\s+)\w+`, "class_name").        // Lookbehind
    AddPattern(`\w+(?=\s*\()`, "function_name").           // Lookahead
    AddPattern(`\b\w+\b`, "identifier").
    MustBuild()

// Test with lookbehind/lookahead patterns
testInput := "class MyClass extends BaseClass { function doSomething() }"
// This would match "MyClass" as "class_name" and "doSomething" as "function_name"
```

### With Custom Options

```go
import "github.com/dlclark/regexp2"

// Case-insensitive matching
engine := NewRegexp2EngineWithOptions(regexp2.IgnoreCase) // Your local implementation
table := helpers.NewRegexTableBuilderWithEngine[string](engine).
    AddPattern(`html?`, "markup").
    AddPattern(`css`, "stylesheet").
    MustBuild()

// Will match "HTML", "html", "CSS", "css", etc.
```

## Step 5: Testing Your Implementation

Create tests **in your application** (not in helpers/):

```go
// File: your-app/regexp2_engine_test.go
package main

import (
    "testing"
    "github.com/dlclark/regexp2"
    "github.com/monogram-project/monogram-go/helpers"
)

func TestRegexp2Engine_Basic(t *testing.T) {
    engine := NewRegexp2Engine()
    
    // Test compilation
    compiled, err := engine.Compile(`\d+`)
    if err != nil {
        t.Fatalf("Failed to compile pattern: %v", err)
    }
    
    // Test matching
    matches := compiled.FindStringSubmatch("abc123def")
    if matches == nil || matches[0] != "123" {
        t.Errorf("Expected match '123', got %v", matches)
    }
}

func TestRegexp2Engine_NamedGroups(t *testing.T) {
    engine := NewRegexp2Engine()
    formatted := engine.FormatNamedGroup("test", "\\d+")
    expected := "(?<test>\\d+)"
    
    if formatted != expected {
        t.Errorf("Expected %q, got %q", expected, formatted)
    }
}

func TestRegexp2Engine_WithRegexTable(t *testing.T) {
    engine := NewRegexp2Engine() // Your local implementation
    table := helpers.NewRegexTableBuilderWithEngine[string](engine).
        AddPattern(`\d+`, "number").
        AddPattern(`[a-z]+`, "word").
        MustBuild()
    
    testCases := []struct {
        input    string
        expected string
        found    bool
    }{
        {"123", "number", true},
        {"hello", "word", true},
        {"HELLO", "", false}, // Case sensitive by default
        {"@#$", "", false},
    }
    
    for _, tc := range testCases {
        value, _, found := table.TryLookup(tc.input)
        if found != tc.found || value != tc.expected {
            t.Errorf("Input %q: expected (%q, %t), got (%q, %t)", 
                tc.input, tc.expected, tc.found, value, found)
        }
    }
}

func TestRegexp2Engine_AdvancedFeatures(t *testing.T) {
    // Test features not available in Go's standard regexp
    engine := NewRegexp2Engine()
    
    // Test lookbehind (if your patterns use it)
    compiled, err := engine.Compile(`(?<=hello\s+)\w+`)
    if err != nil {
        t.Fatalf("Failed to compile lookbehind pattern: %v", err)
    }
    
    matches := compiled.FindStringSubmatch("hello world")
    if matches == nil || matches[0] != "world" {
        t.Errorf("Lookbehind failed: expected 'world', got %v", matches)
    }
}
```

## Step 6: Performance Considerations

### Memory Usage
- regexp2 typically uses more memory than Go's standard regexp
- Consider this for applications with many patterns or high throughput

### Compilation Speed
- regexp2 compilation can be slower than Go's regexp
- The lazy compilation feature of RegexTable helps mitigate this

### Runtime Performance
- regexp2 may be slower for simple patterns
- But can be faster for complex patterns with advanced features
- Benchmark your specific use cases

## Step 7: Error Handling Differences

### Compilation Errors
```go
// regexp2 can return errors during matching, not just compilation
func (r *Regexp2CompiledRegex) FindStringSubmatch(s string) []string {
    match, err := r.regexp.FindStringMatch(s)
    if err != nil {
        // Handle runtime errors (malformed input, stack overflow, etc.)
        // For RegexTable compatibility, we return nil on any error
        return nil
    }
    // ... rest of implementation
}
```

### Pattern Validation
```go
// Add validation helper for regexp2 patterns
func ValidateRegexp2Pattern(pattern string) error {
    _, err := regexp2.Compile(pattern, regexp2.None)
    return err
}
```

## Step 8: Feature Comparison

| Feature | Go regexp | regexp2 | Notes |
|---------|-----------|---------|-------|
| Basic matching | ✅ | ✅ | Both support standard regex |
| Named groups | ✅ | ✅ | Different syntax |
| Lookbehind | ❌ | ✅ | Major advantage of regexp2 |
| Balancing groups | ❌ | ✅ | Advanced .NET feature |
| Performance | ⚡ Fast | 🐌 Slower | For simple patterns |
| Memory usage | 💚 Low | 🟡 Higher | regexp2 uses more memory |
| Unicode support | ✅ | ✅ | Both have good Unicode support |

## Best Practices

1. **Keep helpers dependency-free** - Never add regexp2 as a dependency to the helpers package itself
2. **Use regexp2 when you need advanced features** like lookbehind that aren't available in Go's regexp
3. **Implement in your application** - Keep the regexp2 integration code in your consuming application
4. **Profile your application** to ensure regexp2's performance characteristics work for your use case
5. **Handle errors gracefully** since regexp2 can fail during matching, not just compilation
6. **Test thoroughly** with your specific patterns, especially if using advanced features
7. **Consider a hybrid approach** using Go's regexp for simple patterns and regexp2 for complex ones
8. **Plan for extraction** - When helpers move to their own repo, consider a separate regexp2 companion package

## Package Structure Examples

### Current Structure (In monogram-go)
```
monogram-go/
├── helpers/                    # No external dependencies
│   ├── regextable.go
│   ├── regex_engine.go
│   └── standard_regex_engine.go
└── your-app/
    ├── main.go
    ├── regexp2_engine.go       # Your regexp2 implementation
    └── go.mod                  # Contains regexp2 dependency
```

### Future Structure (Extracted)
```
github.com/your-org/regextable/
├── regextable.go               # Core package (no deps)
├── regex_engine.go
└── standard_regex_engine.go

github.com/your-org/regextable-regexp2/
├── regexp2_engine.go           # Companion package
├── go.mod                      # Contains regexp2 dependency
└── README.md

your-application/
├── main.go
└── go.mod                      # Imports both packages
```

## Troubleshooting Common Issues

### 1. Import Conflicts
```go
// Use aliases to avoid conflicts
import (
    stdregexp "regexp"
    "github.com/dlclark/regexp2"
)
```

### 2. Performance Issues
```go
// Cache compiled patterns aggressively
var patternCache = make(map[string]*regexp2.Regexp)

func getCachedPattern(pattern string) (*regexp2.Regexp, error) {
    if cached, exists := patternCache[pattern]; exists {
        return cached, nil
    }
    
    compiled, err := regexp2.Compile(pattern, regexp2.None)
    if err == nil {
        patternCache[pattern] = compiled
    }
    return compiled, err
}
```

### 3. Group Name Limitations
Since regexp2 doesn't expose group names the same way as Go's regexp, RegexTable's auto-generated group names work perfectly - this is actually an advantage as it avoids the complexity of mapping between different group naming systems.

## Why Keep Helpers Dependency-Free?

### Benefits of Zero External Dependencies

1. **Easy Extraction**: The helpers package can be moved to its own repository without bringing external dependencies
2. **Broad Compatibility**: No version conflicts with applications that use different versions of regexp2
3. **Lightweight**: Keeps the core functionality minimal and focused
4. **Optional Features**: Advanced regex features remain optional and don't bloat the core package
5. **Clear Separation**: Business logic (your app) vs utility code (helpers) boundaries are maintained

### The Plugin Architecture Advantage

The `RegexEngine` interface acts as a plugin system:
- **Core package**: Provides the interface and standard implementation
- **Your application**: Implements advanced engines as needed
- **Future packages**: Can provide pre-built engines without affecting core

This is a common pattern in Go libraries - provide interfaces and let users implement or import implementations as needed.

## Real-World Example

Here's a complete working example showing the recommended integration approach:

```go
// File: your-app/main.go
package main

import (
	"fmt"
	"log"
	"github.com/dlclark/regexp2"
	"github.com/monogram-project/monogram-go/helpers"
)

// Your regexp2 engine implementation (in your app, not in helpers)
type Regexp2Engine struct {
	options regexp2.RegexOptions
}

func NewRegexp2Engine() *Regexp2Engine {
	return &Regexp2Engine{options: regexp2.None}
}

func (e *Regexp2Engine) Compile(pattern string) (helpers.CompiledRegex, error) {
	compiled, err := regexp2.Compile(pattern, e.options)
	if err != nil {
		return nil, err
	}
	return &Regexp2CompiledRegex{regexp: compiled}, nil
}

func (e *Regexp2Engine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?<%s>%s)", groupName, pattern)
}

type Regexp2CompiledRegex struct {
	regexp *regexp2.Regexp
}

func (r *Regexp2CompiledRegex) FindStringSubmatch(s string) []string {
	match, err := r.regexp.FindStringMatch(s)
	if err != nil || match == nil {
		return nil
	}
	
	groups := match.Groups()
	result := make([]string, len(groups))
	for i, group := range groups {
		if group.Length > 0 {
			result[i] = group.String()
		}
	}
	return result
}

func (r *Regexp2CompiledRegex) SubexpNames() []string {
	groupCount := r.regexp.GetGroupNumbers()
	names := make([]string, len(groupCount))
	return names
}

func main() {
	// Use advanced regexp2 features with RegexTable
	engine := NewRegexp2Engine()
	
	table, err := helpers.NewRegexTableBuilderWithEngine[string](engine).
		AddPattern(`(?<=\bclass\s+)\w+`, "class_name").        // Lookbehind!
		AddPattern(`\w+(?=\s*\()`, "function_name").           // Lookahead!
		AddPattern(`\b[A-Z]\w*\b`, "constant").
		Build()
	
	if err != nil {
		log.Fatal(err)
	}
	
	testCode := "class MyService extends BaseService { function process() { const MAX_SIZE = 100; } }"
	
	// This showcases regexp2's advanced features working with RegexTable
	words := []string{"MyService", "BaseService", "process", "MAX_SIZE"}
	for _, word := range words {
		if value, _, found := table.TryLookup(word); found {
			fmt.Printf("%-12s -> %s\n", word, value)
		}
	}
}
```

This approach keeps the helpers package clean and dependency-free while still enabling powerful regexp2 integration in your applications.
