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

The protocol for an external classifier was deliberately kept simple. A 
classifier accepts a series of tokens, one token per line. And, for each input
token, the tool generates a line starting with a one-character classification.

- `S` - Start token (form start, e.g., `def`, `if`, `while`)
- `E` - End token (form end, e.g., `end`, `endif`, `endwhile`)
- `C` - Compound token (multi-part constructs)
- `L` - Label token (identifiers used as labels)
- `P` - Prefix token (operators that come before their operand)
- `O` - Operator token (infix, postfix operators)
- `V` - Variable token (default for unclassified identifiers)

Form-start tokens and operator tokens are followed by additional information:

- For start tokens, the output is followed by the possible matching end tokens:
  e.g. `if` might map into `S end endif`
- Operators tokens are followed by their prefix, infix and postfix
  precedences. Note that 0 indicates that they don't have that role.
  e.g. `O 5 15 0` means an operator which can be used in prefix and
  infix roles but not postfix roles.

Note that a classifier may buffer up all the input before generate any output -
or equally may generate output in lockstep with the input. 


## Architecture Overview

### Core Grammar Preservation

**Essential Syntactic Rules**: Monogram's classification by discovery remains
is used when `--use-classifier` is not specified:.

**External Classification**: Classification of identifiers and signs is
delegated to external programs when `--use-classifier` is specified.


## Migration Impact

**Existing Functionality**: All existing automatic classification behavior is
preserved when no `--use-classifier` flag is provided.

**Provided Classification**: For many scenarios the provided `re-classify` is
more than adequate but still easy to configure.

**New Capability**: But users can now write their own external classification
programs for advanced use cases.

## Additional Notes

- The external classifier approach provides maximum flexibility while keeping the core system simple
- All built-in automatic classification rules are preserved for backward compatibility
- The `--use-classifier` flag is optional - systems work exactly as before when not specified
- External classifier commands are validated to prevent empty command errors
- The approach enables sophisticated classification without adding complexity to the core tokenizer
