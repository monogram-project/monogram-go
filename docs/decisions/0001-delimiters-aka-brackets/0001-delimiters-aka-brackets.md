# 0001 - Delimiters (aka brackets), 2025-02-15

## Issue

There are three kinds of delimiter in common use: parentheses, brackets and
braces. Arguably there is also a fourth, namely less-than (`<`) and greater-than
(`>`). Each of these are used in slightly different ways:

For example, parentheses are commonly used in two syntactic roles:

- As simple delimiters e.g. `(a + b) * c`
- As function-applications e.g. `f(x, y)`

Likewise, brackets have two very similar roles:

- As list brackets e.g. `[p, q, r]`
- And index expressions e.g. `A[i]`

And finally braces are primarily used as simple delimiters `{ x = 1; y = 0 }`.
However in a few situations, such as Nginx configuration, we see syntax like
`events { worker_connections 1024;}`.

Should we treat the three delimiters symmetrically, so that we support both
`(E, ...)` and `F(E, ...)` for all three? Or should we strive to represent and
preseve the variants in meaning and usage?

Also, should we attempt to include angle brackets (`<E>`) in this or not?

## Factors

- Although many programming language support delimiter style, none of them
  support `F(E, ...)` style for all three.
- When a programming language supports one of the usages, the meanings are
  broadly consistent. i.e. paretheses for precedence / function application,
  brackets for array construction / indexing, braces for statement grouping.
- However Monogram is a surface syntax notation and intended to support 
  creative variations on the core theme
- And Monogram prioritises ease of learning and retention through simplicity
  and consistency.
- Angle brackets are principally used as XML tags which is a different
  pattern.

## Options and Outcome

- Option 1: Treat all three delimiters as symmetrical but exclude `<`, `>`.
- Option 2: As 1 but include `<`, `>`
- Option 3: Treat `()` and `[]` similarly but only support `{ ... }` and not `F{E, ...}`. Exclude `<`, `>`

**Outcome**: Option 1

## Pros and Cons of Options

### Option 1

- Pros
    - Simplicity and consistency support ease of learning and retention
    - Aligned with the idea of Monogram as syntax, not semantics
- Cons
    - Use of `apply` to represent `f(x, y)` is biassed towards parentheses
    - `f{x, y}`  is likely to be surprising to most programmers
    - Does not support the bracketed expression `<E>`, type-expressions
      `Map<string, string>` or XML-style tags `<td>...</td>`
- Interesting
    - Also supports `x.f(y, z)` for `x.f[y, z]` and `x.f{y, z}`

### Option 2

- Pros
    - Supports type expressions such as `List<int>`
- Cons
    - Ambiguous when combined with the usual infix usage e.g. `< x > y >`
- Interesting
    - Competes with chained comparisons e.g. `x < y > z`, although my own view
      is that chains should always be one way, so this ambiguity would not arise.
    - Does not support XML-style tags

### Option 3

- Pros
    - AST is closer aligned to likely intentions
- Cons
    - More complicated and misaligned with simplicity and consistency objectives
    - Introduces excessive intentions into a neutral notation
    - Does not support the bracketed expression `<E>`, type-expressions
      `Map<string, string>` or XML-style tags `<td>...</td>`

## Decision

I selected Option 1 because of my concern that Monogram needs to emphasise
syntactic structure in a simple and consistent way. Although some terminology
that relates to intentions (e.g. "apply", "get", "invoke") is mnemonic,
disguishing different intentions when they are not entirely aligned across
programming languages is a drawback.

I excluded other uses of `<` and `>` because their existing and higher
priority role as infix operators would cause ambiguity. I don't really know 
how best to handle these other uses, so that decision is kicked into the 
future.
