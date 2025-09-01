# 0000 - External Classifier Integration, 2025-09-01
_Supersedes regexp-based classification approach from 2025-08-27_

## Issue

The monogram tokenizer needed to support manual/explicit identifier
classification to address two key design goals:

1. **Language Constraint**: Provide a simple way to constrain the recognized
   language by explicitly defining which identifiers should be classified as
   keywords, labels, or special forms, rather than relying solely on automatic
   detection.

2. **Syntactic Marker Elimination**: Enable elimination of the syntactic markers
   `:` and `!` that are currently required for automatic classification (e.g.,
   `else:` for simple labels, `return!` for form prefixes), allowing cleaner
   syntax when explicit classification is preferred.

The system already had sophisticated automatic classification that could infer
token roles based on context and syntactic markers, but explicit classification
provides an alternative approach that can simplify the surface syntax while
giving precise control over recognized syntax.

## Decision Evolution

**Initial Approach (2025-08-27)**: A regex-based classification system was
designed and fully implemented, featuring multiple pattern types
(`simple-label-regex`, `compound-label-regex`, `form-prefix-regex`, etc.) with
sophisticated matching rules and backreference emulation.

**Pivot Decision (2025-09-01)**: After implementation and evaluation, the regex
approach was determined to be somewhat arbitrary and potentially limiting for
users with sophisticated classification needs. A more flexible external
classifier approach was adopted instead.

## Key Decision: External Classifier Architecture

**Rationale**: Rather than constraining users to regex pattern matching, we
chose to delegate classification to external programs via the `--use-classifier`
flag, providing full flexibility for classification logic.

**Benefits**:
- **Full Flexibility**: External classifiers can use any language, algorithms, or techniques for classification
- **Extensibility**: Classification logic can evolve independently of monogram releases

## Architecture Overview

### Core Grammar Preservation

**Essential Syntactic Rules**: Monogram's classification by discovery remains
is used when `--use-classifier` is not specified:.

**External Classification**: Classification of identifiers and signs is
delegated to external programs when `--use-classifier` is specified.

The protocol is described in 

### Configuration Integration

**Command-Line Flag**: 
```go
type FormatOptions struct {
    UseClassifier string  // External classifier command
    // ... other options
}
```

**Config File Support**:
```yaml
option-use-classifier: "python classifier.py"
```

**Validation**: Empty classifier command is treated as an error when explicitly provided:
```bash
monogram --use-classifier=""  # Error: command required
```

### Implementation Details

1. **Flag Integration**: The `--use-classifier` flag is integrated into the main `FormatOptions` structure and config system

2. **External Process Management**: When specified, an external classifier process is spawned and managed for the duration of tokenization

3. **Token Communication**: Tokens are sent to the external classifier, which returns classification decisions

4. **Fallback Behavior**: When no external classifier is specified, the system uses built-in automatic classification rules

## Trade-offs Made

**Flexibility vs Simplicity**: Chose external process complexity over the limitations of regex patterns. While external processes add operational complexity, they provide unlimited flexibility for classification logic.

**Arbitrary Patterns vs Sophisticated Logic**: Moved away from potentially arbitrary regex patterns toward user-defined classification logic that can be as sophisticated as needed.

**Built-in vs External**: Accepted the overhead of external processes in exchange for keeping the core system simple and allowing classification to evolve independently.

**Configuration Complexity vs Implementation Power**: Traded complex regex configuration for the power to implement any classification approach in any language.

## Removed Components

The following regex-based classification components were removed:

- `TokenClassifiers` and `TokenClassifiersCompiled` structures
- `simple-label-regex`, `compound-label-regex`, `form-prefix-regex` patterns
- `form-start-regex`, `form-end-regex`, `form-surround-match` patterns  
- Complex backreference emulation logic
- `regexptable` dependency and compilation logic

These were replaced with simple external classifier delegation.

## Migration Impact

**Existing Functionality**: All existing automatic classification behavior is preserved when no `--use-classifier` flag is provided.

**New Capability**: Users can now specify external classification programs for advanced use cases.

**Simplification**: The core codebase is significantly simpler without the complex regex pattern matching system.

## Future Considerations

- External classifiers can be written in any language and use any classification approach
- Standard classifier implementations could be provided as examples
- Classification protocols could be standardized for interoperability
- Performance optimizations could include classifier process reuse across multiple files

## Additional Notes

- The external classifier approach provides maximum flexibility while keeping the core system simple
- All built-in automatic classification rules are preserved for backward compatibility
- The `--use-classifier` flag is optional - systems work exactly as before when not specified
- External classifier commands are validated to prevent empty command errors
- The approach enables sophisticated classification without adding complexity to the core tokenizer
