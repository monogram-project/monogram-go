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

