package parser

import (
	. "compiler/parser"
	"fmt"
	"testing"
)

func Test_LexParser(t *testing.T) {
	var tests = []struct {
		input          string
		expectedTokens []Token
	}{
		{"1+1", []Token{
			{TokenNumber, "1"},
			{TokenPlus, "+"},
			{TokenNumber, "1"},
			{TokenEOF, ""},
		}},
		{"11+1", []Token{
			{TokenNumber, "11"},
			{TokenPlus, "+"},
			{TokenNumber, "1"},
			{TokenEOF, ""},
		}},
		{"11+155", []Token{
			{TokenNumber, "11"},
			{TokenPlus, "+"},
			{TokenNumber, "155"},
			{TokenEOF, ""},
		}},
		{"1-1", []Token{
			{TokenNumber, "1"},
			{TokenMinus, "-"},
			{TokenNumber, "1"},
			{TokenEOF, ""},
		}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexParser%d", index)
		t.Run(testName, func(t *testing.T) {
			tokens := Lex(tt.input)
			if len(tokens) != len(tt.expectedTokens) {
				t.Errorf("Incorrect number of tokens. Expected %d, got %d", len(tt.expectedTokens), len(tokens))
			}

			for i, token := range tokens {
				expectedToken := tt.expectedTokens[i]
				if token.Type != expectedToken.Type {
					t.Errorf("Incorrect token type. Expected %v, got %v", expectedToken.Type, token.Type)
				}
				if token.Value != expectedToken.Value {
					t.Errorf("Incorrect token value. Expected %v, got %v", expectedToken.Value, token.Value)
				}
			}
		})
	}
}

func Test_LexLiteralParser(t *testing.T) {
	var tests = []struct {
		input         string
		expectedToken Token
	}{
		{"1", Token{TokenNumber, "1"}},
		{"2", Token{TokenNumber, "2"}},
		{`"a"`, Token{TokenString, `a`}},
		{`"ab"`, Token{TokenString, `ab`}},
		{`'ab'`, Token{TokenString, `ab`}},
		{`"123"`, Token{TokenString, `123`}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexLiteralParser%d", index)
		t.Run(testName, func(t *testing.T) {
			token, _ := LexLiteral(tt.input, 0)

			if token == nil {
				t.Errorf("Token was nil")
				return
			}

			expectedToken := tt.expectedToken
			if token.Type != expectedToken.Type {
				t.Errorf("Incorrect token type. Expected %v, got %v", expectedToken.Type, token.Type)
			}
			if token.Value != expectedToken.Value {
				t.Errorf("Incorrect token value. Expected %v, got %v", expectedToken.Value, token.Value)
			}
		})
	}
}
