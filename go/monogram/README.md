# Monogram

Monogram is a Go-based tool and library that parses and transforms Monogram
files (e.g. `example.mg`) into various output formats such as XML, JSON, Mermaid
diagrams, and more. Its architecture is split into two main parts:

- **CLI:** A command-line interface for quick, stand-alone use.
- **Library:** A set of reusable functions and types that can be embedded into
  other projects.

## Project Structure

The source of this tool is kept in the go/monogram subfolder of the repository
and is organized as follows:

```
go/monogram/ 
├── cmd
│   ├── calc
│   │   └── main.go # Example application using lib monogram
│   ├── monogram
│   │   └── main.go # Entry point for the CLI executable (package main)
│   └── typecalc
│       └── main.go # Duplicate application using strongly typed AST
├── lib/ 
│   ├── ast.go 
│   ├── dot_writer.go 
│   ├── json_writer.go 
│   ├── mermaid_writer.go 
│   ├── token.go 
│   ├── tokeniser.go 
│   ├── vscodeClassifyTokens.go
│   ├── xml_writer.go 
│   └── yaml_writer.go # Core library code (package lib) 
├── go.mod
├── go.sum 
└── Justfile # Optional build/run scripts
```

