package core

import (
	"reflect"
	"testing"
)

func TestParseOnePlusTwo(t *testing.T) {
	source := "1+2"
	scanner := NewScanner(source)
	expr := NewBinary(
		NewLiteral(1.0),
		&Token{TokenType: Plus, Lexeme: "+", Line: 1},
		NewLiteral(2.0),
	)
	tokens := scanner.scanTokens()
	parser := NewParser(tokens)
	parsedExpr := parser.parse()

	if !reflect.DeepEqual(expr, parsedExpr) {
		t.Errorf("Test parsing `1+2` got %#v want %#v", parsedExpr, expr)
	}
}

func TestParseComplex(t *testing.T) {
	source := "(1+2) * (3 + 4)"
	expr := NewBinary(
		NewGrouping(
			NewBinary(
				NewLiteral(1.0),
				&Token{TokenType: Plus, Lexeme: "+", Line: 1},
				NewLiteral(2.0),
			),
		),
		&Token{TokenType: Star, Lexeme: "*", Line: 1},
		NewGrouping(
			NewBinary(
				NewLiteral(3.0),
				&Token{TokenType: Plus, Lexeme: "+", Line: 1},
				NewLiteral(4.0),
			),
		),
	)
	scanner := NewScanner(source)
	parser := NewParser(scanner.scanTokens())
	parsedExpr := parser.parse()

	if !reflect.DeepEqual(expr, parsedExpr) {
		t.Errorf("Test parsing (1+2)*(3+4) got %#v want %#v", parsedExpr, expr)
	}
}
