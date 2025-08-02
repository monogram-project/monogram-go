We are collaborating on a Go program, codename `monogram`. The `monogram` tool
is a command-line tool that translates from `monogram` notation into XML, JSON
and other formats. The notation is designed to represent program-like texts.
However it is just a notation and not a programming language, although it does
have an opinionated grammar. Consequently it has no built-in variables, no
built-in operators and even the reserved words are dynamically discovered during
the parse.

We have completed a good first version of that tool in this repo. The command-line
application resides in the `cmd/monogram` directory. The main logic of the
application is in the `mg/monogram` directory.

Our current focus is on a new feature, which allows us to modify the 
sub-categories assigned during tokenization based on regular expressions.

Classifying identifiers follows this logic (roughly):

- In the initial classifier phase:
    - Is this identifier `endfoo` an ending keyword and matched by an opening keyword? If so then mark it as such.
    - Is this identifier `foo` matched by an ending keyword `endfoo`? Then mark it as such.
    - Is this identifier `foo` followed by a `!`? If so then mark it as an IdentifierFormPrefix.

- While subsequently parsing:
    - If it is already classified, then leave it as is.
    - Is this identifier `else` potentially a label and followed by a `:`? If so then mark it as an IdentifierSimpleLabel.
    - If this identifier `else` is followed by `-if` then mark it as a IdentifierCompoundLabel.
    - Otherwise, leave it as an IdentifierVariable.

We want to adapt it to allow for regular expressions to be used in the
classification process. The new logic will be:

- In the initial classifier phase:
    - Does this identifier `endfoo` match the form-end-regex? If so mark it as such.
    - If there is no form-end-regex then:
        - Is this identifier `endfoo` an ending keyword and matched by an opening keyword? If so then mark it as such.
    - Does this identifier `foo` match the form-start-regex? If so mark it as such.
    - If there is no form-start-regex then:
        - Is this identifier `foo` matched by an ending keyword `endfoo`? Then mark it as such.
    - Does this identifier `foo` match the form-prefix-regex? If so mark it as such.
    - If there is no form-prefix-regex then:
        - Is this identifier `foo` followed by a `!`? If so then mark it as an IdentifierFormPrefix.
    - Does this identifier `else` match the simple-label-regex? If so mark it as such.
    - Does this identifier `elseif` match the compound-label-regex? If so mark it as such.

- While subsequently parsing:
    - If it is already classified, then leave it as is.
    - If there is no simple-label-regex then:
        - Is this identifier `else` potentially a label and followed by a `:`? If so then mark it as an IdentifierSimpleLabel.
    - If there is no compound-label-regex then:
        - If this identifier `else` is followed by `-if` then mark it as a IdentifierCompoundLabel.
    - Otherwise, leave it as an IdentifierVariable.

Our first step is to implement the new sub-category IdentifierFormPrefix.
