# 0000 - Regexp based token classification, 2025-08-27

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

## Factors

- **Performance**: Regex matching needs to be efficient as it's applied to every identifier token
- **Configuration flexibility**: Patterns should be specifiable in easy to understand configuration files e.g. JSON, YAML
- **Regex engine limitations**: Go's standard `regexp` package lacks backreferences, which are needed for paired constructs, but is significantly more efficient than general alternatives
- **Maintainability**: The solution should be clean and avoid complex workarounds where possible
- **Backwards compatibility**: Existing behaviour is preserved
- **Classification override**: Manual patterns should optionally override automatic classifiers for matching identifiers
- **Syntactic cleanliness**: Enable syntax without required `:` and `!` markers when explicit classification is used
- **Language constraint**: Provide precise control over which identifiers are recognized for each syntactic role
- **Pattern types needed**:
  - **Simple Labels**: `simple-label-regex` - Identifies simple labels (e.g., "if", "else", "while")
  - **Compound Labels**: `compound-label-regex` - Identifies compound labels (e.g., "else-if", "else-while")
  - **Prefix Forms**: `form-prefix-regex` - Identifies prefix keywords (e.g., "return", "break", "continue")
  - **Surround Forms** (with three sub-classifiers):
    - `form-start-regex` - Identifies opening keywords (e.g., "function", "class", "struct")
    - `form-end-regex` - Identifies closing keywords (e.g., "endfunction", "endclass", "endstruct")
    - `form-surround-match` - Matches paired constructs (e.g., "if endif", "while endwhile", "function endfunction")

## Key Decisions

### Decision 1: Use Go's Built-in Regexp Package

**Rationale**: Despite considering many alternatives, we chose to stick with
Go's built-in `regexp` package for two critical reasons:

1. **Exposed Syntax**: Any regex solution must expose a concrete regexp syntax
   to users in configuration files. Using Go's standard syntax ensures
   familiarity and consistency.

2. **Performance**: Performance is critical for tokenization, and the RE2 engine
   underlying Go's regexp package provides predictable, linear-time performance.

**Future Consideration**: A hybrid solution could be explored in the future, as
`regexptable` supports pluggable regexp engines, potentially allowing fallback
to more feature-rich engines when needed.

### Decision 2: Piecemeal Override Architecture

**Rationale**: Manual configuration was designed to piecemeal override the
automatic classifiers for each role, rather than wholesale replacement of the
classification system.

**Implementation Challenge**: This required ensuring that late-stage automated
classification of labels was properly disabled when manual label classification
was enabled. The system needed to check for manual patterns first and only fall
back to automatic classification when no manual patterns matched.

**Benefit**: This approach preserves the existing automatic classification logic
while providing precise control over specific identifier roles.

### Decision 3: Redundant but Clear Surround-Forms Representation

**Rationale**: For representing the start/end matching rules of surround-forms,
we prioritized ease of explanation over eliminating redundancy.

**Design Choice**: Rather than trying to eliminate the redundancy between
separate start/end rules and matching rules, we chose to provide both:
- `form-start-regex` and `form-end-regex` for individual pattern matching  
- `form-surround-match` for paired construct validation

**Technical Challenge**: This approach required solving the issue of RE2 lacking
backreferences. We implemented manual backreference emulation by using capture
groups and verifying that all captured groups are equal.

**Benefit**: The redundant representation makes the configuration more intuitive
and self-documenting, even though it requires slightly more configuration.

### Implementation Details

1. **Pattern Compilation**: All regex patterns are compiled at startup using `regexptable.RegexpTableBuilder` for efficient runtime matching

2. **Configuration Structure**:
   ```yaml
   # Simple Labels classifier
   simple-label-regex: ["if", "else", "while", "for", "case", "default"]
   
   # Compound Labels classifier  
   compound-label-regex: ["else-if", "else-while"]
   
   # Prefix Forms classifier
   form-prefix-regex: ["return", "break", "continue"]
   
   # Surround Forms classifier (three sub-classifiers)
   form-start-regex: ["function", "class", "struct"]       # form-starts
   form-end-regex: ["endfunction", "endclass", "endstruct"] # form-ends
   form-surround-match: ["if endif", "while endwhile", "function endfunction"] # start/end matches
   ```

3. **Backreference Emulation**: For `form-surround-match` patterns, we use capture groups and manually verify that all captured groups are equal, effectively emulating backreferences:
   ```go
   func (tcc *TokenClassifiersCompiled) MatchesFormSurroundPattern(text string) bool {
       _, captures, ok := tcc.FormSurroundMatchTable.TryLookup(text)
       if ok && len(captures) > 1 {
           // All captured matches must be equal
           for n, m := range captures {
               if n == 0 { continue }
               if m != captures[0] { return false }
           }
       }
       return ok
   }
   ```

4. **Classification Logic**: The system provides four main classifiers that completely override automatic classification when matches occur:

   **During Tokenization (Initial Classification)**:
   - **Prefix Forms**: Manual `form-prefix-regex` patterns → `IdentifierFormPrefix` (overrides automatic)
   - **Surround Forms - Form Ends**: Manual `form-end-regex` patterns → `IdentifierFormEnd` (overrides automatic)
   - **Surround Forms - Form Starts**: Manual `form-start-regex` patterns → `IdentifierFormStart` (overrides automatic)
   - Automatic classification applies only when no manual patterns match

   **During Parsing (Context-Aware Classification)**:
   - **Simple Labels**: Manual `simple-label-regex` patterns → `IdentifierSimpleLabel` (overrides automatic)
   - **Compound Labels**: Manual `compound-label-regex` patterns → `IdentifierCompoundLabel` (overrides automatic)
   - **Surround Forms - Start/End Matching**: `form-surround-match` patterns validate paired constructs
   - Automatic context-based classification applies only when no manual patterns match
   - Fallback → `IdentifierVariable`

5. **Performance Optimization**: 
   - Patterns are pre-compiled into efficient lookup tables
   - Exact string matching where possible
   - Early termination on first match

### Trade-offs Made

**Explicit vs Automatic**: Chose to provide explicit classification as an
alternative to automatic classification with syntactic markers. This enables
cleaner syntax at the cost of requiring configuration.

**Manual vs Automatic**: Enhanced rather than replaced the existing automatic
classification system. Manual patterns provide explicit control while preserving
the intelligent automatic behavior for cases where syntactic markers are
acceptable.

**Performance vs Features**: Chose performance-optimized approach over full
regex feature support. The backreference emulation is a controlled compromise
that handles the specific use case needed.

**Simplicity vs Flexibility**: The exact matching approach limits some regex
features but provides predictable performance and simpler debugging.

**Dependencies vs Features**: Avoided heavy external regex dependencies in favor
of a lightweight solution with manual feature emulation where needed.

### Future Considerations

- A `regexp2re2` library could provide the ideal solution: full regex parsing
  with intelligent fallback to RE2 for performance
- Current solution is adequate for monogram's needs but may need revisiting for
  more complex regex requirements
- Pattern compilation could be cached to disk for very large pattern sets

## Additional Notes

- The implementation provides explicit classification as an alternative to
  syntactic marker-based automatic classification
- Enables cleaner syntax by eliminating the need for `:` and `!` markers when
  explicit patterns are used
- Provides precise language constraint capabilities for defining exactly which
  identifiers have special meaning
- All existing functionality is preserved - the system gracefully falls back to
  automatic classification when no manual patterns are specified
- The solution is well-tested and handles edge cases like malformed patterns
  gracefully
- Configuration is entirely optional - systems work exactly as before when no
  regex patterns are provided
