# This file consists of a series of Monogram examples. You can try 
# parsing this file with the command `monogram -f xml -i example.mg`.

# End-of-lines comments start with a `#` as you can see.

# This first example shows how familiar and program-like we can make
# monogram. This is easily recognisable as "Hello, World".

def main():
    println("Hello, {}!", "world")
enddef

# And this example demonstrates the use of arithmetic operators
# to write arithmetic expressions.

def main():
    celsius = float(input("Enter temperature in Celsius: "))
    fahrenheit = (celsius * 9/5) + 32
    print("Temperature in Fahrenheit:", fahrenheit)
enddef

# And here is a use of nested forms using visually different keywords.
# Monogram does not have any built-in keywords or operators, so this makes
# as much sense to the parser as the Python-like text!

function nfib(n) =>
    if n <= 1:
        1
    else:
        1 + nfib(n-1) + nfib(n-2)
    endif
endfunction

# By contrast, this example shows how we might use Monogram to 
# describe configuration.

Info{
  appName: "RandomApp",
  version: "1.0.0",
  settings: {
    theme: "dark",
    autoSave: true,
    autoSaveInterval: 15
  },
  logging: {
    level: "info",
    outputFile: "/var/log/monogram.log"
  },
  features: {
    enableExperimental: false,
    maxThreads: 4
  }
}

