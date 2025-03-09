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
- Option 4: `\_endfoo`, using backslash to make the escape sequence part of
  and identifier combined with (say) `\_` as the null sequence expander.

**Outcome**: Option 4

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
    - Cannot handle any other use-cases
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
    - Not all existing identifiers are handled as before
- Cons
    - The situation where an underbar is only at one end of an identifier
      needs to be handled differently, which complicates the understanding.
    - Under test, we found the clash with _dunder_ variables was uncomfortable.

### Option 4

- Pros 
    - Easy to use
    - Handles arbitrary identifiers
    - No competition for the syntax
    - All existing variables are handled correctly
- Cons
    - New idea of a null-escape sequence, with `\_` expanding into 
      the empty sequence.
    - Using a null-escape sequence to convertt the type of a token into an
      identifier is quite sneaky.z

### Additional Notes

The driving use-case for this decision was to permit variables such as
`endoscopy`. This would normally be seen as a form-end matching `oscopy`. Option
4 handles this as `\_endoscopy`. This uses the null escape sequence to force the
type of the token into a ordinary identifier.


This supports:

  - Form starts and ends as identifiers e.g. `\_if` and `\_endif`.
  - Arbitrary names as identifiers e.g. `\#123` and `\s`.
  - Discard i.e. `_`
  - Ordinary variables that simply start with an underbar e.g. `_labrador`

It does not compete with:
  
  - Raw string syntax with backslash e.g. \'This is a raw string'. 
  - Note this raw string syntax combines with multi-line syntax.
