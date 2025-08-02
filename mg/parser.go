package mg

import (
	"fmt"
	"regexp"
	"strings"
)

type HowIdentsAreUsed uint8

const (
	NotUsedYet HowIdentsAreUsed = iota
	UsedAsIdentifier
	UsedAsLabel
)

// Parser holds the list of tokens and our current reading position.
type Parser struct {
	currentToken  *Token // The current token being processed.
	UnglueOption  *Token
	IncludeSpans  bool
	Decimal       bool
	CheckLiterals bool // Whether to check regex syntax.
	Idents        map[string]HowIdentsAreUsed
}

func NewParser(init_token *Token, defaultLabel string, includeSpans bool, decimal bool, checkLiterals bool) *Parser {
	return &Parser{
		currentToken:  init_token,
		UnglueOption:  &Token{Type: Identifier, SubType: IdentifierVariable, Text: defaultLabel},
		IncludeSpans:  includeSpans,
		Decimal:       decimal,
		CheckLiterals: checkLiterals,
		Idents:        make(map[string]HowIdentsAreUsed),
	}
}

// hasNext checks if there are tokens left to consume.
func (p *Parser) hasNext() bool {
	return p.currentToken.NextToken.isNotCapstone()
}

// next returns the current token and advances our pointer.
func (p *Parser) next() *Token {
	tok := p.currentToken.NextToken
	p.currentToken = tok
	if tok.isCapstone() {
		return nil // No more tokens available.
	}
	return tok
}

func (p *Parser) safeNext() *Token {
	tok := p.currentToken.NextToken
	p.currentToken = tok
	return tok
}

func (p *Parser) startLineCol() LineCol {
	t := p.currentToken.NextToken
	return LineCol{t.Span.StartLine, t.Span.StartColumn}
}

func (p *Parser) startLineColFromPrevToken() LineCol {
	t := p.currentToken
	return LineCol{t.Span.StartLine, t.Span.StartColumn}
}

func (p *Parser) endLineCol() LineCol {
	t := p.currentToken
	return LineCol{t.Span.EndLine, t.Span.EndColumn}
}

// peek returns the current token without advancing.
func (p *Parser) peek() *Token {
	t := p.currentToken.NextToken
	if t.isCapstone() {
		return nil // No more tokens available.
	}
	return t
}

func (p *Parser) safePeek() *Token {
	return p.currentToken.NextToken
}

func (p *Parser) readExpr(context Context) (*Node, error) {
	n, e := p.readExprPrec(maxPrecedence, context)
	return n, e
}

func (p *Parser) readArguments(subType uint8, context Context) (string, *Node, error) {
	lineCol := p.startLineCol()
	sep, seq, err := p.readExprSeqTo(CloseBracket, subType, true, context.setInsideForm(false))
	if err != nil {
		return "", nil, err
	}
	node := &Node{
		Name:     NameArguments,
		Children: seq,
	}
	if p.IncludeSpans {
		node.Options = map[string]string{
			OptionSpan: lineCol.SpanString(p.endLineCol()),
		}
	}
	return sep, node, nil
}

func (p *Parser) readOptExprPrec(formStart *Token, outer_prec int, context Context) (*Node, error) {
	if !p.hasNext() {
		return nil, nil
	}
	token := p.safePeek()
	if token.Type == Punctuation {
		return nil, nil
	}
	if token.Type == CloseBracket {
		return nil, nil
	}
	if token.Type == Identifier && token.SubType == IdentifierFormEnd {
		return nil, nil
	}
	if token.Type == Identifier && token.SubType == IdentifierVariable && token.IsLabelToken(formStart) {
		return nil, nil
	}
	if token.Type == Sign && token.SubType == SignLabel {
		return nil, nil
	}
	if context.AcceptNewline && token.PrecededByNewline {
		return nil, nil
	}
	return p.readExprPrec(outer_prec, context)
}

func (p *Parser) readExprPrec(outer_prec int, context Context) (*Node, error) {
	startLineCol := p.startLineCol()
	lhs, err := p.readPrimaryExpr(context)
	if err != nil {
		return nil, err
	}
	for p.hasNext() {
		token1 := p.safePeek()
		if context.AcceptNewline && token1.PrecededByNewline {
			break
		}
		if context.InsideForm && token1.Type == Sign && token1.SubType == SignLabel {
			break
		}
		if token1.Type == Literal && token1.SubType == LiteralNumber && (token1.Text[0] == '-' || token1.Text[0] == '+') {
			// Begin the classic hack to handle negative numbers.
			// TODO: Eliminate the duplication of code here.
			endLineCol := LineCol{token1.Span.StartLine, token1.Span.StartColumn + 1}
			subType := SignMinus
			text := "-"
			if token1.Text[0] == '+' {
				subType = SignPlus
				text = "+"
			}
			fake_minus_token := Token{
				Type:                 Sign,
				SubType:              subType,
				Text:                 text,
				Span:                 startLineCol.Span(endLineCol),
				PrecededByNewline:    token1.PrecededByNewline,
				FollowedByWhitespace: false,
				NextToken:            token1,
			}
			prec, ok := fake_minus_token.PrefixPrecedence()
			if !ok || prec > outer_prec {
				break
			}
			token1.Text = token1.Text[1:]
			token1.Span.StartColumn++
			token1.PrecededByNewline = false
			rhs, err := p.readExprPrec(prec, context.setInsideForm(false))
			if err != nil {
				return nil, err
			}
			curr_lhs := lhs
			lhs = &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   fake_minus_token.Text,
					OptionSyntax: ValueInfix,
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
			if p.IncludeSpans {
				curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
				if len(curr_lhs_span) >= 2 {
					lc := p.endLineCol()
					lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], lc.LineNo, lc.ColNo)
				}
			}
			continue
		}
		prec, ok := token1.InfixPrecedence()
		if !ok || prec > outer_prec {
			break
		}
		token2 := p.next()
		c := context.setInsideForm(false)
		curr_lhs := lhs
		if token2.Type == OpenBracket && token2.SubType != BracketBrace {
			sep_text, args, err := p.readArguments(token2.SubType, c)
			if err != nil {
				return nil, err
			}
			dname := token2.DelimiterName()
			lhs = &Node{
				Name: NameApply,
				Options: map[string]string{
					OptionKind:      dname,
					OptionSeparator: sep_text,
				},
				Children: []*Node{lhs, args},
			}
		} else if token2.Type == Sign && token2.SubType == SignDot && p.hasNext() {
			property := p.safeNext()
			if t := p.safePeek(); t.Type == OpenBracket && t.SubType != BracketBrace {
				token3 := p.next()
				sep_text, rhs, err := p.readArguments(token3.SubType, c)
				if err != nil {
					return nil, err
				}
				lhs = &Node{
					Name: NameInvoke,
					Options: map[string]string{
						OptionKind:      token3.DelimiterName(),
						OptionSeparator: sep_text,
						OptionName:      property.Text,
					},
					Children: []*Node{lhs, rhs},
				}
				if p.IncludeSpans {

				}
			} else {
				lhs = &Node{
					Name: NameGet,
					Options: map[string]string{
						OptionName: property.Text,
					},
					Children: []*Node{lhs},
				}
			}
		} else {
			rhs, err := p.readExprPrec(prec, c)
			if err != nil {
				return nil, err
			}
			lhs = &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   token1.Text,
					OptionSyntax: "infix",
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
		}
		if p.IncludeSpans {
			curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
			if len(curr_lhs_span) >= 2 {
				lc := p.endLineCol()
				lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], lc.LineNo, lc.ColNo)
			}
		}
	}
	if p.IncludeSpans {
		lhs.Options[OptionSpan] = startLineCol.SpanString(p.endLineCol())
	}
	return lhs, nil
}

const (
	flagComma     uint8 = 1
	flagSemicolon uint8 = 2
	flagNewline   uint8 = 4
)

func (p *Parser) readExprSeqTo(closingType TokenType, closingSubtype uint8, allowComma bool, context Context) (string, []*Node, error) {
	seq := []*Node{}
	allowFlags := flagSemicolon | flagNewline
	if allowComma {
		allowFlags |= flagComma
	}
	for p.hasNext() {
		t := p.safePeek()
		if t.Type == closingType && t.SubType == closingSubtype {
			p.safeNext()
			break
		}
		expr, err := p.readExpr(context)
		if err != nil {
			return "", nil, err
		}
		seq = append(seq, expr)
		t = p.safePeek()
		if t.Type == Punctuation {
			switch t.SubType {
			case PunctuationComma:
				if allowFlags&flagComma == 0 {
					return "", nil, fmt.Errorf("unexpected comma but found: %s", t.Text)
				}
				allowFlags = flagComma
			case PunctuationSemicolon:
				if allowFlags&flagSemicolon == 0 {
					return "", nil, fmt.Errorf("unexpected semicolon but found: %s", t.Text)
				}
				allowFlags = flagSemicolon
			}
			p.next()
			continue
		}
		if t.Type == closingType && t.SubType == closingSubtype {
			p.next()
			break
		}
		if context.AcceptNewline && t.PrecededByNewline && (allowFlags&flagNewline != 0) {
			allowFlags = flagNewline
			continue
		}
		return "", nil, fmt.Errorf("unexpected token: %s", t.Text)
	}
	sep_text := chooseKindValue(allowFlags)
	return sep_text, seq, nil
}

// chooseKindValue returns the kind value based on the allowFlags. If more
// than one flag is set, it returns ValueUndefined ("undefined").
func chooseKindValue(allowFlags uint8) string {
	switch allowFlags {
	case flagComma:
		return ValueComma
	case flagSemicolon:
		return ValueSemicolon
	case flagNewline:
		return ValueNewline
	default:
		return ValueUndefined
	}
}

func (p *Parser) tryReadExplicitTermination(allowFlags uint8) (bool, uint8, error) {
	token := p.peek()
	if token == nil {
		return false, 0, fmt.Errorf("unexpected end of tokens (missing terminator?)")
	}
	if token.Type == Punctuation {
		switch token.SubType {
		case PunctuationComma:
			if allowFlags&flagComma == 0 {
				return false, 0, fmt.Errorf("unexpected comma: %s", token.Text)
			}
			p.next()
			return true, flagComma, nil
		case PunctuationSemicolon:
			if allowFlags&flagSemicolon == 0 {
				return false, 0, fmt.Errorf("unexpected semicolon: %s", token.Text)
			}
			p.next()
			return true, flagSemicolon, nil
		}
	}
	return false, allowFlags, nil
}

func (p *Parser) requireImplicitTermination(allowFlags uint8) (uint8, error) {
	token := p.peek()
	if token == nil {
		// Token cannot be nil here, defensive programming.
		return 0, fmt.Errorf("unexpected end of tokens (internal error)")
	}
	if (allowFlags&flagNewline == 0) || !token.PrecededByNewline {
		return 0, fmt.Errorf("unexpected token while looking for separator: %s", token.Text)
	}
	return flagNewline, nil
}

type Mode uint8

const (
	alphaMode Mode = iota
	betaMode
	gammaMode
)

/*
    if ALPHA:
		BETA
		GAMMA
		GAMMA
	else-if ALPHA:
		BETA
		GAMMA
		GAMMA
	else-if ALPHA:
		BETA
		GAMMA
		GAMMA
	else:
		BETA
		GAMMA
		GAMMA
	endif
*/

func (p *Parser) readFormExpr(formStart *Token, context Context) (*Node, error) {
	allowFlags := flagComma | flagSemicolon | flagNewline
	closingTokenText := "end" + formStart.Text
	context = context.setInsideForm(true)
	mode := alphaMode
	prev_expr_explicitly_terminated := false
	startLineCol := p.startLineCol()
	var endLineCol LineCol
	builder := NewFormBuilder(formStart.Text, startLineCol, p.IncludeSpans, false)
	for {
		if !p.hasNext() {
			return nil, fmt.Errorf("unexpected end of tokens (missing end of form): %s", closingTokenText)
		}
		token := p.safePeek()
		if token.Type == Identifier && token.SubType == IdentifierFormEnd && token.Text == closingTokenText {
			endLineCol = p.endLineCol()
			p.next()
			break
		}

		if mode == alphaMode {
			n, err := p.readExpr(context)
			if err != nil {
				return nil, err
			}
			if !p.hasNext() {
				return nil, fmt.Errorf("unexpected end of input in form: %s", formStart.Text)
			}
			prev_expr_explicitly_terminated, allowFlags, err = p.tryReadExplicitTermination(allowFlags)
			if err != nil {
				return nil, err
			}
			builder.AddChild(n)
			t := p.safePeek()
			if t.IsLabel() {
				endLC := p.endLineCol()
				p.next()
				builder.BeginNextPart(p.UnglueOption.Text, endLC, p.startLineCol())
				mode = betaMode
			} else {
				mode = gammaMode
			}
		} else if token.IsSimpleLabelToken() {
			lc34 := p.endLineCol()
			t := p.next() // skip the label
			e := p.SetAsSimpleLabel(t)
			if e != nil {
				return nil, e
			}
			p.next() // remove the ':'

			builder.BeginNextPart(token.Text, lc34, p.startLineCol())
			mode = betaMode
		} else if token.IsCompoundLabelToken(formStart) {
			p.SetAsCompoundLabel(token)
			lc34 := p.endLineCol()
			t1 := p.next() // skip the label
			t2 := p.next() // remove the '-
			t3 := p.next() // remove the form-start

			builder.BeginNextPart(t1.Text+t2.Text+t3.Text, lc34, p.startLineCol())
			mode = alphaMode
		} else {
			if mode == gammaMode && !prev_expr_explicitly_terminated {
				var err error
				allowFlags, err = p.requireImplicitTermination(allowFlags)
				if err != nil {
					return nil, err
				}
			}
			n, err := p.readExpr(context)
			if err != nil {
				return nil, err
			}
			builder.AddChild(n)
			if !p.hasNext() {
				return nil, fmt.Errorf("unexpected end of input in form: %s", formStart.Text)
			}
			mode = gammaMode
			prev_expr_explicitly_terminated, allowFlags, err = p.tryReadExplicitTermination(allowFlags)
			if err != nil {
				return nil, err
			}
		}
	}
	return builder.Build(endLineCol, chooseKindValue(allowFlags)), nil
}

// readDelimitedExpr reads a delimited expression.
func (p *Parser) readDelimitedExpr(open *Token, context Context) (*Node, error) {
	sep, seq, err := p.readExprSeqTo(CloseBracket, open.SubType, true, context.setInsideDelimiters(true))
	if err != nil {
		return nil, err
	}
	dname := open.DelimiterName()
	return &Node{
		Name:     NameDelimited,
		Options:  map[string]string{OptionKind: dname, OptionSeparator: sep},
		Children: seq,
	}, nil
}

func (p *Parser) readPrimaryExpr(context Context) (*Node, error) {
	lc12 := p.startLineCol()
	n, e := p.doReadPrimaryExpr(context)
	if e != nil {
		return nil, e
	}
	if p.IncludeSpans {
		n.Options[OptionSpan] = lc12.SpanString(p.endLineCol())
	}
	return n, e
}

func (p *Parser) numberOptions(text string) (map[string]string, error) {
	if !p.Decimal {
		return map[string]string{OptionValue: text}, nil
	}
	options := map[string]string{OptionValue: text}

	// Convert the text to a high-precision floating-point number
	// and store it in the options map. We rely on the fact that the value
	// is extremely constrained at this point.
	decimal_text, err := ConvertToDecimal(text)
	if err != nil {
		return nil, err
	}
	options[OptionsDecimalValue] = decimal_text
	return options, nil
}

func (p *Parser) SetAsSimpleLabel(token *Token) error {
	token.SubType = IdentifierSimpleLabel
	if p.Idents[token.Text] == UsedAsIdentifier {
		return fmt.Errorf("identifier used as label: %s", token.Text)
	}
	p.Idents[token.Text] = UsedAsLabel
	return nil
}

func (p *Parser) SetAsCompoundLabel(token *Token) error {
	token.SubType = IdentifierCompoundLabel
	return nil
}

func (p *Parser) SetAsIdentifier(token *Token) error {
	text := token.Text
	if p.Idents[text] == UsedAsLabel {
		return fmt.Errorf("labels used as identifier: %s", text)
	}
	p.Idents[text] = UsedAsIdentifier
	return nil
}

func (p *Parser) doReadPrimaryExpr(context Context) (*Node, error) {
	if !p.hasNext() {
		return nil, fmt.Errorf("unexpected end of tokens (expression required here)")
	}
	token := p.next()

	switch token.Type {
	case Literal:
		return literalToExpr(token, p)
	case Identifier:
		if token.IsPrefixForm() {
			return p.readPrefixForm(context, token)
		}
		switch token.SubType {
		case IdentifierVariable:
			if !token.EscapeSeen {
				if e := p.SetAsIdentifier(token); e != nil {
					return nil, e
				}
			}
			return &Node{
				Name: NameIdentifier,
				Options: map[string]string{
					OptionName: token.Text,
				},
			}, nil
		case IdentifierFormPrefix:
			return p.readPrefixForm(context, token)
		case IdentifierFormStart:
			return p.readFormExpr(token, context)
		default:
			return nil, fmt.Errorf("unexpected identifier: %s", token.Text)
		}
	case OpenBracket:
		return p.readDelimitedExpr(token, context)
	case Sign:
		if token.SubType == SignDot {
			if !token.FollowedByWhitespace {
				likely_tag_token := token.NextToken
				if likely_tag_token != nil && likely_tag_token.IsStringToken() {
					// This is the degenerate case where we have a dot followed by a string.
					// We simply need to go around the loop again.
					return p.doReadPrimaryExpr(context)
				}
				var likely_string_token *Token
				if likely_tag_token != nil {
					likely_string_token = likely_tag_token.NextToken
				}
				if likely_tag_token == nil || likely_string_token == nil {
					return nil, fmt.Errorf("missing tokens after dot")
				}
				p.next() // consume the tag
				p.next() // consume the string
				if likely_tag_token.IsTaggedString() {
					return &Node{
						Name: NameString,
						Options: map[string]string{
							OptionQuote:     likely_string_token.QuoteWord(),
							OptionValue:     likely_string_token.Text,
							OptionSpecifier: likely_tag_token.Text, // Default specifier for strings
						},
					}, nil
				} else if likely_tag_token.IsTaggedInterpolatedString() {
					interpolationNode, err := p.convertInterpolatedStringSubToken(likely_string_token)
					if err != nil {
						return nil, err
					}
					interpolationNode.Options[OptionSpecifier] = likely_tag_token.Text
					return interpolationNode, nil
				} else if likely_tag_token.IsTaggedMultilineString() {
					multilineNode, err := p.convertMultilineStringSubToken(likely_string_token)
					if err != nil {
						return nil, err
					}
					existing_specifier := multilineNode.Options[OptionSpecifier]
					if existing_specifier != ValueBlank && likely_tag_token.Text != "" && existing_specifier != likely_tag_token.Text {
						return nil, fmt.Errorf("conflicting specifiers for multiline string: %s vs %s", existing_specifier, likely_tag_token.Text)
					}
					multilineNode.Options[OptionSpecifier] = likely_tag_token.Text
					return multilineNode, nil
				} else {
					return nil, fmt.Errorf("unexpected token following dot: %s", likely_string_token.Text)
				}
			} else {
				return nil, fmt.Errorf("leading-dot followed by whitespace: %s", token.Text)
			}
		} else if token.SubType == SignLessThan {
			return p.readXmlElement()
		} else if token.SubType == SignLessThanSlash {
			return nil, fmt.Errorf("unexpected closing XML tag token: %s", token.Text)
		} else if token.SubType != SignLabel {
			prec, valid := token.PrefixPrecedence()
			if valid && prec > 0 {
				expr, err := p.readExprPrec(prec, context.setInsideForm(false))
				if err != nil {
					return nil, err
				}
				return &Node{
					Name: NameOperator,
					Options: map[string]string{
						OptionName:   token.Text,
						OptionSyntax: ValuePrefix,
					},
					Children: []*Node{expr},
				}, nil
			}
		} else {
			return nil, fmt.Errorf("misplace sign token: %s", token.Text)
		}
	default:
		return nil, fmt.Errorf("unexpected token (case #1): %s, %d, %d", token.Text, token.Type, token.SubType)
	}
	return nil, fmt.Errorf("unexpected token (case #2): %s, %d, %d", token.Text, token.Type, token.SubType)
}

func literalToExpr(token *Token, p *Parser) (*Node, error) {
	switch token.SubType {
	case LiteralString:
		if token.Specifier == "re" && p.CheckLiterals && !isValidRegex(token.Text) {
			return nil, fmt.Errorf("invalid regex: %s", token.Text)
		}
		return &Node{
			Name: NameString,
			Options: map[string]string{
				OptionQuote:     token.QuoteWord(),
				OptionValue:     token.Text,
				OptionSpecifier: token.Specifier, // Default specifier for strings
			},
		}, nil

	case LiteralNumber:
		opts, err := p.numberOptions(token.Text)
		if err != nil {
			return nil, err
		}
		return &Node{
			Name:    NameNumber,
			Options: opts,
		}, nil

	case LiteralInterpolatedString: // Handling interpolated strings
		return p.convertInterpolatedStringSubToken(token)

	case LiteralMultilineString:
		return p.convertMultilineStringSubToken(token)

	case LiteralRegex:
		if p.CheckLiterals && !isValidRegex(token.Text) {
			return nil, fmt.Errorf("invalid regex: %s", token.Text)

		}
		return &Node{
			Name: NameString,
			Options: map[string]string{
				OptionQuote:     ValueRegex,
				OptionValue:     token.Text,
				OptionSpecifier: ValueRegex,
			},
		}, nil
	}
	return nil, fmt.Errorf("unexpected literal token: %s", token.Text)
}

func isValidRegex(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}

func (p *Parser) readTagExpr() (bool, *Node, error) {
	token := p.next()
	if token == nil {
		return false, nil, fmt.Errorf("unexpected end of input while reading tag expression")
	}
	if token.Type == Identifier && token.SubType == IdentifierVariable {
		return false, &Node{
			Name:    NameTag,
			Options: map[string]string{OptionName: token.Text},
		}, nil
	}
	if token.Type == OpenBracket {
		subtype := token.SubType
		expr, err := p.readExpr(Context{})
		if err != nil {
			return false, nil, fmt.Errorf("error reading tag expression: %v", err)
		}
		closingToken := p.next()
		if closingToken == nil {
			return false, nil, fmt.Errorf("unexpected end of input while reading tag expression")
		}
		if closingToken.Type != CloseBracket || closingToken.SubType != subtype {
			return false, nil, fmt.Errorf("expected closing bracket for tag expression, but got: %s", closingToken.Text)
		}
		return true, &Node{
			Name: NameDelimited,
			Options: map[string]string{
				OptionKind:      token.DelimiterName(),
				OptionSeparator: ValueUndefined, // Default separator for tags
			},
			Children: []*Node{expr},
		}, nil
	}
	return false, nil, fmt.Errorf("expected identifier or bracketed expression for tag, but got: %s", token.Text)
}

func (p *Parser) readAttrExpr() (*Node, error) {
	token := p.next()
	if token == nil {
		return nil, fmt.Errorf("unexpected end of input while reading attribute expression")
	}
	if token.Type == Literal {
		expr, err := literalToExpr(token, p)
		if err != nil {
			return nil, fmt.Errorf("error reading attribute expression: %v", err)
		}
		return expr, nil
	}
	if token.Type == OpenBracket {
		subtype := token.SubType
		expr, err := p.readExpr(Context{})
		if err != nil {
			return nil, fmt.Errorf("error reading attribute expression: %v", err)
		}
		closingToken := p.next()
		if closingToken == nil {
			return nil, fmt.Errorf("unexpected end of input while reading attribute expression")
		}
		if closingToken.Type != CloseBracket || closingToken.SubType != subtype {
			return nil, fmt.Errorf("expected closing bracket for attribute expression, but got: %s", closingToken.Text)
		}
		return &Node{
			Name: NameDelimited,
			Options: map[string]string{
				OptionKind:      token.DelimiterName(),
				OptionSeparator: ValueUndefined, // Default separator for attributes
			},
			Children: []*Node{expr},
		}, nil
	}
	return nil, fmt.Errorf("expected literal or bracketed expression for attribute, but got: %s", token.Text)
}

func (p *Parser) readXmlElement() (*Node, error) {
	// The initial `<` has been consumed at this point. Using precedence 0 to read the element
	// forces the use of brackets for any non-trivial element names.
	_, element_name, err := p.readTagExpr()
	if err != nil {
		return nil, err
	}
	attrs := &Node{Name: NameElementAttributes, Options: map[string]string{}}
	kids := &Node{Name: NameElementChildren, Options: map[string]string{}}

	element := &Node{Name: NameElement, Options: map[string]string{}}
	element.Children = append(element.Children, element_name)
	element.Children = append(element.Children, attrs)
	element.Children = append(element.Children, kids)
	kids.Options[OptionSeparator] = ValueUndefined // Default separator for kids

	// Read attributes until we hit the closing `/>` or `>`.
	attrsStartLineCol := p.startLineCol()
	var attrsEndLineCol LineCol
	for {
		attrsEndLineCol = p.endLineCol()
		token := p.peek()
		if token == nil {
			return nil, fmt.Errorf("unexpected end of input while reading XML element")
		}

		if token.Type == Sign && token.SubType == SignSlashGreaterThan {
			p.next() // consume the `/>`
			if p.IncludeSpans {
				attrsSpan := attrsStartLineCol.Span(attrsEndLineCol)
				attrs.Options[OptionSpan] = attrsSpan.SpanString()
				kidsSpan := attrsEndLineCol.Span(attrsEndLineCol)
				kids.Options[OptionSpan] = kidsSpan.SpanString()
			}
			return element, nil
		}
		if token.Type == Sign && token.SubType == SignGreaterThan {
			p.next() // consume the `>`
			break
		}

		delimited, lhs, err := p.readTagExpr()
		if err != nil {
			return nil, err
		}
		// Now for the `=` operator.
		eq_token := p.safePeek()
		if eq_token.Type == Sign && eq_token.Text == "=" {
			p.next() // consume the `=`
			rhs, err := p.readAttrExpr()
			if err != nil {
				return nil, err
			}
			attrs.Children = append(attrs.Children, &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   eq_token.Text,
					OptionSyntax: ValueInfix,
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			})
		} else if delimited {
			attrs.Children = append(attrs.Children, lhs)
		} else {
			return nil, fmt.Errorf("expected '=' after XML element attribute name, but got: %s", eq_token.Text)
		}
	}

	if p.IncludeSpans {
		attrsSpan := attrsStartLineCol.Span(attrsEndLineCol)
		attrs.Options[OptionSpan] = attrsSpan.SpanString()
	}

	// Now read elements up to the closing tag.
	kidsStartLineCol := p.startLineCol()
	sep, nodes, err := p.readExprSeqTo(Sign, SignLessThanSlash, true, makeContext())
	if err != nil {
		return nil, err
	}
	kids.Children = append(kids.Children, nodes...)
	kids.Options[OptionSeparator] = sep
	// The previous token was the closing tag, so we need to set the span from that point.
	if p.IncludeSpans {
		kids.Options[OptionSpan] = kidsStartLineCol.SpanString(p.startLineColFromPrevToken())
	}

	// Now read the closing repetition of the element name.
	closingName := p.next()
	if closingName == nil {
		return nil, fmt.Errorf("unexpected end of input while reading closing element name")
	}
	if closingName.Type != Identifier {
		return nil, fmt.Errorf("closing element name is not an identifier: %s", closingName.Text)
	}
	closingLessThen := p.next()
	if closingLessThen == nil || closingLessThen.Type != Sign || closingLessThen.SubType != SignGreaterThan {
		return nil, fmt.Errorf("expected closing '>' after closing element name: %s", closingName.Text)
	}

	// If the closing name is "_", we treat it as a wildcard and do not check it.
	if closingName.Text == "_" {
		return element, nil
	}

	if element_name.Name != NameTag {
		return nil, fmt.Errorf("expected tag name to be an identifier matching: %s (hint: use _)", closingName.Text)
	}

	if element_name.Options[OptionName] != closingName.Text {
		return nil, fmt.Errorf("closing element name '%s' does not match opening element name '%s'", closingName.Text, element_name.Options[OptionName])
	}

	return element, nil
}

func (p *Parser) readPrefixForm(context Context, token *Token) (*Node, error) {
	// Consume the expected "!" token that follows the prefix form identifier
	forceToken := p.next()
	if !(forceToken.Type == Sign && forceToken.SubType == SignForce) {
		return nil, fmt.Errorf("expected '!' after prefix form '%s', but found '%s'", token.Text, forceToken.Text)
	}
	cxt := context.setInsideForm(true)
	startAgain := true

	formBuilder := NewFormBuilder(token.Text, p.startLineCol(), p.IncludeSpans, true)

	for p.hasNext() {
		next := p.safePeek()
		if next.IsSemi() || next.PrecededByNewline {
			break
		}

		if next.IsSimpleLabelToken() {
			t := p.next()
			e := p.SetAsSimpleLabel(t)
			if e != nil {
				return nil, e
			}
			p.next() // Skip :
			formBuilder.BeginNextPart(t.Text, p.endLineCol(), p.startLineCol())
			startAgain = true
		} else if next.IsCompoundLabelToken(token) {
			p.SetAsCompoundLabel(token)
			t1 := p.next() // skip the label
			t2 := p.next() // remove the '-
			t3 := p.next() // remove the form-start
			text := t1.Text + t2.Text + t3.Text
			formBuilder.BeginNextPart(text, p.endLineCol(), p.startLineCol())
			startAgain = true
		} else if startAgain {
			n, e := p.readOptExprPrec(token, maxPrecedence, cxt)
			if e != nil {
				return nil, e
			}
			startAgain = false
			if n == nil {
				break
			}
			formBuilder.AddChild(n)
		} else {
			n, e := p.readOptExprPrec(token, maxPrecedence, cxt)
			if e != nil {
				return nil, e
			}
			if n != nil {
				formBuilder.BeginNextPart(p.UnglueOption.Text, p.endLineCol(), p.startLineCol())
				formBuilder.AddChild(n)
			} else {
				break
			}
		}
	}

	return formBuilder.Build(p.endLineCol(), ValueUndefined), nil
}

func (p *Parser) convertMultilineStringSubToken(token *Token) (*Node, error) {
	multilineNode := &Node{
		Name: NameJoinLines,
		Options: map[string]string{
			OptionQuote:     token.QuoteWord(),
			OptionSpecifier: token.Specifier,
		},
		Children: []*Node{},
	}

	// Process sub-tokens
	for _, subToken := range token.SubTokens {
		err := p.insertConvertedSubToken(subToken, multilineNode)
		if err != nil {
			return nil, err
		}
	}
	return multilineNode, nil
}

func (p *Parser) insertConvertedSubToken(subToken *Token, interpolationNode *Node) error {
	node, err := p.convertSubToken(subToken)
	if err != nil {
		return err
	}

	interpolationNode.Children = append(interpolationNode.Children, node)
	return nil
}

func (p *Parser) convertSubToken(subToken *Token) (*Node, error) {
	switch subToken.SubType {
	case LiteralExpressionString:
		// Recursively parse the expression string
		node, err := p.convertLiteralExpressionStringSubToken(subToken)
		return node, err
	case LiteralString:
		// Handle plain string parts
		n := &Node{
			Name: NameString,
			Options: map[string]string{
				OptionQuote:     subToken.QuoteWord(),
				OptionValue:     subToken.Text,
				OptionSpecifier: ValueBlank, // Default specifier for strings
			},
		}
		if p.IncludeSpans {
			n.Options[OptionSpan] = subToken.SpanString()
		}
		return n, nil
	case LiteralInterpolatedString:
		node, err := p.convertInterpolatedStringSubToken(subToken)
		return node, err
	default:
		return nil, fmt.Errorf("unexpected sub-token subtype: %v", subToken.SubType)
	}
}

func (p *Parser) convertInterpolatedStringSubToken(token *Token) (*Node, error) {
	interpolationNode := &Node{
		Name: NameJoin,
		Options: map[string]string{
			OptionQuote:     token.QuoteWord(),
			OptionSpecifier: ValueBlank, // Default specifier for interpolated strings
		},
		Children: []*Node{},
	}
	if p.IncludeSpans {
		interpolationNode.Options[OptionSpan] = token.SpanString()
	}

	// Process sub-tokens
	for _, subToken := range token.SubTokens {
		err := p.insertConvertedSubToken(subToken, interpolationNode)
		if err != nil {
			return nil, err
		}
	}
	return interpolationNode, nil
}

func (p *Parser) convertLiteralExpressionStringSubToken(subToken *Token) (*Node, error) {
	columnOffset := subToken.Span.StartColumn - 1
	p_opts := &ParserOptions{
		colOffset:    columnOffset,
		DefaultLabel: p.UnglueOption.Text,
		IncludeSpans: p.IncludeSpans,
		Decimal:      p.Decimal,
	}
	expressionNode, err := p_opts.ParseToAST(subToken.Text, "", true)
	expressionNode.Name = NameInterpolate // The outer brackets can be repurposed!
	if p.IncludeSpans {
		span := subToken.Span
		span.StartColumn-- // Now ensure we include the backslash.
		expressionNode.Options[OptionSpan] = span.SpanString()
	}
	delete(expressionNode.Options, OptionSeparator)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression string: %v", err)
	}
	return expressionNode, nil
}

func parseTokensToNodes(initToken *Token, limit bool, defaultLabel string, includeSpans bool, decimal bool, checkLiterals bool) ([]*Node, error) {
	parser := NewParser(initToken, defaultLabel, includeSpans, decimal, checkLiterals)
	nodes := []*Node{}
	for parser.hasNext() {
		node, err := parser.readExpr(makeContext())
		if err != nil {
			return nil, err
		} else {
			nodes = append(nodes, node)
		}
		if limit {
			break
		}
	}
	return nodes, nil
}

func parseToASTArray(input string, limit bool, defaultLabel string, include_spans bool, decodeNumbers bool, checkLiterals bool, colOffset int) ([]*Node, Span, error) {
	// Step 1: Tokenize the input
	initToken, span, terr := tokenizeInput(input, colOffset)
	if terr != nil {
		return nil, Span{}, fmt.Errorf(terr.Message + " (line " + fmt.Sprint(terr.Line) + ", column " + fmt.Sprint(terr.Column) + ")")
	}

	// Step 2: Parse the tokens into nodes
	nodes, err := parseTokensToNodes(initToken, limit, defaultLabel, include_spans, decodeNumbers, checkLiterals)
	if err != nil {
		return nil, Span{}, err
	}

	return nodes, span, nil
}
