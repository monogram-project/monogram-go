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
- Option 3: `_endfoo_`, `_foo+bar_`, using underbars as delimiting quote-marks 
  and allowing escape sequences.

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




