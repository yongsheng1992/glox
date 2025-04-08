package core

import "fmt"

type Token struct {
	// TokenType token的类型
	TokenType TokenType
	//
	Lexeme string
	// Literal 字面量
	Literal interface{}
	Line    int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	}
}

func (token *Token) String() string {
	return fmt.Sprintf("%v %v %v", token.TokenType, token.Lexeme, token.Literal)
}

type TokenType int

const (

	// Single-character tokens.

	LPAREN TokenType = iota
	RPAREN
	LBRACE
	RBRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.

	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.

	IDENTIFIER
	STRING
	NUMBER

	// Keywords.

	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	Eof
)

func (t TokenType) String() string {
	return []string{
		"(", ")", "{", "}", ",", ".", "-", "+", ";", "/", "*",
		"!", "!=", "=", "==", ">", ">=", "<", "<=",
		"IDENTIFIER", "STRING", "NUMBER",
		"and", "class", "else", "false", "fun", "for", "if", "nil", "or", "print", "return", "super", "this", "true", "var", "while",
		"eof",
	}[t]
}
