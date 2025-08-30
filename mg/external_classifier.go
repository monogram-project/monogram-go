package mg

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// ExternalClassifier handles communication with an external token classification program
type ExternalClassifier struct {
	cmd             *exec.Cmd
	stdin           io.WriteCloser
	stdout          io.ReadCloser
	reader          *bufio.Reader
	tokenTexts      []string                       // Collect token texts for batch processing
	classifications []*ExternalTokenClassification // Store results from batch processing
	processed       bool                           // Flag to track if batch processing is complete
}

// ExternalTokenClassification represents the classification result from the external classifier
type ExternalTokenClassification struct {
	Role            string   // V, P, S, E, O, L, C, X
	EndTokens       []string // For S (form-start) tokens, possible end tokens
	PrefixPrec      uint16   // For O (operator) tokens, prefix precedence (0 = not applicable)
	InfixPrec       uint16   // For O (operator) tokens, infix precedence (0 = not applicable)
	PostfixPrec     uint16   // For O (operator) tokens, postfix precedence (0 = not applicable)
	ExceptionReason string   // For X (exception) tokens, the reason
}

// NewExternalClassifier creates a new external classifier using the given command
func NewExternalClassifier(command string) (*ExternalClassifier, error) {
	if command == "" {
		return nil, fmt.Errorf("external classifier command cannot be empty")
	}

	// Split command into parts (simple split by spaces - could be enhanced for quoted args)
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid external classifier command: %s", command)
	}

	cmd := exec.Command(parts[0], parts[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		stdin.Close()
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		stdin.Close()
		stdout.Close()
		return nil, fmt.Errorf("failed to start external classifier command: %w", err)
	}

	return &ExternalClassifier{
		cmd:             cmd,
		stdin:           stdin,
		stdout:          stdout,
		reader:          bufio.NewReader(stdout),
		tokenTexts:      make([]string, 0),
		classifications: make([]*ExternalTokenClassification, 0),
		processed:       false,
	}, nil
}

// AddToken adds a token to the batch for processing
func (ec *ExternalClassifier) AddToken(token string) {
	ec.tokenTexts = append(ec.tokenTexts, token)
}

// ProcessBatch sends all collected tokens to the external classifier and reads all responses
func (ec *ExternalClassifier) ProcessBatch() error {
	if ec.processed {
		return fmt.Errorf("batch has already been processed")
	}

	// Create a channel to collect responses and an error channel
	responseChan := make(chan string, len(ec.tokenTexts))
	errorChan := make(chan error, 1)

	// Start a goroutine to continuously drain the output
	go func() {
		defer close(responseChan)

		for i := 0; i < len(ec.tokenTexts); i++ {
			responseBytes, _, err := ec.reader.ReadLine()
			if err != nil {
				select {
				case errorChan <- fmt.Errorf("failed to read response %d from external classifier: %w", i+1, err):
				default:
				}
				return
			}
			select {
			case responseChan <- string(responseBytes):
			default:
				return
			}
		}
	}()

	// Send all tokens to the external classifier
	for _, token := range ec.tokenTexts {
		if _, err := fmt.Fprintln(ec.stdin, token); err != nil {
			return fmt.Errorf("failed to send token to external classifier: %w", err)
		}
	}

	// Close stdin to signal we're done sending input
	if err := ec.stdin.Close(); err != nil {
		return fmt.Errorf("failed to close stdin: %w", err)
	}
	ec.stdin = nil

	// Collect all responses
	for i := 0; i < len(ec.tokenTexts); i++ {
		select {
		case response, ok := <-responseChan:
			if !ok {
				// Channel closed, check for error
				select {
				case err := <-errorChan:
					return err
				default:
					return fmt.Errorf("response channel closed unexpectedly at response %d", i+1)
				}
			}

			classification, err := parseClassificationResponse(response)
			if err != nil {
				return fmt.Errorf("failed to parse response %d from external classifier: %w", i+1, err)
			}

			ec.classifications = append(ec.classifications, classification)

		case err := <-errorChan:
			return err
		}
	}

	ec.processed = true
	return nil
}

// GetClassification returns the classification for the token at the given index
func (ec *ExternalClassifier) GetClassification(index int) (*ExternalTokenClassification, error) {
	if !ec.processed {
		return nil, fmt.Errorf("batch has not been processed yet")
	}
	if index < 0 || index >= len(ec.classifications) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	return ec.classifications[index], nil
}

// parseClassificationResponse parses the response from the external classifier
func parseClassificationResponse(response string) (*ExternalTokenClassification, error) {
	fields := strings.Fields(response)
	if len(fields) == 0 {
		return nil, fmt.Errorf("empty response from external classifier")
	}

	classification := &ExternalTokenClassification{
		Role: fields[0],
	}

	switch fields[0] {
	case "V": // Variable - no additional fields
		if len(fields) != 1 {
			return nil, fmt.Errorf("variable classification should have no additional fields, got: %s", response)
		}

	case "P": // Prefix form - no additional fields
		if len(fields) != 1 {
			return nil, fmt.Errorf("prefix form classification should have no additional fields, got: %s", response)
		}

	case "S": // Form-start - followed by possible end tokens
		if len(fields) < 2 {
			return nil, fmt.Errorf("form-start classification must specify end tokens, got: %s", response)
		}
		classification.EndTokens = fields[1:]

	case "E": // Form-end - no additional fields
		if len(fields) != 1 {
			return nil, fmt.Errorf("form-end classification should have no additional fields, got: %s", response)
		}

	case "O": // Operator - followed by 3 precedence values
		if len(fields) != 4 {
			return nil, fmt.Errorf("operator classification must have 3 precedence values, got: %s", response)
		}

		prefix, err := strconv.ParseUint(fields[1], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("invalid prefix precedence: %s", fields[1])
		}
		classification.PrefixPrec = uint16(prefix)

		infix, err := strconv.ParseUint(fields[2], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("invalid infix precedence: %s", fields[2])
		}
		classification.InfixPrec = uint16(infix)

		postfix, err := strconv.ParseUint(fields[3], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("invalid postfix precedence: %s", fields[3])
		}
		classification.PostfixPrec = uint16(postfix)

	case "L": // Simple label - no additional fields
		if len(fields) != 1 {
			return nil, fmt.Errorf("simple label classification should have no additional fields, got: %s", response)
		}

	case "C": // Compound label - no additional fields
		if len(fields) != 1 {
			return nil, fmt.Errorf("compound label classification should have no additional fields, got: %s", response)
		}

	case "X": // Exception - followed by reason
		if len(fields) < 2 {
			return nil, fmt.Errorf("exception classification must specify reason, got: %s", response)
		}
		classification.ExceptionReason = strings.Join(fields[1:], " ")

	default:
		return nil, fmt.Errorf("unknown classification role: %s", fields[0])
	}

	return classification, nil
}

// Close closes the external classifier
func (ec *ExternalClassifier) Close() error {
	var errs []error

	if ec.stdin != nil {
		if err := ec.stdin.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close stdin: %w", err))
		}
	}

	if ec.stdout != nil {
		if err := ec.stdout.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close stdout: %w", err))
		}
	}

	if ec.cmd != nil {
		if err := ec.cmd.Wait(); err != nil {
			errs = append(errs, fmt.Errorf("external classifier exited with error: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while closing external classifier: %v", errs)
	}

	return nil
}
