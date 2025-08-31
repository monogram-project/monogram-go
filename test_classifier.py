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
            # print("TOKEN", token, file=sys.stderr)
            sys.stdout.flush()

if __name__ == "__main__":
    main()
