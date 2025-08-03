package mg

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// Test structure matching the YAML structure
type TestCase struct {
	Name               string `yaml:"name"`
	Command            string `yaml:"command"`
	Input              string `yaml:"input"`
	ExpectedOutput     string `yaml:"expected_output"`
	Normalize          string `yaml:"normalize,omitempty"`
	ExpectedExitStatus int    `yaml:"expected_exit_status"`
}

type TestFile struct {
	Normalize string     `yaml:"normalize"`
	Tests     []TestCase `yaml:"tests"`
}

// CheckTranslationToAST is the function to be tested
func CheckTranslationToAST(input string) error {
	p_opts := NewParserOptions()
	node, err := p_opts.ParseToAST(input, "", false)
	if err != nil {
		return err
	}
	_, err = node.ToElement()
	if err != nil {
		return err
	}
	return nil
}

// CheckTranslationToTypedAST tests the strongly-typed AST API
func CheckTranslationToTypedAST(input string) error {
	p_opts := NewParserOptions()
	element, err := p_opts.ParseToElement(input, "", false)
	if err != nil {
		return err
	}

	// Exercise the strongly-typed AST API by traversing it
	return walkElement(element)
}

// walkElement recursively walks through a strongly-typed AST element
// exercising the library API methods
func walkElement(element Element) error {
	// Test basic Element interface methods
	_ = element.Name()
	_ = element.ChildrenCount()
	_ = element.FromTo()

	// Test option access methods using iterator
	options := element.Options()
	for options.Next() {
		key := options.Value()
		_ = element.HasOption(key)
		_ = element.GetOption(key)
	}

	// Test children iteration using the iterator interface
	children := element.Children()
	for children.Next() {
		child := children.Value()
		err := walkElement(child)
		if err != nil {
			return err
		}
	}

	// Test type-specific methods by attempting type assertions
	switch node := element.(type) {
	case *NumberNode:
		_ = node.Value // Access the typed field
	case *IdentifierNode:
		_ = node.Id // Access the typed field
	case *InfixOperatorNode:
		_ = node.Op  // Access the typed field
		_ = node.LHS // Access the typed field
		_ = node.RHS // Access the typed field
	case *PrefixOperatorNode:
		_ = node.Op  // Access the typed field
		_ = node.Arg // Access the typed field
	case *StringNode:
		_ = node.Value // Access the typed field
		_ = node.Quote // Access the typed field
	case *FormNode:
		_ = node.Name // Access the typed field
	case *ApplyNode:
		_ = node.Func // Access the typed field
		_ = node.Args // Access the typed field
	}

	return nil
}

func TestAST(t *testing.T) {
	// Use glob to find all YAML files in the functests directory
	yamlFiles, err := filepath.Glob("../functests/*.yaml")
	if err != nil {
		t.Errorf("Failed to find YAML files in functests directory: %v", err)
		return
	}

	// Iterate over all YAML files
	for _, filePath := range yamlFiles {
		fileName := filepath.Base(filePath)

		t.Run(fileName, func(t *testing.T) {
			// Load the YAML file
			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Errorf("Failed to read YAML file %s: %v", fileName, err)
				return
			}

			// Parse the YAML file
			var testFile TestFile
			if err := yaml.Unmarshal(data, &testFile); err != nil {
				t.Errorf("Failed to parse YAML file %s: %v", fileName, err)
				return
			}

			// Iterate over test cases
			for _, testCase := range testFile.Tests {
				// Only run tests with expected exit status 0
				if testCase.ExpectedExitStatus == 0 {
					t.Run(testCase.Name, func(t *testing.T) {
						err := CheckTranslationToAST(testCase.Input)
						if err != nil {
							t.Errorf("Cannot parse test-case '%s' from file '%s': %v", testCase.Name, fileName, err)
						}
					})

					// Also test the strongly-typed AST
					t.Run(testCase.Name+"_typed", func(t *testing.T) {
						err := CheckTranslationToTypedAST(testCase.Input)
						if err != nil {
							t.Errorf("Cannot parse to typed AST test-case '%s' from file '%s': %v", testCase.Name, fileName, err)
						}
					})
				}
			}
		})
	}
}
