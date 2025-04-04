package core

import (
	"strconv"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"WHILE":  WHILE,
	"Eof":    Eof,
}

// Scanner scan a source get a list of tokens.
type Scanner struct {
	start   int
	current int
	line    int

	source string
	tokens []*Token
}

// isATEnd whether is at the end of the source
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// scanToken
func (s *Scanner) scanToken() {
	ch := s.advance()
	switch ch {
	case '(':
		s.addToken(LPAREN, nil)
	case ')':
		s.addToken(RPAREN, nil)
	case '{':
		s.addToken(LBRACE, nil)
	case '}':
		s.addToken(RBRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '"':
		s.string()
	case ' ', '\t', '\r':
	case '\n':
		s.line++
	default:
		if s.isDigit(ch) {
			s.number()
		} else if s.isAlpha(ch) {
			s.identifier()
		} else {
			panic("Unexpected character. In Line ")
		}
	}
}

func (s *Scanner) identifier() {
	for ; s.isAlphaNumeric(s.peek()); s.advance() {
	}
	if tokeType, exist := keywords[s.source[s.start:s.current]]; exist {
		s.addToken(tokeType, nil)
	} else {
		s.addToken(IDENTIFIER, nil)
	}
}

func (s *Scanner) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (s *Scanner) isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func (s *Scanner) isAlphaNumeric(ch byte) bool {
	return s.isAlpha(ch) || s.isDigit(ch)
}

func (s *Scanner) number() {
	for ; s.isDigit(s.peek()); s.advance() {
	}

	if s.peek() == '.' && s.isDigit(s.nextPeek()) {
		s.advance()
		for ; s.isDigit(s.peek()); s.advance() {
		}
	}
	number, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic(err)
	}
	s.addToken(NUMBER, number)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		panic("Unterminated string")
	}

	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

// advance get a character and increase `current`
func (s *Scanner) advance() byte {
	ch := s.source[s.current]
	s.current++
	return ch
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	token := &Token{
		TokenType: tokenType,
		Lexeme:    string(s.source[s.start:s.current]),
		Literal:   literal,
		Line:      s.line,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}
	return s.source[s.current]
}

func (s *Scanner) nextPeek() byte {
	if s.isAtEnd() {
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) scanTokens() []*Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.addToken(Eof, nil)
	return s.tokens
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		start:   0,
		current: 0,
		line:    1,
		source:  source,
	}
}
