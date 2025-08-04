package mg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Tokenizer struct {
	input              string   // The input string to tokenize
	tokens             []*Token // The array of tokens generated
	lineNo             int      // Current line number
	colNo              int      // Current column number
	pos                int      // Current byte position in the input
	NewlineSeen        bool     // New field to indicate if a newline has been seen
	markStack          []int    // Stack of position markers
	lineNoStack        []int    // Array to store line numbers for each token
	lineColStack       []int    // Array to store column numbers for each token
	SimpleLabelRegex   string   // Regex pattern for simple labels
	CompoundLabelRegex string   // Regex pattern for compound labels
	FormStartRegex     string   // Regex pattern for form start keywords
	FormEndRegex       string   // Regex pattern for form end keywords
	FormPrefixRegex    string   // Regex pattern for form prefix keywords
}

// Create a new Tokenizer
func newTokenizer(input string, _ int, simpleLabelRegex string, compoundLabelRegex string, formStartRegex string, formEndRegex string, formPrefixRegex string) *Tokenizer {
	return &Tokenizer{
		input:              input,
		tokens:             []*Token{},
		lineNo:             1,
		colNo:              1,
		pos:                0,
		NewlineSeen:        false,
		SimpleLabelRegex:   simpleLabelRegex,
		CompoundLabelRegex: compoundLabelRegex,
		FormStartRegex:     formStartRegex,
		FormEndRegex:       formEndRegex,
		FormPrefixRegex:    formPrefixRegex,
	}
}

func (t *Tokenizer) StartLineCol() LineCol {
	return LineCol{t.lineNo, t.colNo}
}

func (t *Tokenizer) markPosition() {
	// Mark the current position in the input
	t.markStack = append(t.markStack, t.pos)
	t.lineNoStack = append(t.lineNoStack, t.lineNo)
	t.lineColStack = append(t.lineColStack, t.colNo)
}

// Reset the position to the last marked position
func (t *Tokenizer) resetPosition() {
	// Reset the position to the last marked position
	if len(t.markStack) == 0 {
		return
	}
	t.pos = t.markStack[len(t.markStack)-1]
	t.lineNo = t.lineNoStack[len(t.lineNoStack)-1]
	t.colNo = t.lineColStack[len(t.lineColStack)-1]
	t.markStack = t.markStack[:len(t.markStack)-1]
	t.lineNoStack = t.lineNoStack[:len(t.lineNoStack)-1]
	t.lineColStack = t.lineColStack[:len(t.lineColStack)-1]
}

func (t *Tokenizer) popMark() string {
	// Pop the last marked position and return the corresponding substring
	if len(t.markStack) == 0 {
		return ""
	}
	start := t.markStack[len(t.markStack)-1]
	t.markStack = t.markStack[:len(t.markStack)-1]
	t.lineNoStack = t.lineNoStack[:len(t.lineNoStack)-1]
	t.lineColStack = t.lineColStack[:len(t.lineColStack)-1]
	return t.input[start:t.pos]
}

// Advance the position within the input, updating line and column numbers
func (t *Tokenizer) advancePosition(r rune) {
	if r == '\n' {
		t.lineNo++
		t.colNo = 1
	} else {
		t.colNo++
	}
	t.pos += utf8.RuneLen(r) // Move the byte position forward
}

// Peek at the current rune in the input without advancing. If it is the end of
// input, return 0 for the rune and 0 for the size. Otherwise, return the rune.
func (t *Tokenizer) peek() (rune, bool) {
	if t.pos >= len(t.input) {
		return rune(0), false // End of input
	}
	r, b := utf8.DecodeRuneInString(t.input[t.pos:])
	return r, b > 0
}

func (t *Tokenizer) isWhiteSpaceNext() bool {
	// Check if the next rune is whitespace
	if t.pos >= len(t.input) {
		return false // End of input
	}
	r, _ := utf8.DecodeRuneInString(t.input[t.pos:])
	return unicode.IsSpace(r)
}

func (t *Tokenizer) peek3() (int, rune, rune, rune) {
	// Peek at the next three runes in the input and return how many we
	// successfully read, along with the runes themselves.
	if t.pos >= len(t.input) {
		return 0, rune(0), rune(0), rune(0) // End of input
	}
	r1, b1 := utf8.DecodeRuneInString(t.input[t.pos:])
	if b1 <= 0 {
		return 0, rune(0), rune(0), rune(0) // Invalid UTF-8
	}
	if t.pos >= len(t.input) {
		return 1, r1, rune(0), rune(0) // End of input
	}
	r2, b2 := utf8.DecodeRuneInString(t.input[t.pos+b1:])
	if b2 <= 0 {
		return 1, r1, rune(0), rune(0) // Invalid UTF-8
	}
	if t.pos >= len(t.input) {
		return 2, r1, r2, rune(0) // End of input
	}
	r3, b3 := utf8.DecodeRuneInString(t.input[t.pos+b1+b2:])
	if b3 <= 0 {
		return 2, r1, r2, rune(0) // Invalid UTF-8
	}
	return 3, r1, r2, r3 // Successfully read three runes
}

// hasMoreInput checks whether there is any remaining input to be processed.
// It returns true if the current position has not reached the end of the input
// string, indicating that there is more content to tokenize.
func (t *Tokenizer) hasMoreInput() bool {
	return t.pos < len(t.input)
}

func (t *Tokenizer) peekN(n int) (rune, bool) {
	currentPos := t.pos // Start at the current position
	var r rune
	var size int

	// Iterate through the input to locate the nth rune
	for range n {
		if currentPos >= len(t.input) {
			// If we reach the end of input before finding the nth rune, return false
			return 0, false
		}

		r, size = utf8.DecodeRuneInString(t.input[currentPos:])
		if r == utf8.RuneError {
			// Handle invalid UTF-8 character by returning false
			return 0, false
		}

		// Advance to the next rune
		currentPos += size
	}

	// Return the nth rune and true (indicating success)
	return r, true
}

// Consume the current rune and advance the position
func (t *Tokenizer) consume() rune {
	r, ok := t.peek()
	if !ok {
		return r
	}
	t.advancePosition(r)
	if r == '\n' {
		t.NewlineSeen = true
	} else if !unicode.IsSpace(r) {
		t.NewlineSeen = false
	}
	return r
}

func (t *Tokenizer) discard2() {
	//	TODO: Optimise this.
	t.consume() // Consume the first rune
	t.consume() // Consume the second rune
}

func (t *Tokenizer) consumeN(n int) {
	// Consume n runes from the input
	for range n {
		if t.hasMoreInput() {
			t.consume()
		} else {
			break // Stop if we reach the end of input
		}
	}
}

// Add a token to the token list
func (t *Tokenizer) addToken(tokenType TokenType, subType uint8, text string, startLine int, startCol int) *Token {
	lc := LineCol{startLine, startCol}
	token := t.newTokenLineCol(tokenType, subType, text, lc)
	t.appendToken(token)
	return token
}

// Add a token to the token list
func (t *Tokenizer) addTokenLineCol(tokenType TokenType, subType uint8, text string, start LineCol) *Token {
	span := start.Span(LineCol{})
	token := t.newToken(tokenType, subType, text, span)
	t.appendToken(token)
	return token
}

// Pop the last token from the token list
func (t *Tokenizer) popToken() *Token {
	// Pop the last token from the token list
	if len(t.tokens) == 0 {
		return nil // No tokens to pop
	}
	token := t.tokens[len(t.tokens)-1]
	t.tokens = t.tokens[:len(t.tokens)-1] // Remove the last token
	return token
}

func (t *Tokenizer) markFollowedByWhitespace(token *Token) {
	token.FollowedByWhitespace = t.isWhiteSpaceNext()
}

// Append a token to the token list
func (t *Tokenizer) appendToken(token *Token) {
	// Append the token to the token list
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) newToken(tokenType TokenType, subType uint8, text string, span Span) *Token {
	// Create a new token with the current line and column numbers
	return &Token{
		Type:    tokenType,
		SubType: subType,
		Text:    text,
		Span:    span,
	}
}

func (t *Tokenizer) newTokenLineCol(tokenType TokenType, subType uint8, text string, start LineCol) *Token {
	// Create a new token with the current line and column numbers
	return &Token{
		Type:    tokenType,
		SubType: subType,
		Text:    text,
		Span:    start.Span(LineCol{t.lineNo, t.colNo}),
	}
}

func (t *Tokenizer) tokenize() *MonogramError {
	for t.hasMoreInput() {
		r, _ := t.peek()
		seen := t.NewlineSeen

		// Skip whitespace
		if unicode.IsSpace(r) {
			t.consume()
			continue
		}

		// Skip comments
		if r == '#' {
			for t.hasMoreInput() {
				r, _ := t.peek()
				if r == '\n' || r == '\r' {
					t.tryConsumeNewline()
					break
				}
				t.consume()
			}
			continue
		}

		// Match strings
		if isOpeningQuoteChar(r) {
			_, ok := t.tryPeekTripleOpeningQuotes()
			if ok {
				token, terr := t.readMultilineString(false)
				if terr != nil {
					return terr
				}
				token.SetSeen(t, seen)
				continue
			}
			token, terr := t.readString(false, r)
			if terr != nil {
				return terr
			}
			token.SetSeen(t, seen)
			continue
		}

		// Match numbers
		if unicode.IsDigit(r) || ((r == '-' || r == '+') && t.IsNumberFollowing()) {
			token, terr := t.readNumber()
			if terr != nil {
				return terr
			}
			token.SetSeen(t, seen)
			continue
		}

		// Match non-finite number symbols.
		if r == '∞' || r == '⦰' {
			lc := t.StartLineCol()
			t.consume()
			t.addTokenLineCol(Literal, LiteralNumber, string(r), lc).SetSeen(t, seen)
			continue
		}

		if r == '-' || r == '+' {
			r1, ok := t.peekN(2)
			if ok && (r1 == '∞' || r1 == '⦰') {
				lc := t.StartLineCol()
				t.consume()
				t.consume()
				var text strings.Builder
				text.WriteRune(r)
				text.WriteRune(r1)
				t.addTokenLineCol(Literal, LiteralNumber, text.String(), lc).SetSeen(t, seen)
				continue
			}
		}

		// Match identifiers
		if unicode.IsLetter(r) || r == '_' {
			_, err := t.readIdentifierSetSeen(seen)
			if err != nil {
				return err
			}
			continue
		}

		// Match punctuation
		if r == ',' || r == ';' {
			t.readPunctuation().SetSeen(t, seen)
			continue
		}

		// Match brackets
		if r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' {
			t.readBracket().SetSeen(t, seen)
			continue
		}

		// Match signs
		if t.isSign(r) {
			t.readSign().SetSeen(t, seen)
			continue
		}

		// Match tokens starting with backslash (`\`)
		if r == '\\' {
			_, err := t.readIdentifierSetSeen(seen)
			if err != nil {
				return err
			}
			continue
		}

		// Match tokens starting with `@`.
		if r == '@' {
			// See if we have a tag.
			t.consume() // Consume the '@'
			tagText := ""
			r, ok := t.peek()
			if ok && (unicode.IsLetter(r) || r == '_') {
				tagText = t.takeTagText()
			}
			r, ok = t.peek()
			if ok && isOpeningQuoteChar(r) {
				_, is_triple := t.tryPeekTripleOpeningQuotes()
				var token *Token
				var terr *MonogramError
				if is_triple {
					token, terr = t.readMultilineString(true)
				} else {
					token, terr = t.readRawString(false, r)
				}
				if terr != nil {
					return terr
				}
				if token.Specifier != "" && tagText != "" && token.Specifier != tagText {
					return &MonogramError{
						Message: fmt.Sprintf("Tag specifier '%s' does not match existing specifier '%s'", tagText, token.Specifier),
						Line:    t.lineNo,
						Column:  t.colNo,
					}
				}
				if tagText != "" {
					token.Specifier = tagText
				}
				token.SetSeen(t, seen) // Process as a raw string
			} else {
				return &MonogramError{
					Message: fmt.Sprintf("Expected opening quote after '@%s'", tagText),
					Line:    t.lineNo,
					Column:  t.colNo,
				}
			}
			continue
		}

		// Match regex literals.
		if r == '⫽' {
			terr := t.readRegexLiteral()
			if terr != nil {
				return terr
			}
			continue
		}

		// Discard unexpected characters
		return &MonogramError{Message: fmt.Sprintf("Unexpected character: %c", r), Line: t.lineNo, Column: t.colNo}
	}
	return nil
}

func (t *Tokenizer) readRegexLiteral() *MonogramError {
	t.consume() // Consume the '⫽' character
	txt, err := t.readToRune('⫽')
	if err != nil {
		return err
	}
	token := t.addToken(Literal, LiteralRegex, txt, t.lineNo, t.colNo)
	t.markFollowedByWhitespace(token)
	return nil
}

func (t *Tokenizer) IsNumberFollowing() bool {
	r, b := t.peekN(2)
	return b && unicode.IsDigit(r)
}

func (t *Tokenizer) isSign(r rune) bool {
	return strings.ContainsRune(signChars, r)
}

func (t *Tokenizer) readToRune(target rune) (string, *MonogramError) {
	// Read characters until the target rune is found or end of line/input.
	var text strings.Builder
	for t.hasMoreInput() {
		r, _ := t.peek()
		if r == target {
			t.consume() // Consume the target rune
			return text.String(), nil
		}
		if r == '\n' || r == '\r' {
			return "", &MonogramError{Message: "Unexpected newline while reading to rune", Line: t.lineNo, Column: t.colNo}
		}
		text.WriteRune(t.consume()) // Consume the current character
	}
	return "", &MonogramError{Message: fmt.Sprintf("End of input reached before finding rune %c", target), Line: t.lineNo, Column: t.colNo}
}

func (t *Tokenizer) readSign() *Token {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	prev_was_gt := false
	for t.hasMoreInput() {
		r, _ := t.peek()
		if !t.isSign(r) {
			break
		}
		if prev_was_gt && r == '<' {
			break
		}
		prev_was_gt = (r == '>')

		t.consume()
	}

	// Add the sign token
	text := t.input[start:t.pos]
	subType := classifySign(text)
	token := t.addToken(Sign, subType, text, startLine, startCol)
	t.markFollowedByWhitespace(token) // Mark if followed by whitespace
	return token
}

func (t *Tokenizer) readBracket() *Token {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the bracket character

	// Determine the subtype
	var subType uint8
	var ttype TokenType
	switch r {
	case '(':
		ttype = OpenBracket
		subType = BracketParenthesis
	case ')':
		ttype = CloseBracket
		subType = BracketParenthesis
	case '[':
		ttype = OpenBracket
		subType = BracketBracket
	case ']':
		ttype = CloseBracket
		subType = BracketBracket
	case '{':
		ttype = OpenBracket
		subType = BracketBrace
	case '}':
		ttype = CloseBracket
		subType = BracketBrace
	}

	// Add the bracket token
	return t.addToken(ttype, subType, string(r), startLine, startCol)
}

func (t *Tokenizer) readPunctuation() *Token {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the punctuation character

	// Determine the subtype
	var subType uint8
	switch r {
	case ',':
		subType = PunctuationComma
	case ';':
		subType = PunctuationSemicolon
	}

	// Add the punctuation token
	return t.addToken(Punctuation, subType, string(r), startLine, startCol)
}

func (t *Tokenizer) tryPeekTripleOpeningQuotes() (rune, bool) {
	return t.tryPeekTripleQuotes(true)
}

func (t *Tokenizer) tryPeekTripleQuotes(is_opening bool) (rune, bool) {
	// Peek the first character to check if it’s a valid quote character
	r1, ok1 := t.peek()
	if !ok1 {
		return 0, false // End of input
	}
	if is_opening {
		if !isOpeningQuoteChar(r1) {
			return 0, false // Invalid opening quote character
		}
	} else {
		if !isClosingQuoteChar(r1) {
			return 0, false // Invalid closing quote character
		}
	}

	// Check if the next two characters match the first one
	r2, ok2 := t.peekN(2)
	r3, ok3 := t.peekN(3)
	if !(ok2 && ok3 && r2 == r1 && r3 == r1) {
		return 0, false // Not a triple quote
	}

	return r1, true // Successfully read triple quotes
}

func (t *Tokenizer) consumeTripleClosingQuotes(quote rune) *MonogramError {
	r, b := t.tryReadTripleClosingQuotes()
	if !b {
		return &MonogramError{Message: "Missing triple quotes", Line: t.lineNo, Column: t.colNo}
	}
	if r != quote {
		return &MonogramError{Message: fmt.Sprintf("Expected %c, but found %c", quote, r), Line: t.lineNo, Column: t.colNo}
	}
	return nil
}

func (t *Tokenizer) tryReadTripleClosingQuotes() (rune, bool) {
	return t.tryReadTripleQuotes(false)
}

func (t *Tokenizer) tryReadTripleOpeningQuotes() (rune, bool) {
	return t.tryReadTripleQuotes(true)
}

func (t *Tokenizer) tryReadTripleQuotes(is_opening bool) (rune, bool) {
	r, b := t.tryPeekTripleQuotes(is_opening)

	if b {
		// Consume all three quotes
		t.consume() // Consume the first quote
		t.consume() // Consume the second quote
		t.consume() // Consume the third quote
	}

	return r, b
}

func (t *Tokenizer) skipSpacesUpToNewline() {
	// Skip whitespace characters
	for t.hasMoreInput() {
		r, ok := t.peek()
		if !ok || r == '\n' || r == '\r' {
			break
		}
		if !unicode.IsSpace(r) {
			break // Stop if a non-whitespace character is found
		}
		t.consume() // Consume the whitespace character
	}
}

// Method to read the specifier of a multi-line string / code-fence.
func (t *Tokenizer) readSpecifier() (string, *MonogramError) {
	// Read all the characters until a newline or end of input.
	var text strings.Builder
	for t.hasMoreInput() {
		r := t.consume()
		if r == '\n' || r == '\r' {
			if r == '\r' {
				t.tryConsumeRune('\n') // Consume \n if it follows
			}
			break // End of line
		}
		text.WriteRune(r)
	}
	strtext := strings.TrimSpace(text.String())
	if strings.Contains(strtext, " ") {
		return "", &MonogramError{Message: "Spaces inside code-fence specifier", Line: t.lineNo, Column: t.colNo}
	}
	//  Check the specifier matches the regex ^\w*$. This reserves wriggle room
	//  for future expansion.
	if len(strtext) > 0 {
		m, e := regexp.MatchString(`^[a-zA-Z_]\w*$`, strtext)
		if !m || e != nil {
			return "", &MonogramError{Message: "Invalid code-fence specifier", Line: t.lineNo, Column: t.colNo}
		}
	}
	return strtext, nil
}

// Calculates the closing indent if we are on the last line of a multiline string.
func textIsWhitespaceFollowedBy3Quotes(text string, quote rune) (bool, string) {
	// Check if the text is whitespace followed by three quotes
	if len(text) < 3 {
		return false, ""
	}
	var indent strings.Builder
	whitespace := true
	count := 0
	for _, r := range text {
		if whitespace {
			if unicode.IsSpace(r) {
				indent.WriteRune(r)
				continue
			} else if r == quote {
				whitespace = false
				count++
			} else {
				return false, ""
			}
		} else {
			if r == quote {
				count++
				if count >= 3 {
					return true, indent.String()
				}
			} else {
				return false, ""
			}
		}
	}
	return false, ""
}

func (t *Tokenizer) findClosingIndent() (rune, string, string, int, *MonogramError) {
	t.markPosition()

	// Validate and consume the opening triple quotes
	opening_quote, ok := t.tryReadTripleOpeningQuotes()
	if !ok {
		return 0, "", "", 0, &MonogramError{Message: "Malformed opening triple quotes", Line: t.lineNo, Column: t.colNo}
	}
	closing_quote := getMatchingCloseQuote(opening_quote) // Get the matching closing quote

	// Ensure no other non-space characters appear on the opening line
	specifier, terr := t.readSpecifier()
	if terr != nil {
		return 0, "", "", 0, terr
	}

	// Now read each line in order until we find the closing line.
	startLine, startCol := t.lineNo, t.colNo
	lines := []string{}
	var match bool
	var closingIndent string
	for t.hasMoreInput() {
		line := t.readRestOfLine()
		match, closingIndent = textIsWhitespaceFollowedBy3Quotes(line, closing_quote)
		if match {
			break
		}
		lines = append(lines, line)
	}

	if !match {
		return 0, "", "", 0, &MonogramError{Message: "Closing triple quote not found", Line: t.lineNo, Column: t.colNo}
	}

	for i, line := range lines {
		// Allow empty lines
		if line == "" {
			continue
		}
		// Check if the line starts with the closing indent
		if !strings.HasPrefix(line, closingIndent) {
			return 0, "", "", 0, &MonogramError{
				Message: "not indented consistently with the closing triple quote",
				Line:    startLine + i,
				Column:  startCol,
			}
		}
	}

	t.resetPosition()
	return closing_quote, closingIndent, specifier, len(lines), nil
}

func (t *Tokenizer) readRestOfLine() string {
	// Read the rest of the line until a newline or end of input
	var text strings.Builder
	for t.hasMoreInput() {
		r, _ := t.peek()
		if r == '\n' || r == '\r' {
			break // End of line
		}
		text.WriteRune(t.consume())
	}
	t.tryConsumeNewline()
	return text.String()
}

func (t *Tokenizer) readMultilineString(rawFlag bool) (*Token, *MonogramError) {
	startLine, startCol := t.lineNo, t.colNo
	var subTokens []*Token

	openingQuote, closingIndent, specifier, nlines, terr := t.findClosingIndent()
	if terr != nil {
		return nil, terr
	}
	closingQuote := getMatchingCloseQuote(openingQuote) // Get the matching closing quote

	// Discard the rest of this line, which are the opening quotes.
	t.readRestOfLine()

	// The next N lines should be either all whitespace or start with the
	// closing indent.
	for range nlines {
		if t.tryConsumeText(closingIndent) {
			if rawFlag {
				t.readRawString(true, openingQuote)
			} else {
				t.readString(true, openingQuote)
			}
		} else {
			token := t.addToken(Literal, LiteralString, "", t.lineNo, t.colNo)
			token.QuoteRune = openingQuote
		}
		subTokens = append(subTokens, t.popToken())
	}

	// Discard the rest of the next line, which will be the closing quotes.
	t.skipSpacesUpToNewline()

	terr = t.consumeTripleClosingQuotes(closingQuote)
	if terr != nil {
		return nil, terr // Return error if closing quotes are malformed
	}

	// Add the multiline string token
	token := t.addToken(Literal, LiteralMultilineString, "", startLine, startCol)
	token.IsMultiLine = true
	token.Specifier = specifier
	token.QuoteRune = openingQuote
	token.SubTokens = subTokens

	return token, nil
}

func (t *Tokenizer) tryConsumeRune(char rune) bool {
	// Check if the next rune matches the given character
	r, ok := t.peek()
	if !ok {
		return false // End of input
	}
	if r != char {
		return false // No match
	}
	t.consume() // Consume the character
	return true
}

func (t *Tokenizer) tryConsumeNewline() bool {
	// Consume '\r' and optionally '\n' to handle both '\n' and '\r\n' line endings.
	// IMPORTANT: This direct indexing is only safe because we are testing against
	// the ASCII range. In this range, the UTF-8 encoding is identical to the ASCII.
	if t.hasMoreInput() && t.input[t.pos] == '\r' {
		t.consume() // Consume '\r'
		if t.hasMoreInput() && t.input[t.pos] == '\n' {
			t.consume() // Consume '\n' if it follows
		}
		return true
	} else if t.hasMoreInput() && t.input[t.pos] == '\n' {
		t.consume() // Consume '\n'
		return true
	}
	return false // No newline consumed
}

func (t *Tokenizer) tryConsumeText(text string) bool {
	// Check if the next characters match the given text
	if strings.HasPrefix(t.input[t.pos:], text) {
		t.consumeN(len(text)) // Consume the matching text
		return true
	}
	return false
}

func getMatchingCloseQuote(openingQuote rune) rune {
	// Return the matching closing quote for the given opening quote
	if openingQuote == '«' {
		return '»'
	}
	return openingQuote // For other quotes, return the same character
}

func (t *Tokenizer) readRawString(unquoted bool, default_quote rune) (*Token, *MonogramError) {
	startLine, startCol := t.lineNo, t.colNo
	quote := default_quote
	if !unquoted {
		quote = getMatchingCloseQuote(t.consume()) // Consume the opening quote
	}
	var text strings.Builder

	for {
		if !t.hasMoreInput() {
			return nil, &MonogramError{Message: "Unterminated raw string", Line: startLine, Column: startCol}
		}
		r := t.consume()
		if r == quote { // Closing quote found
			break
		} else if r == '\n' || r == '\r' { // Handle newlines
			if unquoted {
				if r == '\r' {
					t.tryConsumeRune('\n') // Consume '\n' if it follows
				}
				break
			}
			return nil, &MonogramError{Message: "Line break in raw string", Line: startLine, Column: startCol}
		}
		// Backslashes are treated as normal characters in raw strings
		text.WriteRune(r)
	}

	// Add the raw string token
	token := t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
	token.QuoteRune = quote
	return token, nil
}

func (t *Tokenizer) readString(unquoted bool, default_quote rune) (*Token, *MonogramError) {
	startLine, startCol := t.lineNo, t.colNo
	currSpan := Span{startLine, startCol, -1, -1}
	quote := default_quote
	if !unquoted {
		quote = getMatchingCloseQuote(t.consume()) // Consume the opening quote
	}
	var text strings.Builder
	var interpolationTokens []*Token

	for {
		if !t.hasMoreInput() {
			return nil, &MonogramError{Message: "Unterminated string", Line: startLine, Column: startCol}
		}
		r := t.consume()
		if !unquoted && r == quote { // Closing quote found
			break
		}
		if r == '\\' && t.hasMoreInput() { // Handle escape or interpolation
			next, _ := t.peek()
			if next == '(' || next == '[' || next == '{' {
				// End the current StringToken and handle interpolation
				if text.Len() > 0 {
					currSpan.EndLine, currSpan.EndColumn = t.lineNo, t.colNo-1 // Do not include the backslash
					current := t.newToken(Literal, LiteralString, text.String(), currSpan)
					current.QuoteRune = quote
					interpolationTokens = append(interpolationTokens, current)
					text.Reset()
				}
				interpolatedToken, err := t.readStringInterpolation()
				if err != nil {
					return nil, err
				}
				interpolationTokens = append(interpolationTokens, interpolatedToken)
				currSpan = Span{t.lineNo, t.colNo, -1, -1}
			} else {
				text.WriteString(handleEscapeSequence(t))
			}
		} else if r == '\n' || r == '\r' { // Handle newlines
			if unquoted {
				if r == '\r' {
					t.tryConsumeRune('\n') // Consume '\n' if it follows
				}
				break
			}
			return nil, &MonogramError{Message: "Line break in string", Line: startLine, Column: startCol}
		} else {
			text.WriteRune(r)
		}
	}

	// Add the final StringToken if there's remaining text
	if text.Len() > 0 {
		currSpan.EndLine, currSpan.EndColumn = t.lineNo, t.colNo
		token := t.newToken(Literal, LiteralString, text.String(), currSpan)
		token.QuoteRune = quote
		interpolationTokens = append(interpolationTokens, token)
	}

	// Is this just a literal string?
	if len(interpolationTokens) == 1 && interpolationTokens[0].SubType == LiteralString {
		t.appendToken(interpolationTokens[0])
		return interpolationTokens[0], nil
	}

	// Combine into a StringInterpolationToken if interpolation occurred
	compoundToken := t.addToken(Literal, LiteralInterpolatedString, "", startLine, startCol)
	compoundToken.QuoteRune = quote
	compoundToken.SubTokens = interpolationTokens
	return compoundToken, nil
}

func (t *Tokenizer) readStringInterpolation() (*Token, *MonogramError) {
	span := Span{t.lineNo, t.colNo, -1, -1}
	state := 0       // State 0: inside expression, State 1: inside string
	var stack []rune // Pushdown stack

	t.markPosition()                   // Mark the position for the interpolation
	openingRune := t.consume()         // Consume the opening bracket
	stack = append(stack, openingRune) // Push opening bracket onto stack

	for {
		if !t.hasMoreInput() {
			return nil, &MonogramError{Message: "Unterminated interpolation", Line: span.StartLine, Column: span.StartColumn}
		}
		r := t.consume()
		switch state {
		case 0: // Inside expression
			switch r {
			case '\\': // Escape sequence
				handleEscapeSequence(t)
			case '(', '[', '{': // Opening brackets
				stack = append(stack, r)
			case ')', ']', '}': // Closing brackets
				if len(stack) > 0 && matches(stack[len(stack)-1], r) {
					stack = stack[:len(stack)-1] // Pop stack
					if len(stack) == 0 {         // End of interpolation
						text := t.popMark() // Pop the marked position
						span.EndLine, span.EndColumn = t.lineNo, t.colNo
						token := t.newToken(Literal, LiteralExpressionString, text, span)
						return token, nil
					}
				} else {
					return nil, &MonogramError{Message: "Mismatched bracket", Line: span.StartLine, Column: span.StartColumn}
				}
			case '"', '\'', '`', '«': // Enter string state
				stack = append(stack, getMatchingCloseQuote(r))
				state = 1
			case 'r', '\n': // Line breaks are not allowed
				return nil, &MonogramError{Message: "Line break in interpolation", Line: t.lineNo, Column: t.colNo}
			}
		case 1: // Inside string
			switch r {
			case '\\': // Escape sequence
				if t.hasMoreInput() {
					next, _ := t.peek()
					if next == '(' || next == '[' || next == '{' {
						_, err := t.readStringInterpolation()
						if err != nil {
							return nil, err
						}
					} else {
						handleEscapeSequence(t)
					}
				} else {
					return nil, &MonogramError{Message: "Unterminated escape sequence", Line: span.StartLine, Column: span.StartColumn}
				}
			case stack[len(stack)-1]: // Matching closing quote
				stack = stack[:len(stack)-1] // Pop stack
				state = 0
			}
		}
	}
}

// Helper to check if brackets match
func matches(open, close rune) bool {
	return (open == '(' && close == ')') || (open == '[' && close == ']') || (open == '{' && close == '}')
}

// Helper method to process escape sequences
func handleEscapeSequence(t *Tokenizer) string {
	var text strings.Builder
	r := t.consume() // Consume the escape character
	switch r {
	case 'b':
		text.WriteRune('\b')
	case 'f':
		text.WriteRune('\f')
	case 'n':
		text.WriteRune('\n')
	case 'r':
		text.WriteRune('\r')
	case 't':
		text.WriteRune('\t')
	case '\\', '/', '"', '\'', '`', '»': // Escaped backslash, slash, or matching quote
		text.WriteRune(r)
	case 'u': // Unicode escape sequence
		text.WriteString(t.readUnicodeEscape())
	case '_': // Non-standard escape sequence: \_
		// Expand into no characters (do nothing)
		// This has a couple of use-cases. 1. It helps break up a dense sequence
		// of characters, making it easier to read. 2. It can be used to introduce
		// a non-standard identifier.
	default:
		text.WriteRune('\\') // Keep invalid escape sequences as-is
		text.WriteRune(r)
	}

	return text.String()
}

func (t *Tokenizer) readUnicodeEscape() string {
	var code strings.Builder
	for range 4 {
		if t.hasMoreInput() {
			r, size := utf8.DecodeRuneInString(t.input[t.pos:])
			if r == utf8.RuneError {
				break // Handle invalid UTF-8
			}
			code.WriteRune(r)
			t.pos += size // Advance by the size of the rune
		} else {
			break // Stop if there are fewer than 4 runes remaining
		}
	}
	if code.Len() == 4 {
		if decoded, err := decodeUnicodeEscape(code.String()); err == nil {
			return string(decoded)
		}
	}
	return "\\u" + code.String()
}

// Decode a Unicode escape sequence (\uXXXX) into a rune
func decodeUnicodeEscape(code string) (rune, error) {
	if r, err := strconv.ParseInt(code, 16, 32); err == nil {
		return rune(r), nil
	} else {
		return 0, err
	}
}

const hexRune = 'x'
const binRune = 'b'
const octRune = 'o'
const nonFiniteRune = 'n'
const balancedTernaryRune = 't'
const radixRune = 'r'

const hexPrefix = "0" + string(hexRune)
const binPrefix = "0" + string(binRune)
const octPrefix = "0" + string(octRune)
const nonFinitePrefix = "0" + string(nonFiniteRune)
const balancedTernaryPrefix = "0" + string(balancedTernaryRune)

type NumericCategory int

const (
	NumericBase NumericCategory = iota
	NumericBalancedTernary
	NumericNonFinite
)

func (t *Tokenizer) readBase(startLine int, startCol int) (NumericCategory, int, *MonogramError) {
	var base int = 10
	var category NumericCategory = NumericBase
	n, r1, r2, r3 := t.peek3()
	if n >= 2 {
		if r1 == '0' {
			switch r2 {
			case hexRune:
				base = 16
				t.discard2() // Consume the '0x' prefix.
			case binRune:
				base = 2
				t.discard2() // Consume the '0b' prefix.
			case octRune:
				base = 8
				t.discard2() // Consume the '0o' prefix.
			case nonFiniteRune:
				base = 2
				category = NumericNonFinite
				t.discard2() // Consume the '0n' prefix.
			case balancedTernaryRune:
				base = 3
				category = NumericBalancedTernary
				t.discard2() // Consume the '0t' prefix.
			}
		} else if unicode.IsDigit(r1) {
			if r2 == radixRune {
				// One digit radix.
				base = int(r1 - '0')
				if base <= 1 || base > 9 {
					return category, 0, &MonogramError{Message: "Invalid number format", Line: startLine, Column: startCol}
				}
				t.discard2() // Consume the '\dr'
			} else if n >= 3 && r3 == radixRune && unicode.IsDigit(r2) {
				base = int(r1-'0')*10 + int(r2-'0')
				if base <= 1 || base > 36 {
					return category, 0, &MonogramError{Message: "Invalid number format", Line: startLine, Column: startCol}
				}
				t.discard2() // Consume the '\d\d'
				t.consume()  // Consume the 'r'
			}
		}
	}
	return category, base, nil
}

func XDigitValue(r rune, category NumericCategory, base int) (int, error) {
	if category == NumericBalancedTernary {
		switch r {
		case '0':
			return 0, nil
		case '1':
			return 1, nil
		case 'T':
			return -1, nil
		}
		return 0, fmt.Errorf("invalid character for balanced ternary: %c", r)
	}
	d0 := int(r - '0')
	if 0 <= d0 && d0 <= 9 {
		if d0 < base {
			return d0, nil
		}
		return 0, fmt.Errorf("invalid character for base %d: %c", base, r)
	}
	dA := int(r-'A') + 10
	if 10 <= dA && dA <= 35 {
		if dA < base {
			return dA, nil
		}
		return 0, fmt.Errorf("invalid character for base %d: %c", base, r)
	}
	return 0, fmt.Errorf("invalid character for base %d: %c", base, r)
}

func IsXDigit(r rune, category NumericCategory, base int) bool {
	_, err := XDigitValue(r, category, base)
	return err == nil
}

func (t *Tokenizer) readNumber() (*Token, *MonogramError) {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	// Initialize prev to zero.
	var prev rune = 0

	// Handle an optional leading '-' sign.
	if t.tryConsumeRune('-') {
		prev = '-' // Assign immediately since this affects later parsing.
	} else if t.tryConsumeRune('+') {
		prev = '+' // Assign immediately since this affects later parsing.
	}

	category, base, terr := t.readBase(startLine, startCol)
	if terr != nil {
		return nil, terr
	}

	if category == NumericNonFinite {
		if t.tryConsumeRune('1') || t.tryConsumeRune('0') {
			if t.tryConsumeRune('.') {
				if t.consume() != '0' {
					return nil, &MonogramError{Message: "Invalid number format", Line: startLine, Column: startCol}
				}
			}
			r, b := t.peek()
			if b && unicode.IsDigit(r) {
				return nil, &MonogramError{Message: "Invalid non-finite number", Line: startLine, Column: startCol}
			}
			text := t.input[start:t.pos]
			token := t.addToken(Literal, LiteralNumber, text, startLine, startCol)
			return token, nil
		}
		return nil, &MonogramError{Message: "Invalid non-finite number", Line: startLine, Column: startCol}
	}

	hasDot := false

	for t.hasMoreInput() {
		r, _ := t.peek()

		if r == '.' {
			// Reject if multiple dots or if the previous character was an underscore.
			if hasDot || prev == '_' {
				break
			}
			hasDot = true
			t.consume()
		} else if IsXDigit(r, category, base) {
			t.consume()
		} else if r == '_' {
			// Allow underscores only if the previous character was a digit.
			if !IsXDigit(prev, category, base) {
				// Invalid underscore placement
				break
			}
			// Use peekIf to verify that the following character is a digit.
			r2, b := t.peekN(2)
			if !b || !IsXDigit(r2, category, base) {
				// Invalid underscore placement
				break
			}
			t.consume() // Consume the underscore
		} else {
			break
		}

		// Update prev at the end of each iteration.
		prev = r
	}
	if prev == '.' {
		return nil, &MonogramError{Message: "Floating point not followed by valid digit", Line: startLine, Column: startCol}
	}

	// Now for any exponential notation.
	if t.tryConsumeRune('e') || t.tryConsumeRune('E') {

		_ = t.tryConsumeRune('+') || t.tryConsumeRune('-')

		if r, ok := t.peek(); ok && unicode.IsDigit(r) {
			prev = t.consume()
			for t.hasMoreInput() {
				r, _ := t.peek()
				if unicode.IsDigit(r) {
					t.consume() // Consume the digit
				} else if r == '_' {
					// Allow underscores only if the previous character was a digit.
					if !unicode.IsDigit(prev) {
						// Invalid underscore placement
						break
					}
					// Use peekIf to verify that the following character is a digit.
					r2, b := t.peekN(2)
					if !b || !unicode.IsDigit(r2) {
						// Invalid underscore placement
						break
					}
					t.consume() // Consume the underscore
				} else {
					break
				}
				prev = r // Update prev at the end of each iteration.
			}
		} else {
			// If no digits follow 'e' or 'E', it's an error.
			return nil, &MonogramError{Message: "Invalid number format", Line: startLine, Column: startCol}
		}
	}

	// If no runes were consumed or the only rune consumed was a sign, return an error.
	if start == t.pos || (start == t.pos-1 && prev == '-') {
		return nil, &MonogramError{Message: "Invalid number format", Line: startLine, Column: startCol}
	}

	r, b := t.peek()
	if b && unicode.IsDigit(r) {
		return nil, &MonogramError{Message: "Invalid number with extra trailing digits", Line: startLine, Column: startCol}
	}

	text := t.input[start:t.pos]
	token := t.addToken(Literal, LiteralNumber, text, startLine, startCol)

	return token, nil
}

func (t *Tokenizer) readIdentifierSetSeen(seen bool) (*Token, *MonogramError) {
	tok, err := t.readIdentifier()
	if err != nil {
		return tok, err
	}
	tok.SetSeen(t, seen)
	return tok, nil
}

func (t *Tokenizer) readIdentifier() (*Token, *MonogramError) {
	startLineCol := t.StartLineCol()
	var text strings.Builder
	var escSeen bool = false

	for t.hasMoreInput() {
		r, _ := t.peek()

		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			// Handle escape sequences when a backslash is encountered
			if r == '\\' {
				t.consume() // Consume the backslash
				escapeText := handleEscapeSequence(t)
				text.WriteString(escapeText)
				escSeen = true
				continue
			}
			break
		}

		t.consume() // Consume the current character
		text.WriteRune(r)
	}

	// Add the identifier token with the new field
	token := t.addTokenLineCol(Identifier, IdentifierVariable, text.String(), startLineCol)
	token.EscapeSeen = escSeen
	t.markFollowedByWhitespace(token)
	return token, nil
}

func (t *Tokenizer) takeTagText() string {
	var text strings.Builder

	for t.hasMoreInput() {
		r, _ := t.peek()
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			break // Stop if the character is not part of an identifier
		}
		t.consume() // Consume the character
		text.WriteRune(r)
	}

	return text.String()
}

func isOpeningQuoteChar(r rune) bool {
	return r == '\'' || r == '"' || r == '`' || r == '«'
}

func isClosingQuoteChar(r rune) bool {
	return r == '\'' || r == '"' || r == '`' || r == '»'
}

func (t *Tokenizer) markReservedTokens() *MonogramError {
	ident_exists := make(map[string]bool)
	is_reserved := make(map[string]bool) // A subset of ident_exists

	// Collect all identifiers.
	for _, token := range t.tokens {
		if token.Type == Identifier {
			ident_exists[token.Text] = true
		}
	}

	// Mark identifiers that are used as prefixes.
	for n, token := range t.tokens {
		if token.Type != Identifier {
			continue
		}

		// Classify as a IdentifierLabel if simple-label-regex is specified.
		if t.SimpleLabelRegex != "" {
			if matched, err := regexp.MatchString(t.SimpleLabelRegex, token.Text); err == nil && matched {
				token.SubType = IdentifierSimpleLabel
				continue // Skip further processing for this token
			}
		}

		var next *Token
		if n < len(t.tokens)-1 {
			next = t.tokens[n+1]
		}
		if next != nil && next.Type == Sign && next.SubType == SignForce && !token.FollowedByWhitespace {
			if strings.HasPrefix(token.Text, "end") {
				//return fmt.Errorf("cannot use %s as an opening keyword", token.Text)
				return &MonogramError{
					Message: fmt.Sprintf("cannot use '%s' as an opening keyword", token.Text),
					Line:    token.Span.StartLine,
					Column:  token.Span.StartColumn,
				}
			}
			token.SubType = IdentifierFormPrefix
			is_reserved[token.Text] = true
		}
	}

	// Mark surround-form identifiers.
	for _, token := range t.tokens {
		if token.Type != Identifier || strings.HasPrefix(token.Text, "endend") {
			continue
		}
		if token.SubType != IdentifierVariable {
			continue // Already classified
		}
		if strings.HasPrefix(token.Text, "end") {
			stem := token.Text[3:]
			if ident_exists[stem] {
				token.SubType = IdentifierFormEnd
			}
		} else if is_reserved[token.Text] {
			token.SubType = IdentifierFormStart
		} else {
			if ident_exists["end"+token.Text] {
				token.SubType = IdentifierFormStart
			}
		}
	}
	return nil
}

// Link tokens into a singly-linked list.
func (t *Tokenizer) chainTokens() {
	for n, token := range t.tokens {
		if n == 0 {
			continue
		}
		prev := t.tokens[n-1]
		prev.NextToken = token
	}
}

// Always safe to ask for the previous token.
func (t *Tokenizer) addInitToken() *Token {
	// Create capstone token for the start of input
	startToken := t.addToken(Capstone, CapstoneStart, "<START>", t.lineNo, t.colNo)
	startToken.Span.EndLine = t.lineNo
	startToken.Span.EndColumn = t.colNo
	t.tokens = append(t.tokens, startToken) // Add to the list of tokens
	return startToken
}

// Always safe to ask for the next token.
func (t *Tokenizer) addFiniToken() *Token {
	// Create capstone token for the end of input
	endToken := t.addToken(Capstone, CapstoneEnd, "<END>", t.lineNo, t.colNo)
	endToken.Span.EndLine = t.lineNo
	endToken.Span.EndColumn = t.colNo
	endToken.NextToken = endToken         // Point to itself
	t.tokens = append(t.tokens, endToken) // Add to the list of tokens
	return endToken
}

func tokenizeInput(input string, colOffset int, simpleLabelRegex string, compoundLabelRegex string, formStartRegex string, formEndRegex string, formPrefixRegex string) (*Token, Span, *MonogramError) {
	// Create a new Tokenizer instance
	tokenizer := newTokenizer(input, colOffset, simpleLabelRegex, compoundLabelRegex, formStartRegex, formEndRegex, formPrefixRegex)

	initToken := tokenizer.addInitToken() // Add capstone token for the start of input

	// Perform tokenization
	terr := tokenizer.tokenize()
	if terr != nil {
		return nil, Span{}, terr
	}

	tokenizer.addFiniToken() // Add capstone token for the end of input

	terr = tokenizer.markReservedTokens()
	if terr != nil {
		return nil, Span{}, terr
	}

	tokenizer.chainTokens()

	if colOffset > 0 {
		for _, token := range tokenizer.tokens {
			token.Span.StartColumn += colOffset
			token.Span.EndColumn += colOffset
		}
	}

	// Discard the array of tokens.
	tokenizer.tokens = make([]*Token, 0)

	// Return the list of tokens
	return initToken, Span{1, 1, tokenizer.lineNo, tokenizer.colNo}, nil
}
