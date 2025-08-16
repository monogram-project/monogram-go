package mg

import (
	"strings"
)

type TokenType int

const (
	// Major Types
	Capstone TokenType = iota
	Literal
	Identifier
	Punctuation
	OpenBracket
	CloseBracket
	Sign
)

// Subtypes for Capstone
const (
	CapstoneStart uint8 = iota
	CapstoneEnd
)

// Subtypes for Literal
const (
	LiteralString uint8 = iota
	LiteralNumber
	LiteralInterpolatedString
	LiteralExpressionString
	LiteralMultilineString
	LiteralRegex
)

// Subtypes for Identifier
const (
	IdentifierVariable uint8 = iota
	IdentifierFormPrefix
	IdentifierFormStart
	IdentifierFormEnd
	IdentifierSimpleLabel
	IdentifierCompoundLabel
)

// Subtypes for Punctuation
const (
	PunctuationComma uint8 = iota
	PunctuationSemicolon
)

// Subtypes for Bracket
const (
	BracketParenthesis uint8 = iota
	BracketBrace
	BracketBracket
)

// Subtypes for Sign
const (
	SignLabel uint8 = iota
	SignForce
	SignDot
	SignMinus
	SignPlus
	SignLessThan
	SignLessThanSlash
	SignGreaterThan
	SignSlashGreaterThan
	SignOperator
)

type Token struct {
	Type                 TokenType // The type of token (Sign, Bracket, etc.)
	SubType              uint8     // The specific subtype of the token (if any)
	Text                 string    // The raw text of the token
	Specifier            string    // The specifier for the token (if any) for multiline strings
	Span                 Span      // The span of the token in the source code
	PrecededByNewline    bool      // New field to indicate if the token is preceded by a newline
	FollowedByWhitespace bool      // New field to indicate if the token is followed by whitespace
	EscapeSeen           bool      // New field to indicate if an escape sequence was seen
	IsMultiLine          bool      // New field to indicate if the token is a multi-line string
	QuoteRune            rune      // New field to indicate the quote rune for strings
	NextToken            *Token    // The next token in the chain

	SubTokens []*Token // Subtokens for interpolated string tokens

	// Cache for precedence
	precValue int  // Cached precedence value
	precValid bool // Indicates if the precedence has been computed
	errFlag   bool // Cached error flag for precedence validity
}

func (t *Token) SpanString() string {
	return t.Span.SpanString()
}

// Based on the closing quote character.
func (t *Token) QuoteWord() string {
	switch t.QuoteRune {
	case '"':
		return "double"
	case '\'':
		return "single"
	case '`':
		return "backtick"
	case '»':
		return "chevron"
	default:
		return "undefined"
	}
}

func (t *Token) IsSemi() bool {
	return t.Type == Punctuation && t.SubType == PunctuationSemicolon
}

func (t *Token) IsLabelToken(formStart *Token) bool {
	return t.IsSimpleLabelToken() || t.IsCompoundLabelToken(formStart) || t.IsEffectivelyCompoundLabelToken()
}

func (t *Token) IsSimpleLabelToken() bool {
	if t.Type != Identifier {
		return false
	}
	if t.SubType == IdentifierSimpleLabel {
		return true // Simple labels are always valid
	}
	if t.SubType != IdentifierVariable {
		return false
	}
	hasFollowingSignLabel := !t.FollowedByWhitespace && t.NextToken != nil && t.NextToken.Type == Sign && t.NextToken.SubType == SignLabel
	return hasFollowingSignLabel // Variable labels must be followed by a sign label
}

func (t *Token) IsEffectivelyCompoundLabelToken() bool {
	return t.Type == Identifier && t.SubType == IdentifierCompoundLabel
}

func (t *Token) IsCompoundLabelToken(formStart *Token) bool {
	if t.Type != Identifier || t.SubType != IdentifierVariable {
		return false
	}
	return t.ContinuesLikeCompoundLabelToken(formStart)
}

func (t *Token) ContinuesLikeCompoundLabelToken(formStart *Token) bool {
	if t.FollowedByWhitespace {
		return false
	}
	t1 := t.NextToken
	if t1 == nil {
		return false
	}
	if t1.Type != Sign || t1.SubType != SignMinus {
		return false
	}
	t2 := t1.NextToken
	if t2 == nil {
		return false
	}
	if t2.Type != Identifier || (t2.SubType != IdentifierFormStart && t2.SubType != IdentifierFormPrefix) {
		return false
	}
	if t2.Text != formStart.Text {
		return false
	}
	return true
}

func (t *Token) IsLabel() bool {
	return t.Type == Sign && t.SubType == SignLabel
}

func (t *Token) IsPrefixForm() bool {
	return t.Type == Identifier && t.SubType == IdentifierFormPrefix
}

func (t *Token) IsTaggedString() bool {
	return t.IsTagged(LiteralString)
}

func (t *Token) IsTaggedInterpolatedString() bool {
	return t.IsTagged(LiteralInterpolatedString)
}

func (t *Token) IsTaggedMultilineString() bool {
	return t.IsTagged(LiteralMultilineString)
}

func (t *Token) IsTagged(subtype uint8) bool {
	if t.Type != Identifier || t.FollowedByWhitespace {
		return false
	}
	t1 := t.NextToken
	if t1 == nil || (t1.Type != Literal || (t1.SubType != LiteralString && t1.SubType != subtype)) {
		return false
	}
	if t.FollowedByWhitespace {
		return false
	}
	return true
}

func (t *Token) IsStringToken() bool {
	return t.Type == Literal && (t.SubType == LiteralString || t.SubType == LiteralInterpolatedString || t.SubType == LiteralMultilineString)
}

func (t *Token) SetSeen(tokenizer *Tokenizer, seen bool) {
	t.PrecededByNewline = seen
	t.Span.EndLine = tokenizer.lineNo
	t.Span.EndColumn = tokenizer.colNo
}

const signChars = ".*/%+-<>~!&^|?:="

// These roughly correspond to the precedence of the operators in the C language.
// Note how the prefix operators sneak ahead of their infix counterparts. Blanks
// are used to encode gaps in the precedence order.
const precCharactersInfix = ".([                */%+-<>~!&^|?:="
const precCharactersPrefx = "   .*/%-+<>~!&^|?:="

func (t *Token) DelimiterName() string {
	switch t.Type {
	case OpenBracket, CloseBracket:
		switch t.SubType {
		case BracketParenthesis:
			return "parentheses"
		case BracketBrace:
			return "braces"
		case BracketBracket:
			return "brackets"
		}
	}
	return ""
}

const (
	maxPrecedence int = 999
)

func (t *Token) InfixPrecedence() (int, bool) {
	return t.Precedence(true) // Use infix precedence
}

func (t *Token) PrefixPrecedence() (int, bool) {
	return t.Precedence(false) // Use prefix precedence
}

func (t *Token) Precedence(infix bool) (int, bool) {
	// Check if precedence is already cached
	if t.precValid {
		return t.precValue, !t.errFlag // Return cached result
	}

	// Precedence is only meaningful for Signs and Brackets
	if t.Type != Sign && t.Type != OpenBracket {
		return setCacheNoValidPrecedence(t) // Cache that this token has no valid precedence
	}

	if t.Type == Sign && (t.SubType == SignLessThanSlash || t.SubType == SignSlashGreaterThan) {
		return setCacheNoValidPrecedence(t) // Cache that this token has no valid precedence
	}

	P, ok := textPrecedence(t.Text, infix)
	if !ok {
		return setCacheNoValidPrecedence(t)
	}

	// Cache the precedence result and success
	t.precValue = P
	t.precValid = true
	t.errFlag = false // Cache success (no error)

	return P, true
}

func textPrecedence(text string, infix bool) (int, bool) {
	// Get the first rune of the token's text
	runes := []rune(text)
	// We need at least one rune. And we use spaces to encode a non-matching character in the precedence strings.
	// So if the first rune is a space, we can set it as a non-operator.
	if len(runes) == 0 || runes[0] == ' ' {
		// Invalid token with empty text
		return 0, false
	}
	firstRune := runes[0]

	// Find the position of the first rune in the signs string
	precCharacters := precCharactersInfix
	if !infix {
		precCharacters = precCharactersPrefx
	}
	pos := strings.IndexRune(precCharacters, firstRune)
	if pos == -1 {
		// If the rune is not in the signs string
		return 0, false
	}

	// Calculate precedence
	P := (pos + 1) * 10
	if len(runes) > 1 && runes[0] == runes[1] {
		// If the first rune occurs twice in the token, subtract 1
		P--
	}

	return P, true
}

func setCacheNoValidPrecedence(t *Token) (int, bool) {
	t.precValue = 0
	t.precValid = true
	t.errFlag = true
	return 0, false
}

// VSCodeTokenType maps the token's type and subtype to a VSCode semantic token type.
func (t *Token) VSCodeTokenType() string {
	switch t.Type {
	case Literal:
		switch t.SubType {
		case LiteralString:
			return "string"
		case LiteralNumber:
			return "number"
		default:
			return "literal"
		}
	case Identifier:
		switch t.SubType {
		case IdentifierVariable:
			return "variable"
		case IdentifierFormStart, IdentifierFormPrefix:
			// Assuming a callable-like entity
			return "function"
		case IdentifierFormEnd:
			// End markers can be styled as keywords
			return "keyword"
		case IdentifierSimpleLabel, IdentifierCompoundLabel:
			return "operator"
		default:
			return "identifier"
		}
	case Punctuation:
		switch t.SubType {
		case PunctuationComma, PunctuationSemicolon:
			return "delimiter"
		default:
			return "punctuation"
		}
	case OpenBracket, CloseBracket:
		switch t.SubType {
		case BracketParenthesis:
			return "delimiter.parenthesis"
		case BracketBrace:
			return "delimiter.brace"
		case BracketBracket:
			return "delimiter.bracket"
		default:
			return "delimiter"
		}
	case Sign:
		switch t.SubType {
		case SignLabel:
			return "macro"
		default:
			return "operator"
		}
	default:
		return "unknown"
	}
}

func (t *Token) isCapstone() bool {
	return t.Type == Capstone
}

func (t *Token) isNotCapstone() bool {
	return t.Type != Capstone
}

func classifySign(text string) uint8 {
	// Classify the sign based on its text
	switch text {
	case ".":
		return SignDot
	case ":":
		return SignLabel
	case "-":
		return SignMinus
	case "+":
		return SignPlus
	case "!":
		return SignForce
	case "<":
		return SignLessThan
	case "</":
		return SignLessThanSlash
	case ">":
		return SignGreaterThan
	case "/>":
		return SignSlashGreaterThan
	default:
		return SignOperator // Default case for other signs
	}
}

func (t *Token) SplitGlued() {
	// Find the first occurence of >< in the token text and split down the middle.
	parts := strings.SplitN(t.Text, "><", 2)
	if len(parts) != 2 {
		return // If the token does not contain ><, stop.
	}
	textLHS := parts[0] + ">" // Keep the first part and add the closing '>'
	textRHS := "<" + parts[1] // Keep the second part and add the opening '<'

	span := t.Span // Store the original span

	t.Text = textLHS                                     // Update the original token's text to the first part
	t.SubType = classifySign(textLHS)                    // Classify the first part
	t.Span.EndColumn = t.Span.StartColumn + len(textLHS) // Adjust the end column

	t2 := &Token{
		Type:    Sign,
		SubType: classifySign(textRHS),
		Text:    textRHS,
	}

	t2.Span = span                       // Copy the span from the original token
	t2.Span.StartColumn = span.EndColumn // Adjust the start column

	// Now tie the new token into the chain.
	t2.NextToken = t.NextToken
	t.NextToken = t2
}
