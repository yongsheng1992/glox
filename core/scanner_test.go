package core

import (
	"reflect"
	"testing"
)

type testcase struct {
	input string
	want  []*Token
}

func TestSingleCharacterTokens(t *testing.T) {
	testcases := []testcase{
		{
			input: `var imAVariable = "here is my value";`,
			want: []*Token{
				{
					TokenType: Var,
					Lexeme:    "var",
					Literal:   nil,
					Line:      1,
				},
				{
					TokenType: Identifier,
					Lexeme:    "imAVariable",
					Literal:   nil,
					Line:      1,
				},
				{
					TokenType: Equal,
					Lexeme:    "=",
					Literal:   nil,
					Line:      1,
				},
				{
					TokenType: String,
					Lexeme:    `"here is my value"`,
					Literal:   "here is my value",
					Line:      1,
				},
				{
					TokenType: SemiColon,
					Lexeme:    `;`,
					Literal:   nil,
					Line:      1,
				},
			},
		},
	}

	for _, tc := range testcases {
		scanner := NewScanner(tc.input)
		tokens := scanner.scanTokens()
		if !reflect.DeepEqual(tokens, tc.want) {
			t.Errorf("TestSingleCharacterTokens got %v want %v", tokens, tc.want)
		}

	}
}

func TestMultipleCharacterTokens(t *testing.T) {

}

func TestLiteralTokens(t *testing.T) {

}

func TestKeywordTokens(t *testing.T) {

}
