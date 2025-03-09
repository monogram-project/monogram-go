# 0013 - Prefix operators, 2025-03-09

## Issue

At this time we only support infix operators. However, programming languages do support a fair range of prefix operators:

- unary minus e.g. `-x`, the most important case
- unary plus e.g. `+x`
- pointer deference and reference e.g. `*x`, `&x`
- self references e.g. `^x`
- logical not `!x` - probably the second most important case, by sheer popularity (even if the design is stupid)
- bitwise negation `~x`

And less importantly but in the mix: 
- pre-increment/decrement, `++x`, `--x` 

It also might be nice to include the plus-or-minus operator
- `+/-`


## Factors

- This does open the door to unwanted expressions such as `~ - *x`, which though syntactically fine are confusable with errors of omission  i.e. instead  of writing `x++y` we erroneously write `++y`.
    - The important cases are all single letters. Maybe it makes sense to limit it to single letters?
    - One might require a naming convention, such as unary operators must begin with an underbar? e.g. `_-x`.
- Post-fix operators might be wanted later on which would lead to ambiguity in situations such as `x * * y`. Is that `(x *) * y` or `x * (* y)`?

## Options and Outcome

- Option 1: Single character signs are prefix operators.
- Option 2: Prefix operators must be flagged by a symbol
- Option 3: Using infix operators with a symbol meaning 'No value'
- Option 4: Allow signs to be used as both prefix and infix operators freely

**Outcome**: Option 4

## Consequences

- The grammar is a bit less robust against typos.
- Postfix operators cannot be added without extra work.

## Pros and Cons of Options

### Option 1: Single character signs are prefix operators.

- Pros
    - Covers all the important use-cases
    - Limits the weakening of the grammar
- Cons
    - Grammar is weakened in all the important cases!


### Option 2: Prefix operators must be flagged by a symbol

- Pros
    - Grammar remains fully robust
- Cons
    - Visually extremely bad


### Option 3: Using infix operators with a symbol meaning 'No value'

- Pros
    - Grammar is robust
- Cons
    - Looks very bad
    - Adds extra thing to learn

### Option 4: Allow signs to be used as both prefix and infix operators freely

- Pros
    - Looks very natural
    - Very simple rule for teaching
- Cons
    - Makes grammar slightly less robust against typos
    - Cannot extend same freedom to postfix operators without ambiguity arising
- Interesting
    - Prefix operators are more important than postfix operators in mainstream languages
    - I experimented with this feature before making the selection and did not 
      find any issues beyond the obvious vulnerability to omitted variables. 
      So the vulnerability is relatively modest.
