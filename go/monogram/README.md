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
├── cmd/ 
│   └── monogram/ 
│       └── main.go # Entry point for the CLI executable (package main) 
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

## Prerequisites

- Go 1.24 or later installed on your machine.
- A working `GOPATH` or a Go workspace with module support.

## Installing the CLI

To install the Monogram CLI, run the following command:

```sh
go install github.com/sfkleach/monogram/go/monogram/cmd/monogram@latest
```

This command will compile the CLI and install the monogram executable in your
$GOBIN directory. Make sure $GOBIN is in your system's PATH so that you can
invoke monogram directly from the terminal.

## Using the Library

If you want to integrate Monogram's functionality into your own project, you can import the library package:

```go
import "github.com/sfkleach/monogram/go/monogram/lib"
```

For example, you might call an exported function like this:

```go
package main

import (
    "os"
    "log"
    "github.com/sfkleach/monogram/go/monogram/lib"
)

func main() {
    // PrintASTXML is an exported function in lib/xml_writer.go
    if err := lib.PrintASTXML(myAST, os.Stdout); err != nil {
        log.Fatalf("Error printing AST to XML: %v", err)
    }
}
```

## Building from Source

If you'd like to build everything from source manually:

### Clone the Repository

```sh
git clone https://github.com/sfkleach/monogram.git
cd monogram/go/monogram
```

### Build the CLI

From the module root (go/monogram), run:
```sh
go build -o monogram ./cmd/monogram
```

go build -o monogram ./cmd/monogram

### Test the Library:

go test ./lib

### Usage

Once installed, run the CLI tool with:
```sh
monogram [flags] [arguments]
```

For help and command-line options, use:
```sh
monogram -h
```

## Contributing

Contributions, bug reports, and feature requests are welcome. If you’d like to
help improve Monogram, feel free to fork this repository and submit a pull
request.

