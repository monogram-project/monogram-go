# Get Started

## Prerequisites

- Go 1.23 or later installed on your machine.
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

If you'd like to build everything from source manually then follow these steps.

### Step 1: Clone the Repository

```sh
git clone https://github.com/sfkleach/monogram.git
cd monogram/go/monogram
```

### Step 2: Build the CLI

From the module root (go/monogram), run:
```sh
go build -o monogram ./cmd/monogram
```

### Step 3: Test the Library:

```
go test ./lib
```

### Step 4: Usage

Once installed, run the CLI tool with:
```sh
monogram [flags] [arguments]
```

For help and command-line options, use:
```sh
monogram -h
```