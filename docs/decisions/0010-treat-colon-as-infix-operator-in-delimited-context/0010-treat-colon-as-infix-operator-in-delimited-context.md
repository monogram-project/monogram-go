# 0010 - Treat colon as infix operator in delimited context, 2025-03-01

## Issue

Because Python is a relatively strong influence on the design of Monogram, the
`:` keyword is already overloaded for its role in named and anonymous breakers.
However breakers can only occur in a form-context and not in a
delimited-context. 

And there is a specific delimited-context in which colon-as-operator is
particularly useful; namely when we are imitating JSON. At this point in the
development of Monogram it really is 100% obvious whether or not that is an
important use case. However it seems very likely that it is a good idea.

## Factors

- Ambiguity
- Ease of learning

## Options and Outcome

1. Treat colon as an operator in delimited-context
2. Limit colon to its role as a breaker-marker

*Outcome*: Option 1

## Consequences

- Keeps open the option of JSON as a subset
- Now `k: v` means different things in different syntactic contexts
  (disambiguated largely by newlines in practice)

## Pros and Cons of Options

### Option 1

- Pros
    - JSON is potentially a strict subset
- Cons
    - We lose the flexibility of allowingh breakers in delimited-context (which
      might support 'list comprehension'-style syntax.)
- Interesting
    - To align with JSON there is an impact on tokenisation, although
      I think it is an acceptable compromise.
    - Python's 'list-comprehension' syntax is still unavailable.

### Option 2

- Pros
  - Easier to teach
  - Retains flexibility of breakers in delimited-context
- Cons
    - Incompatible with JSON

## Additional Notes

*Contexts*:
- Every region of code in a monogram document is nested inside delimiters and forms.
- The closest level of nesting determines the _context_.
- At top-level code might not be nested at all. In which case we say it is in
  _document_ context.
