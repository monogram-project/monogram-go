[![Build](https://github.com/monogram-project/monogram-go/actions/workflows/build-and-test.yml/badge.svg?branch=main)](https://github.com/monogram-project/monogram-go/actions/workflows/build-and-test.yml)
![License](https://img.shields.io/github/license/monogram-project/monogram-go)
![Release](https://img.shields.io/github/v/release/monogram-project/monogram-go)

# Monogram

Monogram is a Go-based tool and library that parses and transforms Monogram
files (e.g. `example.mg`) into various output formats such as XML, JSON, Mermaid
diagrams, and more. Its architecture is split into two main parts:

- **CLI:** A command-line interface for quick, stand-alone use.
- **Library:** A set of reusable functions and types that can be embedded into
  other projects.

## Project Structure

The source of this tool is organized as follows:

```
❯ tree monogram-go/
monogram-go/
├── CHANGELOG.md
├── cmd
│   ├── calc
│   │   └── main.go
│   ├── monogram
│   │   ├── main.go
│   │   ├── noserver.go
│   │   └── server.go
│   └── typecalc
│       └── main.go
├── CODE_OF_CONDUCT.md
├── Dockerfile
├── docs
│   └── decisions
│       └── decision-template.md
├── example.mg
├── functests
│   ├── functest.py
│   ├── json-tests.yaml
│   ├── poetry.lock
│   ├── pyproject.toml
│   ├── xml-tests.yaml
│   └── yaml-tests.yaml
├── go
│   └── monogram
├── go.mod
├── go.sum
├── install.sh
├── jumpstart.sh
├── Justfile
├── lib
│   └── version.go
├── LICENSE.txt
├── mg
│   ├── asciitree_writer.go
│   ├── ... many go files here ...
│   └── yaml_writer.go
├── prompt.md
├── README.md
├── scripts
│   ├── bump.py
│   └── decisions.py
└── version.txt
```

