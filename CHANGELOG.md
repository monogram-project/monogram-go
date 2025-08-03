# Change Log for Monogram-Go

Following the style in https://keepachangelog.com/en/1.0.0/

## Unreleased 

### Added

- Configuration file support via `--config` (`-c`) flag. Supports YAML format with both regex patterns for identifier classification and default options. See [docs/config.md](docs/config.md) for details.

### Changed

- `--default-breaker` flag renamed to `--default-label` for clarity and consistency.
- Configuration file option fields now follow `option-{commandOptionName}` naming scheme (e.g., `option-format`, `option-indent`).

### Removed

- Experimental `--options-file` flag has been withdrawn in favor of the new YAML configuration file system.

## [0.7.1] XML start/end tags

### Changed

- Prefix operators now bind less tightly than `.` or `(`.

- XML-style tags have been added to the grammar. Technically this is a change
  because it prevents the use of `<` as a prefix operator. However you can now
  write the following:

```txt
<monogram>
  <entry id="1">
    <name> "Sample Entry" </name>
    <description> "This is a simple XML example for Monogram." </description>
  </entry>
</monogram>
```

Note that this feature only provides the _tags_ and not character data. This
is not an attempt to reinvent XML, merely to allow you to express XML's 
structure and nesting in an intuitive fashion.

The above example is translated into XML like this:

```xml
<unit>
  <element>
    <identifier name="monogram" />
    <attributes />
    <children separator="undefined">
      <element>
        <identifier name="entry" />
        <attributes>
          <operator name="=" syntax="infix">
            <identifier name="id" />
            <string quote="double" specifier="" value="1" />
          </operator>
        </attributes>
        <children separator="newline">
          <element>
            <identifier name="name" />
            <attributes />
            <children separator="undefined">
              <string quote="double" specifier="" value="Sample Entry" />
            </children>
          </element>
          <element>
            <identifier name="description" />
            <attributes />
            <children separator="undefined">
              <string quote="double" specifier="" value="This is a simple XML example for Monogram." />
            </children>
          </element>
        </children>
      </element>
    </children>
  </element>
</unit>
```

## [0.6.0] Expansion of strings and statement syntax

### Changed

- Raw strings are no longer marked with a leading `\`. Instead we have
  adopted `@` to indicate raw-strings (c.f. C#). For an explanation see
  [discussion on extensible literals](https://github.com/monogram-project/monogram/blob/main/docs/decisions/0020-extensible-literals/0020-extensible-literals.md). e.g. `@"To insert a tab type '\t' at the terminal."` 

- The interior of forms and bracketed (delimited) expressions now all
  support commas, semi-colons and newlines as separators.
  
### Added

- String can now use chevrons (guillemets) as alternative quotes e.g.
  `«The quick brown \(animal) jumps over the lazy dog.»` See 
  [non-ASCII characters in Monogram](https://github.com/monogram-project/monogram/blob/main/docs/decisions/0019-unicode-characters-beyond-ascii/0019-unicode-characters-beyond-ascii.md).

- In-line raw strings can now be tagged as having a special meaning in the same
  way as multi-line strings. In this case we use the space after the `@` to 
  add the tag indicating what the string represents e.g. `@date«2025-05-27»`.

- In-line normal (non-raw) strings also support tags. In this case they use
  a leading `.` e.g. `.date«\(yyyy)-\(mm)-\(dd)»`

- Regular expressions now have their own literal syntax e.g. `⫽ab.*yz⫽`. This
  can be syntax checked with the new `--check-literals` option.

### Fixed

- The add and minus signs are treated equally. Expressions such as `x+2` now
  work in the same way as `x-2`. Previous you would have needed to add
  spaces.

- JSON can now be nested inside forms. Previously it would incorrectly give
  a syntax error.

- The `^` character is now allowed in operators. It was always meant to be
  part of the base character set, analogous to the C language.


## [0.5.5] Bug Fixes

### Fixed

- #149 Empty forms now generate forms with at least one (empty) part. (e.g. `foo endfoo`)
- Trailing labels: now generate a part. (e.g. `foo x mylabel: endfoo`)
- #150 Semi-colons no longer eaten by prefix forms inside curly braces.

### Other

- `monogram-go` now split off from the original repo.


## [0.5.2] Prefix Syntax, Numeric Literals, Library now "mg.*"

### Changed

- Extended prefix-forms for **C-style syntax**
  - To accommodate this we had to remove `f{x, y}`, which is a breaking change.
  - N.B. The major version is not bumped as we have not reached our first release.

  - This allows us to accommodate loop syntax such as `while! (x) { ... }`, in
    the style of C, and `while! x { ... }` in the style of Swift/Go. You can
    even imitate cascaded conditionals:

    ```
    # ! marks a prefix form.
    if! predicate(x) {
      action(x)
    } else-if test(y) {     # else-if is recognised as a 'label'
      action(y)
    } else: {               # else: is also recognised as a 'label'
      0
    }
    ```

  - Note that prefix-forms keep reading expressions until they find one that
    finishes on a line-break. So you need to follow the Python-like convention
    of putting new-lines inside brackets.

- The Go module is renamed from `lib` to `mg` to provide a short and distinctive
  import name.

- Simple labels (e.g. `else:`) are now blocked from also being used as an
  ordinary identifier, improving the robustness of the parse in a variety of
  situations.


### Added

- **Exponential notation for numbers** now added e.g. 1.23e+8

- Non-finite numeric literals added, supporting integer and floating-point domains.
  - Includes representations for positive infinity (`∞`, `0n1`), negative infinity (`-∞`, `-0n1`), and nullity (`⦰`, `0n0`).
  - Integer and floating point non-finite values distinguished by decimal point (`-∞.0`, `-0n1.0)`
  - Under `--decimal`, outputs conform to the IEEE 754 standard with:
    - `inf` for positive infinity
    - `-inf` for negative infinity
    - `nan` for nullity.

- Support for **non-decimal literal integers**, using upper case-letters A-Z as
  digits for bases higher than 10.
  - The usual hex literals are supported e.g. 0xFF = 255
  - Binary literals e.g. 0b1100 = 12
  - Octal numbers e.g. 0o022 = 18
  - Arbitrary bases between 2-36 using `r`
    - e.g. 36rZZ = 1295
    - This syntax borrowed from Common Lisp, Prolog, Smalltalk and Pop-11, which
      is as close to a consensus as I could find.

- Support for **non-decimal floating point**, using upper case-letters A-Z
  - The prefixes `0x`, `0b`, `0o` and `0r` apply here
  - e.g. 0x1.8 = 1.5
  
- The `monogram` and `monogram-mini` tools have a **new command-line option**
  `--decimal`.
  - This causes number nodes to have a new attribute `decimal` that holds the
    numerical value converted to a decimal string.

- Multi-line strings can now **imitate markdown code-fences** by having
  a content-specifier put after the opening triple quotes, like this:

  ``````
  ```py
  def hello():
    print("hello, world")
  ```
  ``````
  This translates into XML as follows:
  ```xml
  <unit>
    <joinlines quote="backtick" specifier="py">
      <string quote="backtick" value="def hello():" />
      <string quote="backtick" value="  print(&quot;hello, world&quot;)" />
    </joinlines>
  </unit>
  ```

  As one might hope, this works nicely for all three types of quotes, not
  just backticks.


- Balanced ternary literals supported, with the `0t` prefix. Includes integers
  (`0t101T` for decimal 28), floating-point values (`0t1.T`), and exponent
  notation (`0tT11e3`).

### Fixed

- Fixed abrupt exit from local web server when parse errors were encountered.

## [0.4.3] Fixed CodeQL Warnings

### Fixed

- [Internal] CodeQL warnings were being generated for GitHub workflows without explicit
  permissions. These have been added.

## [0.4.2] Fixed Docker image

### Fixed

- The Docker image now works correctly with `--test`. Now the test server 
  permits incoming connections from the host machine. This means that you can
  run the test server in a Docker container and access it from your host machine.

## [0.4.0] Asciitree Format

### Added

- A new output format `asciitree` has been added, suitable for inclusion
  in plain text documents. For example `{ y = x + 1; z = y * y; }` turns into:

```txt
unit
└─ delimited
   │  • kind: braces
   │  • separator: semicolon
   ├─ operator
   │  │  • name: =
   │  │  • syntax: infix
   │  ├─ identifier
   │  │     • name: y
   │  └─ operator
   │     │  • name: +
   │     │  • syntax: infix
   │     ├─ identifier
   │     │     • name: x
   │     └─ number
   │           • value: 1
   └─ operator
      │  • name: =
      │  • syntax: infix
      ├─ identifier
      │     • name: z
      └─ operator
         │  • name: *
         │  • syntax: infix
         ├─ identifier
         │     • name: y
         └─ identifier
               • name: y
```

- A new `install.sh` script that can be used as a quick and dirty way to 
  get monogram installed. Good for CI/CD pipelines.


## [0.3.0] Local Test Server

### Added 

- The new `--test` option runs `monogram` in an exploratory mode, 
  opens a web browser window with a test page.
- New executable `monogram-mini` that is built without the web server 
  for a reduced executable size (12MB -> 3MB).

![Screenshot of test server page](docs/images/test-option-screenshot.png)

## [0.2.8.5] More Bug Fixes

### Internal 

- Improved process for bumping version numbers

## [0.2.8] More Bug Fixes

### Fixes

- Fixed incorrect formating and span locations in error messages (#110)
- Fixed loophole around endendXXX identifiers (#111)
- Cleaned up messages in unit-tests.
- `foo!` now marks `foo` as a form-opener. See [decision record](docs/decisions/0014-prefix-forms-introduce-reserved-words/0014-prefix-forms-introduce-reserved-words.md).

## [0.2.7] Bug Fixes

### Fixes

- Inconsistent phrasing in the [README.md](README.md) has been fixed.

- A leading backslash is supposed to introduce an identifer. However this 
  was actually not being recognised as valid syntax. This is now fixed.

```
❯ echo '\_x' | monogram -f xml 
<unit>
  <identifier name="x" />
</unit>
```

- Statements were simple sequences of expressions. This was not intentional
  and the documented grammar is now enforced. This means that the statements 
  of a form are now separated by semi-colons or line breaks.

```
❯ echo 'block a = b * b; x = y + z endblock' | monogram -f xml 
<unit>
  <form syntax="surround">
    <part keyword="block">
      <operator name="=" syntax="infix">
        <identifier name="a" />
        <operator name="*" syntax="infix">
          <identifier name="b" />
          <identifier name="b" />
        </operator>
      </operator>
      <operator name="=" syntax="infix">
        <identifier name="x" />
        <operator name="+" syntax="infix">
          <identifier name="y" />
          <identifier name="z" />
        </operator>
      </operator>
    </part>
  </form>
</unit>

```

Or equivalently

```
❯ cat << "EOF" | ./monogram -f xml
block
    a = b * b
    x = y + z
endblock
EOF
<unit>
  <form syntax="surround">
    ETC ... deleted for brevity. Identical to above.
  </form>
</unit>

```

## [0.2.6] Docker image and Underscores in Numbers

### Added

- Underscores in numbers now supported

- Prebuilt binaries on release for popular platforms:
    - Linux on x86_64 and Arm64
    - MacOS on x86_64 and Arm64
    - Windows on Intel
- Docker image sfkleach/monogram pushed on release to Docker Hub
    - Use with `docker run --rm -i sfkleach/monogram [OPTIONS] < STDIN > STDOUT`

## [0.2.4] Fixes span info

### Added 

- CHANGELOG.md

### Fixed

- Span info for "unit" nodes
- Span info for chained infix operators


## [0.2.2.0] Fixes cosing triple quotes

### Fixes

- Closing triple quotes do not have to be on their own but may be followed by other text. e,g, for infix operators.
- Test recipes working.


## [0.2.1.1] Adds string interpolation and multi-line

### Added

- New `just` recipes `unittest` and `functest` for running the types types of tests.
- String interpolation
- Multi-line strings


## [0.2.0.5] Example implementations: calc and typecalc

### Added

- New --version option, integrated with GH workflow
- Example application calc: using dynamic 'monotype' AST
- Example application typecalc: using typed AST
- More functional tests
- Working GH Codespace
- More examples added to docs
- Status 'badges' added to the README.md

### Fixes

- Subtraction handled correctly in expressions such as `3-2`


## [0.2.0] First version of reference implementation in Golang

## [0.1] Proof of concept implemented in Pop-11


