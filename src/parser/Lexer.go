package parser

import "errors"

type TokenType int
type Lexer func(input string, start int) (*string, int, bool, error)

type LexerData struct {
	Type TokenType
	Lexer
}

const (
	TokenEOF TokenType = iota
	TokenNumber
	TokenBool
	TokenString
	TokenPlus
	TokenMinus
	TokenMultiply
	TokenDivide
	TokenOpenParenthesis
	TokenCloseParenthesis
	TokenQuestionMark
	TokenColon
	TokenEQ
	TokenNeq
	TokenGT
	TokenLT
	TokenOr
	TokenAnd
	TokenNot
)

var (
	LexerEOF              = LexerData{TokenEOF, simpleLexer("")}
	LexerNumber           = LexerData{TokenNumber, lexInt}
	LexerBool             = LexerData{TokenBool, lexBool}
	LexerString           = LexerData{TokenString, lexString}
	LexerPlus             = LexerData{TokenPlus, simpleLexer("+")}
	LexerMinus            = LexerData{TokenMinus, simpleLexer("-")}
	LexerMultiply         = LexerData{TokenMultiply, simpleLexer("*")}
	LexerDivide           = LexerData{TokenDivide, simpleLexer("/")}
	LexerOpenParenthesis  = LexerData{TokenOpenParenthesis, simpleLexer("(")}
	LexerCloseParenthesis = LexerData{TokenCloseParenthesis, simpleLexer(")")}
	LexerQuestionMark     = LexerData{TokenQuestionMark, simpleLexer("?")}
	LexerColon            = LexerData{TokenColon, simpleLexer(":")}
	LexerEQ               = LexerData{TokenEQ, simpleLexer("==")}
	LexerNeq              = LexerData{TokenNeq, simpleLexer("!=")}
	LexerGT               = LexerData{TokenGT, simpleLexer(">")}
	LexerLT               = LexerData{TokenLT, simpleLexer("<")}
	LexerOr               = LexerData{TokenOr, simpleLexer("||")}
	LexerAnd              = LexerData{TokenAnd, simpleLexer("&&")}
	LexerNot              = LexerData{TokenNot, simpleLexer("!")}
)

func simpleLexer(symbol string) Lexer {
	return func(input string, start int) (*string, int, bool, error) {
		if start+len(symbol) > len(input) {
			return nil, start, false, nil
		}
		return ptr(symbol), start + len(symbol) - 1, input[start:start+len(symbol)] == symbol, nil
	}
}

func Lex(input string) ([]Token, error) {
	var tokens []Token
	i := 0
	for ; i < len(input); i++ {
		token, index, err := parseToken(input[i:], 0)
		if err != nil {
			return nil, err
		}
		if token == nil {
			continue
		}
		tokens = append(tokens, *token)
		i += index
	}
	data, _, ok, err := LexerEOF.Lexer(input, i)
	if err != nil {
		return nil, err
	}
	if ok {
		tokens = append(tokens, Token{LexerEOF.Type, *data})
	}
	return tokens, nil
}

func parseToken(input string, i int) (*Token, int, error) {
	lexers := []LexerData{
		LexerPlus, LexerMinus, LexerMultiply, LexerDivide, LexerOpenParenthesis, LexerCloseParenthesis,
		LexerQuestionMark, LexerColon, LexerEQ, LexerNeq, LexerGT, LexerLT, LexerOr, LexerAnd, LexerNot,
	}
	if input[i] == ' ' {
		return nil, i, nil
	}
	for _, lexer := range lexers {
		data, index, ok, err := lexer.Lexer(input, i)
		if err != nil {
			return nil, i, err
		}
		if ok {
			return &Token{lexer.Type, *data}, index, nil
		}
	}

	literal, newIndex, err := LexLiteral(input, i)
	if err != nil {
		return nil, i, err
	}
	if literal == nil {
		return nil, i, nil
	}
	return literal, newIndex, nil
}

func LexLiteral(input string, i int) (*Token, int, error) {
	lexers := []LexerData{LexerBool, LexerString, LexerNumber}

	for _, lx := range lexers {
		data, index, ok, err := lx.Lexer(input, i)
		if err != nil {
			return nil, i, err
		}
		if ok {
			return &Token{lx.Type, *data}, index, nil
		}
	}

	return nil, i + 1, nil
}

func lexBool(input string, i int) (*string, int, bool, error) {
	switch {
	case i+4 <= len(input) && input[i:i+4] == "true":
		return ptr("true"), i + 3, true, nil
	case i+5 <= len(input) && input[i:i+5] == "false":
		return ptr("false"), i + 4, true, nil
	}
	return nil, i, false, nil
}

func lexString(input string, i int) (*string, int, bool, error) {
	start := i

	if input[i] == '"' || input[i] == '\'' {
		quoteType := input[i]
		for i+1 <= len(input) && input[i+1] != quoteType {
			i++
		}
		if input[i+1] == quoteType {
			return ptr(input[start+1 : i+1]), i + 1, true, nil
		}
		return nil, start, false, errors.New("Missing closing quotes from string")
	}
	return nil, start, false, nil
}

func lexInt(input string, i int) (*string, int, bool, error) {
	start := i
	if input[i] >= '0' && input[i] <= '9' {
		for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' {
			i++
		}
		return ptr(input[start : i+1]), i, true, nil
	}
	return nil, start, false, nil
}

func ptr[T any](v T) *T { return &v }
