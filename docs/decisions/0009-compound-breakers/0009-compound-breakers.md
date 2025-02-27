# 0009 - Compound breakers, 2025-02-27

## Issue

We want to be able to simulate a cascaded-if. Here's an example in Python:
```py
if test:
    statements
elif test2:
    statements
else:
    statements
```

However if we only have simple breakers then we would need to write:

```ruby
if test:
    statements
elif: test2 then:
    statements
else:
    statements
```

This looks visually cluttered because of the `:` markers. We would like to have
a different style of breaker that does not require the colon marker - and also
permits the elision of the breaker's name.


## Factors

- The new type of breaker must not make undermine the robustness of Mongram.
- It should be easy-to-read
- And easy to understand 

## Options and Outcome

1. A prefix marked token e.g. `&elif`
2. A compound token that repeats the opener e.g. `else.if`, `else-if`. No
   spaces would be permitted within a compound token.
3. Any compound token is a breaker (e.g. `also-do`) but it only allows 
   breaker-name elision if the last part of the compound is the opener.

*Outcome*: Option 2

## Consequences

- We may need to move to Option 3 in the future for greater syntactic
  flexibility.

## Pros and Cons of Options

### Option 1

```py
if test:
    statements
&elif test2:
    statements
else:
    statements
```

- Pros
    - Easy to understand and use once you are familiar with it.
    - By choosing the right marker it is very robust.
- Cons
    - Unfamiliar

### Option 2

```ruby
if test:
    statements
else-if test2:
    statements
else:
    statements
```

- Pros
    - Easy to understand and use immediately.
    - Immediately familiar.

- Cons
    - Makes the grammar less robust because simple arithmetic expressions
      such as `a-b` are potentially confusable. By requiring that the last
      token is a repeat of the opener, the potential for ambiguity is very
      virtually eliminated. (An opener cannot be used as an ordinary 
      variable.)

- Interesting
    - Difficult to choose the right 'glue' symbol. I settled on `-` after
      a lot of experimentation because it was the most readable - and no other
      choice was problem free either.

## Option 3

```py
if test:
    statements
else::if test2:
    statements
else:
    statements
```
```bash
for i <- L:
    statements
break::if test do:
    statements
endfor
```

- Pros
    - Easy to understand and use after a while
    - Adds a lot of expressiveness to the grammar
    - Keeps the grammar robust

- Cons
    - Initially very unfamiliar
    - Not as readable

Interesting
    - The rules around breaker-name elision might need relaxing

