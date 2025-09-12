# Configuration File Support

The monogram tool now supports configuration files in YAML format using the
`--config` (`-c`) flag.

## Usage

```bash
monogram --config config.yaml [other options]
```

## Configuration Options

The configuration file supports the following options:

- `option-format`: Default output format (xml, json, yaml, mermaid, dot)
- `option-indent`: Default indentation (number of spaces)
- `option-default-label`: Default label text
- `option-include-spans`: Include source spans in output (true/false)
- `option-decimal`: Decode numbers in base 10 (true/false)
- `option-check-literals`: Check regex and literal string validity (true/false)
- `option-use-classifier`: Optional command to use to classify tokens.
- `option-trim-token-on-output`: Trims the width of tokens to a maximum for
  display purposes (e.g. with Mermaid output)


## Example Configuration

```yaml
# Default configuration values
option-format: "xml"
option-indent: 2
option-default-label: "seq"
option-include-spans: false
option-decimal: false
option-check-literals: true
```

## Precedence

Command-line flags always take precedence over configuration file settings. This
allows you to:

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
