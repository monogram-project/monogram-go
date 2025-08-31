# re-classify Tool

The `re-classify` tool is a command-line utility that classifies tokens based on configurable regular expression patterns. It reads tokens from standard input and outputs classification codes for each token.

## Usage

```bash
go run ./cmd/re-classify/ <config.yaml>
```

Or build and run:
```bash
go build -o re-classify ./cmd/re-classify/
./re-classify <config.yaml>
```

The tool reads tokens from standard input, **one token per line**, and outputs a single classification character for each token.

### Example

```bash
printf "if\nvariable\nend\n" | go run ./cmd/re-classify/ config.yaml
```

Or using echo with newlines:
```bash
echo -e "if\nvariable\nend" | go run ./cmd/re-classify/ config.yaml
```

## Classification Codes

The tool outputs single-character codes for each token:

- `S` - Start token (form start, e.g., `def`, `if`, `while`)
- `E` - End token (form end, e.g., `end`, `endif`, `endwhile`)
- `C` - Compound token (multi-part constructs)
- `L` - Label token (identifiers used as labels)
- `P` - Prefix token (operators that come before their operand)
- `O` - Operator token (infix, postfix operators)
- `V` - Variable token (default for unclassified identifiers)

For start tokens, the output may include expected end tokens:
- `S enddef endfunction` - Start token with possible endings

## Configuration File Format

The configuration file is in YAML format with the following structure:

### Basic Structure

```yaml
surround-regexp:
  - start: "start_pattern"
    endings: ["end_pattern_1", "end_pattern_2"]
  - start: "another_pattern"
    end: "single_end_pattern"

form-start-regexp:
  - "pattern1"
  - "pattern2"

form-end-regexp:
  - "end_pattern1"
  - "end_pattern2"

form-prefix-regexp:
  - "prefix_pattern"

simple-label-regexp:
  - "label_pattern"

compound-label-regexp:
  - "compound_pattern"

operator-regexp:
  - pattern: "operator_pattern"
    prefix-prec: 100
    infix-prec: 50
    postfix-prec: 75
    end-tokens: ["end1", "end2"]  # optional
```

### Pattern Types

#### 1. Surround Patterns (`surround-regexp`)

Surround patterns define start tokens and their corresponding end patterns. These support pattern substitution.

```yaml
surround-regexp:
  - start: "def"
    endings: ["enddef", "end"]
  - start: "if|while"
    endings: ["end$0", "$0_end"]  # $0 substitutes the matched text
  - start: "begin"
    end: "end"  # Single end pattern (alternative to endings array)
```

#### 2. Form Start Patterns (`form-start-regexp`)

Legacy patterns for identifying form start tokens:

```yaml
form-start-regexp:
  - "def"
  - "class"
  - "if|while|for"
```

#### 3. Form End Patterns (`form-end-regexp`)

Legacy patterns for identifying form end tokens:

```yaml
form-end-regexp:
  - "end.*"
  - "fi"
  - "done"
```

#### 4. Form Prefix Patterns (`form-prefix-regexp`)

Patterns for identifying prefix operators:

```yaml
form-prefix-regexp:
  - "not"
  - "\\+"  # Literal + character
```

#### 5. Simple Label Patterns (`simple-label-regexp`)

Patterns for identifying simple labels:

```yaml
simple-label-regexp:
  - "[a-z]+:"
  - "label[0-9]+"
```

#### 6. Compound Label Patterns (`compound-label-regexp`)

Patterns for identifying compound labels:

```yaml
compound-label-regexp:
  - "else-if"
  - "[a-z]+-[a-z]+"
```

#### 7. Operator Patterns (`operator-regexp`)

Operators with precedence values for prefix, infix, and postfix positions:

```yaml
operator-regexp:
  - pattern: "="
    prefix-prec: 0
    infix-prec: 1
    postfix-prec: 0
  - pattern: "\\+\\+"
    prefix-prec: 100
    infix-prec: 0
    postfix-prec: 75
    end-tokens: [")", "]"]  # Optional end tokens for form-start operators
```

### Pattern Substitution

Patterns in `surround-regexp` support substitution using:

- `$0` - The entire matched string

```yaml
surround-regexp:
  - start: "if|while"
    endings: ["end$0", "$0_end"]  # if -> endif, if_end; while -> endwhile, while_end
```

### Example Configuration

```yaml
surround-regexp:
  - start: "def"
    endings: ["enddef", "end"]
  - start: "if|while"
    endings: ["end$0", "$0_end"]
  - start: "begin"
    end: "end"

form-prefix-regexp:
  - "not"

simple-label-regexp:
  - "[a-z]+:"

operator-regexp:
  - pattern: "="
    prefix-prec: 0
    infix-prec: 1
    postfix-prec: 0
  - pattern: "\\+\\+"
    prefix-prec: 100
    infix-prec: 0
    postfix-prec: 75
```

### Testing

Create a simple test with the existing test configuration (one token per line):

```bash
printf "if\nvariable\nendif\n" | go run ./cmd/re-classify/ ./cmd/re-classify/test-config.yaml
```

Expected output:
```
S endif if_end
V
E
```

This shows:
- `if` classified as Start token (`S`) with expected endings `endif` and `if_end`
- `variable` classified as Variable (`V`)  
- `endif` classified as End token (`E`)

**Important**: The tool expects **one token per line** on standard input, not space-separated tokens. The number of output lines will always match the number of input lines.

## Implementation Notes

- The tool uses `RegexpTable` for efficient pattern matching
- Start tokens are processed first to build end token mappings
- Pattern matching is case-sensitive unless specified otherwise
- Capture groups in patterns enable flexible token transformation
- The tool processes all tokens in two phases for optimal classification accuracy
