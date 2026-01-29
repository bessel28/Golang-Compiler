package parser

import (
	. "compiler/components/ast"
	. "compiler/components/conditions"
	"compiler/components/literals"
	. "compiler/components/operators"
	"strconv"
)

type TokenType int

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

type Token struct {
	Type  TokenType
	Value string
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(input string) *Parser {
	return &Parser{Lex(input), 0}
}

func (p *Parser) current() *Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}
func (p *Parser) advance() *Token {
	current := p.current()
	p.pos++
	return current
}

func (p *Parser) match(types ...TokenType) bool {
	match := p.NonAdvMatch(types...)
	if match {
		p.advance()
	}
	return match
}

func (p *Parser) NonAdvMatch(types ...TokenType) bool {
	token := p.current()
	if token == nil {
		return false
	}
	for _, t := range types {
		if token.Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) ParseExpr() Expression {
	return p.parseTernary()
}

func (p *Parser) parseTernary() Expression {
	conditionExpr := p.parseOrExpr()

	if p.match(TokenQuestionMark) {
		condition, ok := conditionExpr.(Condition)
		if !ok {
			panic("Invalid condition expression for ternary operation")
		}

		trueExpr := p.ParseExpr()

		if !p.match(TokenColon) {
			panic("Expected ':' from ternary expression")
		}

		falseExpr := p.ParseExpr()
		return &TernaryOperator{Left: trueExpr, Right: falseExpr, Condition: condition}
	}
	return conditionExpr
}

func (p *Parser) parseOrExpr() Expression {
	left := p.parseAndExpr()

	for p.match(TokenOr) {

		right := p.parseAndExpr()
		left = &OrCondition{&BinaryCondition{left, right}}
	}

	return left
}
func (p *Parser) parseAndExpr() Expression {
	left := p.parseEqualityExpr()

	for p.match(TokenAnd) {
		right := p.parseEqualityExpr()
		left = &AndCondition{&BinaryCondition{left, right}}
	}

	return left
}
func (p *Parser) parseEqualityExpr() Expression {
	left := p.parseRelationalExpr()

	if p.match(TokenEQ) {
		right := p.parseRelationalExpr()
		return &EqualCondition{&BinaryCondition{left, right}}
	} else if p.match(TokenNeq) {
		right := p.parseRelationalExpr()
		return &NotEqualCondition{&BinaryCondition{left, right}}
	}
	return left
}
func (p *Parser) parseRelationalExpr() Expression {
	left := p.parseAdditive()

	if p.match(TokenGT) {
		right := p.parseAdditive()
		return &GreaterThanCondition{&BinaryCondition{left, right}}
	} else if p.match(TokenLT) {
		right := p.parseRelationalExpr()
		return &LessThanCondition{&BinaryCondition{left, right}}
	}
	return left
}

func (p *Parser) parseAdditive() Expression {
	left := p.parseMultiplicative()

	for p.NonAdvMatch(TokenPlus, TokenMinus) {
		currentOp := p.advance()
		right := p.parseMultiplicative()
		switch {
		case currentOp == nil:
			break
		case currentOp.Type == TokenPlus:
			left = &BinaryOperator{left, right, AddOperator{}}
		case currentOp.Type == TokenMinus:
			left = &BinaryOperator{left, right, SubOperator{}}
		}
	}
	return left
}

func (p *Parser) parseMultiplicative() Expression {
	left := p.parseUnary()

	for p.NonAdvMatch(TokenMultiply, TokenDivide) {
		currentOp := p.advance()
		right := p.parseUnary()
		switch {
		case currentOp == nil:
			break
		case currentOp.Type == TokenMultiply:
			left = &BinaryOperator{left, right, MultOperator{}}
		case currentOp.Type == TokenDivide:
			left = &BinaryOperator{left, right, DivOperator{}}
		}
	}
	return left
}
func (p *Parser) parseUnary() Expression {
	if p.match(TokenNot) {
		left := p.parseTerm()
		return &NotCondition{&UnaryCondition{left}}
	}
	return p.parseTerm()
}
func (p *Parser) parseTerm() Expression {
	switch {
	case p.match(TokenEOF):
		panic("End of File token being parsed")
	case p.match(TokenOpenParenthesis):
		expr := p.ParseExpr()

		current := p.current()
		if current == nil || current.Type != TokenCloseParenthesis {
			panic("Missing Closed Parenthesis")
		}
		p.advance()
		return expr
	case p.NonAdvMatch(TokenNumber):
		current := p.advance()
		if current == nil {
			panic("Nil token")
		}
		intVal, err := strconv.Atoi(current.Value)
		if err != nil {
			panic("Invalid number")
		}
		return literals.NewLiteral(intVal)
	case p.NonAdvMatch(TokenBool):
		current := p.advance()
		if current == nil {
			panic("Nil token")
		}
		if current.Value == "true" {
			return literals.NewLiteral(true)
		} else if current.Value == "false" {
			return literals.NewLiteral(false)
		}
		panic("Invalid bool value")
	case p.NonAdvMatch(TokenString):
		current := p.advance()
		if current == nil {
			panic("Nil token")
		}
		return literals.NewLiteral(current)
	default:
		panic("Unknown token " + p.current().Value)
	}
}

func Lex(input string) []Token {
	tokens := []Token{}

	for i := 0; i < len(input); i++ {
		token, index := parseToken(input[i:], 0)
		if token == nil {
			continue
		}
		tokens = append(tokens, *token)
		i += index
	}
	tokens = append(tokens, Token{Type: TokenEOF})
	return tokens
}

func parseToken(input string, i int) (*Token, int) {
	switch input[i] {
	case '+':
		return &Token{TokenPlus, "+"}, i
	case '-':
		return &Token{TokenMinus, "-"}, i
	case '*':
		return &Token{TokenMultiply, "*"}, i
	case '/':
		return &Token{TokenDivide, "/"}, i
	case '(':
		return &Token{TokenOpenParenthesis, "("}, i
	case ')':
		return &Token{TokenCloseParenthesis, ")"}, i
	case '>':
		return &Token{TokenGT, ">"}, i
	case '<':
		return &Token{TokenGT, "<"}, i
	case '?':
		return &Token{TokenQuestionMark, "?"}, i
	case ':':
		return &Token{TokenColon, ":"}, i
		//TODO Finish adding the remaining Lex Tokens
	default:
		literal, newIndex := LexLiteral(input, i)
		if literal == nil {
			return nil, i
		}
		return literal, newIndex
	}
}

func LexLiteral(input string, i int) (*Token, int) {
	start := i

	if i+4 < len(input) && input[i:i+4] == "true" {
		return &Token{TokenBool, "true"}, i + 3
	} else if i+5 < len(input) && input[i:i+5] == "false" {
		return &Token{TokenBool, "false"}, i + 5
	} else if input[i] == '"' {
		for i+1 < len(input) && input[i+1] != '"' {
			i++
		}
		if input[i+1] == '"' {
			return &Token{TokenString, input[start : i+2]}, i
		}
		panic("Missing end quote from string token")
	} else if input[i] == '\'' {
		for i+1 < len(input) && input[i+1] != '\'' {
			i++
		}
		if input[i+1] == '\'' {
			return &Token{TokenString, input[start : i+2]}, i
		}
		panic("Missing end single quote from string token")
	} else if input[i] >= '0' && input[i] <= '9' {
		for i+1 < len(input) && input[i+1] >= '0' && input[i+1] <= '9' {
			i++
		}
		return &Token{TokenNumber, input[start : i+1]}, i
	}
	return nil, i + 1
}
