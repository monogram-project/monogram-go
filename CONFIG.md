# Configuration File Support

The monogram tool now supports configuration files in YAML format using the `--config` (`-c`) flag.

## Usage

```bash
monogram --config config.yaml [other options]
```

## Configuration Options

The configuration file supports the following options:

### Regex Patterns (Future Use)
These patterns will be used for regex-based identifier classification:

- `simple-label-regex`: Pattern for simple labels (e.g., "if", "else", "while")
- `compound-label-regex`: Pattern for compound labels (e.g., "else-if")
- `form-start-regex`: Pattern for form start identifiers (e.g., "function", "class")
- `form-end-regex`: Pattern for form end identifiers (e.g., "endfunction", "endclass")
- `form-prefix-regex`: Pattern for prefix form identifiers (e.g., "return", "break")

### Default Options
These options provide defaults that are applied when the corresponding command-line flag is not specified:

- `default-format`: Default output format (xml, json, yaml, mermaid, dot)
- `default-indent`: Default indentation (number of spaces)
- `default-label`: Default label/breaker text
- `include-spans`: Include source spans in output (true/false)
- `decimal`: Decode numbers in base 10 (true/false)
- `check-literals`: Check regex and literal string validity (true/false)

## Example Configuration

```yaml
# Regex patterns for identifier classification
simple-label-regex: "^(if|else|while|for|case|default)$"
compound-label-regex: "^(else-if|else-while)$"
form-start-regex: "^(function|class|struct)$"
form-end-regex: "^end(function|class|struct)$"
form-prefix-regex: "^(return|break|continue)$"

# Default configuration values
default-format: "xml"
default-indent: 2
default-label: "seq"
include-spans: false
decimal: false
check-literals: true
```

## Precedence

Command-line flags always take precedence over configuration file settings. This allows you to:

1. Set common defaults in your config file
2. Override specific settings for individual runs using command-line flags

## Example Usage

```bash
# Use config defaults
echo 'hello world' | monogram --config my-config.yaml

# Override the format from config
echo 'hello world' | monogram --config my-config.yaml --format json

# Override indentation from config
echo 'hello world' | monogram --config my-config.yaml --indent 0
```
