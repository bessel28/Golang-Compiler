package parsing

import (
	. "compiler/src/parser"
	"fmt"
	"testing"
)

func Test_LexParser(t *testing.T) {
	var tests = []struct {
		input          string
		expectedTokens []Token
	}{
		{"1+1", []Token{
			{TokenNumber, "1", -1},
			{TokenPlus, "+", -1},
			{TokenNumber, "1", -1},
			{TokenEOF, "", -1},
		}},
		{"11+1", []Token{
			{TokenNumber, "11", -1},
			{TokenPlus, "+", -1},
			{TokenNumber, "1", -1},
			{TokenEOF, "", -1},
		}},
		{"11+155", []Token{
			{TokenNumber, "11", -1},
			{TokenPlus, "+", -1},
			{TokenNumber, "155", -1},
			{TokenEOF, "", -1},
		}},
		{"1-1", []Token{
			{TokenNumber, "1", -1},
			{TokenMinus, "-", -1},
			{TokenNumber, "1", -1},
			{TokenEOF, "", -1},
		}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexParser%d", index)
		t.Run(testName, func(t *testing.T) {
			tokens, err := Lex(tt.input)
			if err != nil {
				t.Fatalf("Lex error: %s", err)
			}

			if len(tokens) != len(tt.expectedTokens) {
				t.Fatalf("Incorrect number of tokens. Expected %d, got %d", len(tt.expectedTokens), len(tokens))
			}

			for i, token := range tokens {
				expectedToken := tt.expectedTokens[i]
				if token.Type != expectedToken.Type {
					t.Fatalf("Incorrect token type. Expected %v, got %v", expectedToken.Type, token.Type)
				}
				if token.Value != expectedToken.Value {
					t.Fatalf("Incorrect token value. Expected %v, got %v", expectedToken.Value, token.Value)
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
		{"1", Token{TokenNumber, "1", -1}},
		{"2", Token{TokenNumber, "2", -1}},
		{`"a"`, Token{TokenString, `a`, -1}},
		{`"ab"`, Token{TokenString, `ab`, -1}},
		{`'ab'`, Token{TokenString, `ab`, -1}},
		{`"123"`, Token{TokenString, `123`, -1}},
		{`abc`, Token{TokenIdentifier, `abc`, -1}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexLiteralParser%d", index)
		t.Run(testName, func(t *testing.T) {
			tokens, err := Lex(tt.input)
			if err != nil {
				t.Fatalf("LexLiteral error: %s", err)
			}

			if tokens == nil {
				t.Fatalf("Token was nil")
				return
			}

			expectedToken := tt.expectedToken
			if len(tokens) != 2 {
				t.Fatalf("Incorrect number of tokens. Expected %d, got %d", 2, len(tokens))
			}
			if tokens[0].Type != expectedToken.Type {
				t.Fatalf("Incorrect token type. Expected %v, got %v", expectedToken.Type, tokens[0].Type)
			}
			if tokens[0].Value != expectedToken.Value {
				t.Fatalf("Incorrect token value. Expected %v, got %v", expectedToken.Value, tokens[0].Value)
			}
		})
	}
}
