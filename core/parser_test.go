package core

import (
	"reflect"
	"testing"
)

func TestParseOnePlusTwoExpression(t *testing.T) {
	source := "1+2"
	scanner := NewScanner(source)
	expr := NewBinary(
		NewLiteral(1.0),
		&Token{TokenType: PLUS, Lexeme: "+", Line: 1},
		NewLiteral(2.0),
	)
	tokens := scanner.scanTokens()
	parser := NewParser(tokens)
	parsedExpr := parser.expression()

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
				&Token{TokenType: PLUS, Lexeme: "+", Line: 1},
				NewLiteral(2.0),
			),
		),
		&Token{TokenType: STAR, Lexeme: "*", Line: 1},
		NewGrouping(
			NewBinary(
				NewLiteral(3.0),
				&Token{TokenType: PLUS, Lexeme: "+", Line: 1},
				NewLiteral(4.0),
			),
		),
	)
	scanner := NewScanner(source)
	parser := NewParser(scanner.scanTokens())
	parsedExpr := parser.expression()

	if !reflect.DeepEqual(expr, parsedExpr) {
		t.Errorf("Test parsing (1+2)*(3+4) got %#v want %#v", parsedExpr, expr)
	}
}

func TestVarStmt(t *testing.T) {
	stmt := NewVarStmt(
		NewToken(IDENTIFIER, "a", nil, 1),
		NewLiteral("1"),
	)

	source := `var a = "1";`
	scanner := NewScanner(source)
	parser := NewParser(scanner.scanTokens())
	stmts := parser.parse()
	if len(stmts) != 1 {
		t.Fatalf(`Test parsing 'var a = "1";' got %#v want %#v`, stmts, stmt)
	} else {
		if !reflect.DeepEqual(stmt, stmts[0]) {
			t.Fatalf(`Test parsing 'var a = "1"' got %#v want %#v`, stmts[0], stmt)
		}
	}

}
