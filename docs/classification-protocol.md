# Protocol for External Classifiers

A classifier is any command that reads tokens from standard input, **one token
per line**, and outputs a single classification character for each token.


## Classification Codes

For each token, the tool outputs a line starting with a single-character code:

- `S` - Start token (form start, e.g., `def`, `if`, `while`)
- `E` - End token (form end, e.g., `end`, `endif`, `endwhile`)
- `C` - Compound token (multi-part constructs)
- `L` - Label token (identifiers used as labels)
- `P` - Prefix token (operators that come before their operand)
- `O` - Operator token (infix, postfix operators)
- `V` - Variable token (identifiers used as variables)
- `U` - Unclassified (continue with the initially assigned role)

Form-start tokens and operator tokens are followed by additional information:

- For start tokens, the output is followed by the possible matching end tokens:
  e.g. `if` might map into `S end endif`
- Operators tokens are followed by their prefix, infix and postfix
  precedences. Note that 0 indicates that they don't have that role.
  e.g. `O 5 15 0` means an operator which can be used in prefix and
  infix roles but not postfix roles.


## Example of a Classifier (Python)

This is a simple implementation of a classfier in Python.

```py
#!/usr/bin/python3

import sys

# Simple test classifier that recognizes "if" as form-start with "fi" as end token
def classify_token(token):
    if token == "if":
        return "S fi"   # Form start
    elif token == "fi":
        return "E"      # Form-end
    elif token == "then":
        return "L"      # Simple label
    else:
        return "V"      # Variable

def main():
    for line in sys.stdin:
        token = line.strip()
        if token:
            classification = classify_token(token)
            print(classification)
            sys.stdout.flush()

if __name__ == "__main__":
    main()
```

## See Also

This package comes with an easy-to-use but powerful classification tool 
called `re-classify`. You can read more about it in detail [here](../cmd/re-classify/README.md).
