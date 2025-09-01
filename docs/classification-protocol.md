# Protocol for External Classifiers

A classifier is any command that reads tokens from standard input, **one token
per line**, and outputs a single classification character for each token.


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

## The re-classify tool

This package comes with an easy-to-use but powerful classification tool 
called `re-classify`. You can read more about it in detail [here](../cmd/re-classify/README.md).


## Example of a Classifier (Python)

```py
#!/usr/bin/python3

import sys

# Simple test classifier that recognizes "if" as form-start with "fi" as end token
def classify_token(token):
    if token == "if":
        return "S fi"
    elif token == "fi":
        return "E"
    elif token == "then":
        return "L"  # Simple label
    else:
        return "V"  # Variable

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