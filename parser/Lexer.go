package parser

type TokenType int
type Lexer func(input string, start int) (*string, int, bool)

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
	return func(input string, start int) (*string, int, bool) {
		if start+len(symbol) > len(input) {
			return nil, start, false
		}
		return ptr(symbol), start + len(symbol) - 1, input[start:start+len(symbol)] == symbol
	}
}

func Lex(input string) []Token {
	var tokens []Token
	i := 0
	for ; i < len(input); i++ {
		token, index := parseToken(input[i:], 0)
		if token == nil {
			continue
		}
		tokens = append(tokens, *token)
		i += index
	}
	if data, _, ok := LexerEOF.Lexer(input, i); ok {
		tokens = append(tokens, Token{LexerEOF.Type, *data})
	}
	return tokens
}

func parseToken(input string, i int) (*Token, int) {
	lexers := []LexerData{
		LexerPlus, LexerMinus, LexerMultiply, LexerDivide, LexerOpenParenthesis, LexerCloseParenthesis,
		LexerQuestionMark, LexerColon, LexerEQ, LexerNeq, LexerGT, LexerLT, LexerOr, LexerAnd, LexerNot,
	}
	if input[i] == ' ' {
		return nil, i
	}
	for _, lexer := range lexers {
		if data, index, ok := lexer.Lexer(input, i); ok {
			return &Token{lexer.Type, *data}, index
		}
	}

	literal, newIndex := LexLiteral(input, i)
	if literal == nil {
		return nil, i
	}
	return literal, newIndex
}

func LexLiteral(input string, i int) (*Token, int) {
	lexers := []LexerData{LexerBool, LexerString, LexerNumber}

	for _, lx := range lexers {
		if data, index, ok := lx.Lexer(input, i); ok {
			return &Token{lx.Type, *data}, index
		}
	}

	return nil, i + 1
}

func lexBool(input string, i int) (*string, int, bool) {
	switch {
	case i+4 <= len(input) && input[i:i+4] == "true":
		return ptr("true"), i + 3, true
	case i+5 <= len(input) && input[i:i+5] == "false":
		return ptr("false"), i + 4, true
	}
	return nil, i, false
}

func lexString(input string, i int) (*string, int, bool) {
	start := i

	if input[i] == '"' || input[i] == '\'' {
		quoteType := input[i]
		for i+1 <= len(input) && input[i+1] != quoteType {
			i++
		}
		if input[i+1] == quoteType {
			return ptr(input[start+1 : i+1]), i + 1, true
		}
		panic("Missing end quote from string token")
	}
	return nil, start, false
}

func lexInt(input string, i int) (*string, int, bool) {
	start := i
	if input[i] >= '0' && input[i] <= '9' {
		for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' {
			i++
		}
		return ptr(input[start : i+1]), i, true
	}
	return nil, start, false
}

func ptr[T any](v T) *T { return &v }
