package core

type Parser struct {
	tokens  []*Token
	current int
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (parser *Parser) parse() []Stmt {
	stmts := make([]Stmt, 0)
	for !parser.isAtEnd() {
		stmts = append(stmts, parser.declaration())
	}
	return stmts
}

func (parser *Parser) statement() Stmt {
	if parser.match(PRINT) {
		return parser.printStmt()
	} else {
		return parser.exprStmt()
	}
}

func (parser *Parser) declaration() Stmt {
	if parser.match(VAR) {
		return parser.varDeclaration()
	} else {
		return parser.statement()
	}
}

func (parser *Parser) varDeclaration() Stmt {
	if parser.match(IDENTIFIER) {
		identifier := parser.previous()
		var expr Expr
		if parser.match(EQUAL) {
			expr = parser.expression()
		}
		parser.consume(SEMICOLON, "Expect ';' after expression.")
		return NewVarStmt(identifier, expr)
	}
	panic("Expect identifier after 'var' keyword.")
}

func (parser *Parser) exprStmt() Stmt {
	expr := parser.expression()
	parser.consume(SEMICOLON, "Expect ';' after expression.")
	return NewExprStmt(expr)
}

func (parser *Parser) printStmt() Stmt {
	expr := parser.expression()
	parser.consume(SEMICOLON, "Expect ';' after value.")
	return NewPrint(expr)
}
func (parser *Parser) expression() Expr {
	return parser.equality()
}

func (parser *Parser) equality() Expr {
	left := parser.comparison()

	for parser.match(EQUAL_EQUAL, BANG_EQUAL) {
		right := parser.comparison()
		operator := parser.previous()
		return NewBinary(left, operator, right)
	}

	return left
}

func (parser *Parser) comparison() Expr {
	left := parser.term()

	for parser.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		right := parser.term()
		operator := parser.previous()
		return NewBinary(left, operator, right)
	}

	return left
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
		return NewLiteral(parser.previous().Literal)
	} else if parser.match(LPAREN) {
		expr := parser.expression()
		parser.consume(RPAREN, "Expect ')' after expression.")
		return NewGrouping(expr)
	} else if parser.match(IDENTIFIER) {
		token := parser.previous()
		return NewVarExpr(token)
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
	return parser.peek().TokenType == Eof
}
