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

func NewToken(tokenType TokenType, lexeme string, literal string, line int) *Token {
	return &Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	}
}

func (token *Token) String() string {
	return fmt.Sprintf("%s %s %s", token.TokenType, token.Lexeme, token.Literal)
}

type TokenType int

const (

	// Single-character tokens.

	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	SemiColon
	Slash
	Star

	// One or two character tokens.

	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals.

	Identifier
	String
	Number

	// Keywords.

	And
	Class
	Else
	False
	Fun
	For
	IF
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	Eof
)

func (t TokenType) String() string {
	return []string{
		"(", ")", "{", "}", ",", ".", "-", "+", ";", "/", "*",
		"!", "!=", "=", "==", ">", ">=", "<", "<=",
		"Identifier", "String", "NUmber",
		"and", "class", "else", "false", "fun", "for", "if", "nil", "or", "print", "return", "super", "this", "true", "var", "while",
		"eof",
	}[t]
}
