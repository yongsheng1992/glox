package core

import "fmt"

const (
	reportFmt = "[line %d] Error %s : %s"
)

type ParseError struct {
	token   *Token
	message string
}

func NewParseError(token *Token, message string) *ParseError {
	return &ParseError{
		token:   token,
		message: message,
	}
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf(reportFmt, pe.token.Line, pe.token.Lexeme, pe.message)
}

type RuntimeError struct {
}
