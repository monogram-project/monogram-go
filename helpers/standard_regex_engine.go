package helpers

import (
	"fmt"
	"regexp"
)

// StandardRegexEngine implements RegexEngine using Go's built-in regexp package.
type StandardRegexEngine struct{}

// NewStandardRegexEngine creates a new StandardRegexEngine.
func NewStandardRegexEngine() *StandardRegexEngine {
	return &StandardRegexEngine{}
}

// Compile compiles a regex pattern using Go's regexp.Compile.
func (e *StandardRegexEngine) Compile(pattern string) (CompiledRegex, error) {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return NewStandardCompiledRegex(compiled), nil
}

// FormatNamedGroup formats a named capture group using Go's (?P<name>pattern) syntax.
func (e *StandardRegexEngine) FormatNamedGroup(groupName, pattern string) string {
	return fmt.Sprintf("(?P<%s>%s)", groupName, pattern)
}

// StandardCompiledRegex wraps a Go *regexp.Regexp to implement CompiledRegex.
type StandardCompiledRegex struct {
	regexp *regexp.Regexp
}

// NewStandardCompiledRegex creates a new StandardCompiledRegex wrapping the given regexp.
func NewStandardCompiledRegex(regexp *regexp.Regexp) *StandardCompiledRegex {
	return &StandardCompiledRegex{regexp: regexp}
}

// FindStringSubmatch delegates to the wrapped regexp.
func (r *StandardCompiledRegex) FindStringSubmatch(s string) []string {
	return r.regexp.FindStringSubmatch(s)
}

// SubexpNames delegates to the wrapped regexp.
func (r *StandardCompiledRegex) SubexpNames() []string {
	return r.regexp.SubexpNames()
}
