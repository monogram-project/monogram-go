package helpers

// RegexEngine defines the minimal interface needed by RegexTable for regex operations.
// This abstraction allows different regex engines to be used with RegexTable.
type RegexEngine interface {
	// Compile compiles a regex pattern and returns a CompiledRegex or an error.
	Compile(pattern string) (CompiledRegex, error)

	// FormatNamedGroup formats a pattern with a named capture group using the engine's syntax.
	// For example: Go uses (?P<name>pattern), .NET uses (?<name>pattern), etc.
	FormatNamedGroup(groupName, pattern string) string
}

// CompiledRegex represents a compiled regex pattern that can perform matches.
type CompiledRegex interface {
	// FindStringSubmatch finds the first match and returns all submatches.
	// Returns nil if no match is found.
	// The first element is the full match, subsequent elements are capture groups.
	FindStringSubmatch(s string) []string

	// SubexpNames returns the names of the capturing groups.
	// The first element corresponds to the entire match and is empty.
	SubexpNames() []string
}
