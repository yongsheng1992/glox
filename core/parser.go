package core

type Parser struct {
	tokens  []*Token
	current int
}

func NewParser(token []*Token) *Parser {
	return &Parser{
		tokens:  token,
		current: 0,
	}
}

func (parser *Parser) parse() Expr {
	return parser.expression()
}

func (parser *Parser) expression() Expr {
	return parser.term()
}

func (parser *Parser) term() Expr {
	expr := parser.factor()

	for parser.match(MINUS, PLUS) {
		operator := parser.previous()
		expr = NewBinary(expr, operator, parser.factor())
	}

	return expr
}

func (parser *Parser) factor() Expr {
	expr := parser.unary()

	for parser.match(STAR, SLASH) {
		operator := parser.previous()
		expr = NewBinary(expr, operator, parser.unary())
	}

	return expr
}

func (parser *Parser) unary() Expr {
	if parser.match(BANG, MINUS) {
		return NewUnary(parser.previous(), parser.unary())
	} else {
		return parser.primary()
	}
}

func (parser *Parser) primary() Expr {
	if parser.match(NUMBER) {
		return NewLiteral(parser.previous().Literal)
	} else if parser.match(STRING) {
		panic(NewParseError(parser.peek(), "STRING is currently unsupported!"))
	} else if parser.match(LPAREN) {
		expr := parser.expression()
		parser.consume(RPAREN, "Expect ')' after expression.")
		return NewGrouping(expr)
	} else {
		panic(NewParseError(parser.peek(), "Expect expression."))
	}
}

// match
func (parser *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if parser.check(tokenType) {
			parser.advance()
			return true
		}
	}
	return false
}

func (parser *Parser) check(tokenType TokenType) bool {
	if parser.isAtEnd() {
		return false
	}
	return parser.peek().TokenType == tokenType
}

// peek
func (parser *Parser) peek() *Token {
	if parser.isAtEnd() {
		return &Token{TokenType: Eof}
	}
	return parser.tokens[parser.current]
}

// advance go through the parse process.
func (parser *Parser) advance() *Token {
	if !parser.isAtEnd() {
		parser.current += 1
	}
	return parser.previous()
}

// previous
func (parser *Parser) previous() *Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) consume(tokenType TokenType, message string) *Token {
	if parser.check(tokenType) {
		return parser.advance()
	}
	panic(NewParseError(parser.peek(), message))
}

// isAtEnd
func (parser *Parser) isAtEnd() bool {
	return parser.current >= len(parser.tokens)
}
