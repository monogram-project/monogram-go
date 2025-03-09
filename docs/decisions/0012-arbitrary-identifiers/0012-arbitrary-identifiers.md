# 0012 - Arbitrary identifiers, 2025-03-04

## Issue

The use of the `end` prefix to mark form keywords does have the potential
for awkward name clashes. In addition, we might want to have completely
arbitrary names for identifiers.

## Factors

- Syntax should be easy to learn and easy to understand
- Should be simple to use for identifers whose name starts with `end`
- Must allow for completely arbitrary identifiers

## Options and Outcome

- Option 1: `\endfoo`, using backslash to escape an identifier
- Option 2: `\"foo+bar"`, escaping a string
- Option 3: `_endfoo_`, `_foo\+bar_`, using underbars as delimiting quote-marks 
  and requiring escape sequences.

**Outcome**: Option 3

## Consequences

- Underbar has a special status
- The tokenisation rules are made more complicated to teachh
- We can handle arbitrary identifiers
- And the backslash is free to use for raw-strings.

## Pros and Cons of Options

### Option 1

- Pros
    - Easy to learn and understand
    - Handles form-ends neatly
- Cons
    - Cannot handle any other exceptions
- Interesting
    - Can be combined with Option 2 for improved coverage

### Option 2

- Pros
    - Easy to learn and understand
    - Handles arbitrary identifiers
- Cons
    - Competes with the use of `\` in raw-string syntax

### Option 3

- Pros
    - Easy to use
    - Handles arbitrary identifiers
    - No competition for the syntax
    - Does not consume the symbol space
- Cons
    - The situation where an underbar is only at one end of an identifier
      needs to be handled differently, which complicates the understanding.
    - The fix is to require all non-alphanumerics to be escaped.
    - This only works because the only non-alphabetic escape sequences are
      for interpolation, which we do not need to support.

### Additional Notes

The driving use-case for this decision was to permit variables such as
`endoscopy`. This would normally be seen as a form-end matching `oscopy`. Option
3 handles this as `_endoscopy_`. This is an identifier that starts _and
finishes_ with an underbar. The underbars act as delimiting quotes,
escape-syntax (`\`) is enabled but whitespace must be quoted. 

N.B. Starting with an underbar and not finishing on an underbar but using escape
syntax is a token-error e.g. `_\s\n`.

This supports:

  - Form starts and ends as identifiers e.g. `_if_` and `_endif_`.
  - Arbitrary names as identifiers e.g. `_\#123_` and `_\s_`.
  - Discard i.e. `_`
  - Ordinary variables that simply start with an underbar e.g. `_labrador`
  - Raw string syntax with backslash e.g. \'This is a raw string'. Note the syntax combines with multi-line syntax.


