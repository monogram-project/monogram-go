package mg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Position struct {
	Line   int `json:"line"`   // Line number (0-indexed)
	Column int `json:"column"` // Column number (0-indexed)
}

type TokenClassification struct {
	Type  string   `json:"type"`  // e.g., "keyword", "identifier", "literal"
	Value string   `json:"value"` // The actual token string
	Start Position `json:"start"` // Start position in the input
	End   Position `json:"end"`   // End position in the input
}

func VSCodeClassifyTokens(input io.Reader, output io.Writer) {
	// Read the entire input as a string
	data, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("Error: Failed to read input: %v", err)
	}
	initToken, _, terr := tokenizeInput(string(data), 0, nil)
	if terr != nil {
		jsonOutput, err := json.MarshalIndent(map[string]interface{}{
			"error": terr, // Nest the TokenizerError under an "error" field
		}, "", "  ")
		if err != nil {
			log.Fatalf("Error serializing TokenizerError to JSON: %v", err)
		}
		fmt.Println(string(jsonOutput)) // Output JSON to stdout
		return
	}

	// Parse the tokens into nodes, which will side effect the tokens in the array
	// allowing us to detect labels accurately. We can ignore any errors as we
	// are only after the side-effect.
	coreOptions := &CoreFormatOptions{
		DefaultLabel:  "_",
		IncludeSpans:  false,
		Decimal:       false,
		CheckLiterals: false,
	}
	parseTokensToNodes(initToken, false, coreOptions)

	var classifications []TokenClassification = []TokenClassification{}

	token := initToken.NextToken
	for token.Type != Capstone {
		classifications = append(classifications, TokenClassification{
			Type:  token.VSCodeTokenType(),
			Value: token.Text,
			Start: Position{Line: token.Span.StartLine, Column: token.Span.StartColumn},
			End:   Position{Line: token.Span.EndLine, Column: token.Span.EndColumn}, // TODO!
		})
		token = token.NextToken
	}

	// Serialize the classifications into JSON
	jsonOutput, err := json.MarshalIndent(classifications, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing JSON: %v", err)
	}
	fmt.Println(string(jsonOutput)) // Output JSON to stdout
}
