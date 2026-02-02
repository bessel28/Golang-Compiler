package parser

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	"compiler/src/components/literals"
	. "compiler/src/components/operators"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type  TokenType
	Value string
}

func mapTokens(tokens []Token) []string {
	var stringVals = make([]string, len(tokens))
	for i, t := range tokens {
		if t.Type == TokenEOF {
			stringVals[i] = "-EOF-"
			continue
		}
		stringVals[i] = t.Value
	}
	return stringVals
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(input string) (*Parser, error) {
	tokens, err := Lex(input)
	return &Parser{tokens, 0}, err
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

func (p *Parser) match(types ...TokenType) (bool, error) {
	match, err := p.nonAdvMatch(types...)
	if match {
		p.advance()
	}
	return match, err
}

func (p *Parser) nonAdvMatch(types ...TokenType) (bool, error) {
	token := p.current()
	if token == nil {
		return false, errors.New("No tokens left to match")
	}
	for _, t := range types {
		if token.Type == t {
			return true, nil
		}
	}
	return false, nil
}

// ParseExpr Attempts to parse all the tokens in a Parser's tokens field.
// NOTE : This function expects ALL tokens to be parsed and will return an error if there are unparsed tokens remaining
func (p *Parser) ParseExpr() (Expression, error) {
	expr, exprErr := p.parseTernary()
	if current := p.current(); current != nil && current.Type != TokenEOF {
		unparsedTokenDebug := fmt.Sprintf("Current Token: %+v\nTotal tokens remaining: %v\n%v",
			*current, len(p.tokens)-p.pos, strings.Join(mapTokens(p.tokens[p.pos:]), ", "))
		return expr, errors.New(fmt.Sprintf("Unparsed tokens remaining: %v", unparsedTokenDebug))
	}
	return expr, exprErr
}

func (p *Parser) parseTernary() (Expression, error) {
	conditionExpr, conditionErr := p.parseOrExpr()
	if conditionErr != nil {
		return nil, conditionErr
	}

	matchFound, questionMarkErr := p.match(TokenQuestionMark)
	if questionMarkErr != nil {
		return nil, questionMarkErr
	}

	if matchFound {
		condition, ok := conditionExpr.(Condition)
		if !ok {
			return nil, errors.New("Invalid condition expression for ternary operation")
		}

		trueExpr, trueErr := p.parseTernary()
		if trueErr != nil {
			return nil, trueErr
		}

		colonMatch, colonErr := p.match(TokenColon)
		if colonErr != nil {
			return nil, colonErr
		}
		if !colonMatch {
			return nil, errors.New("Missing Colon in ternary expression")
		}

		falseExpr, falseErr := p.parseTernary()
		if falseErr != nil {
			return nil, falseErr
		}
		return &TernaryOperator{Left: trueExpr, Right: falseExpr, Condition: condition}, nil
	}
	return conditionExpr, nil
}

func (p *Parser) parseOrExpr() (Expression, error) {
	left, leftErr := p.parseAndExpr()
	if leftErr != nil {
		return nil, leftErr
	}

	for {
		orMatch, orErr := p.match(TokenOr)
		if orErr != nil {
			return nil, orErr
		}
		if !orMatch {
			break
		}
		right, rightErr := p.parseAndExpr()
		if rightErr != nil {
			return nil, rightErr
		}
		left = &OrCondition{&BinaryCondition{left, right}}
	}

	return left, nil
}

func (p *Parser) parseAndExpr() (Expression, error) {
	left, leftErr := p.parseEqualityExpr()
	if leftErr != nil {
		return nil, leftErr
	}

	for {
		andMatch, andErr := p.match(TokenAnd)
		if andErr != nil {
			return nil, andErr
		}
		if !andMatch {
			break
		}
		right, rightErr := p.parseEqualityExpr()
		if rightErr != nil {
			return nil, rightErr
		}
		left = &AndCondition{&BinaryCondition{left, right}}
	}

	return left, nil
}
func (p *Parser) parseEqualityExpr() (Expression, error) {
	left, leftErr := p.parseRelationalExpr()
	if leftErr != nil {
		return nil, leftErr
	}

	eqMatch, eqErr := p.match(TokenEQ)
	if eqErr != nil {
		return nil, eqErr
	}
	if eqMatch {
		right, rightErr := p.parseRelationalExpr()
		if rightErr != nil {
			return nil, rightErr
		}
		return &EqualCondition{&BinaryCondition{left, right}}, nil
	}
	neqMatch, neqErr := p.match(TokenNeq)
	if neqErr != nil {
		return nil, neqErr
	}
	if neqMatch {
		right, rightErr := p.parseRelationalExpr()
		if rightErr != nil {
			return nil, rightErr
		}
		return &NotEqualCondition{&BinaryCondition{left, right}}, nil
	}
	return left, nil
}
func (p *Parser) parseRelationalExpr() (Expression, error) {
	left, leftErr := p.parseAdditive()
	if leftErr != nil {
		return nil, leftErr
	}

	gtMatch, gtErr := p.match(TokenGT)
	if gtErr != nil {
		return nil, gtErr
	}
	if gtMatch {
		right, rightErr := p.parseAdditive()
		if rightErr != nil {
			return nil, rightErr
		}

		return &GreaterThanCondition{&BinaryCondition{left, right}}, nil
	}

	ltMatch, ltErr := p.match(TokenGT)
	if ltErr != nil {
		return nil, ltErr
	}
	if ltMatch {
		right, rightErr := p.parseRelationalExpr()
		if rightErr != nil {
			return nil, rightErr
		}

		return &LessThanCondition{&BinaryCondition{left, right}}, nil
	}
	return left, nil
}

func (p *Parser) parseAdditive() (Expression, error) {
	left, leftErr := p.parseMultiplicative()
	if leftErr != nil {
		return nil, leftErr
	}

	for {
		matchFound, matchErr := p.nonAdvMatch(TokenPlus, TokenMinus)
		if matchErr != nil {
			return nil, matchErr
		}
		if !matchFound {
			break
		}

		currentOp := p.advance()
		right, rightErr := p.parseMultiplicative()
		if rightErr != nil {
			return nil, rightErr
		}

		switch {
		case currentOp == nil:
			break
		case currentOp.Type == TokenPlus:
			left = &BinaryOperator{left, right, AddOperator{}}
		case currentOp.Type == TokenMinus:
			left = &BinaryOperator{left, right, SubOperator{}}
		}
	}
	return left, nil
}

func (p *Parser) parseMultiplicative() (Expression, error) {
	left, leftErr := p.parseUnary()
	if leftErr != nil {
		return nil, leftErr
	}

	for {
		matchFound, matchErr := p.nonAdvMatch(TokenMultiply, TokenDivide)
		if matchErr != nil {
			return nil, matchErr
		}
		if !matchFound {
			break
		}
		currentOp := p.advance()
		right, rightErr := p.parseUnary()
		if rightErr != nil {
			return nil, rightErr
		}

		switch {
		case currentOp == nil:
			break
		case currentOp.Type == TokenMultiply:
			left = &BinaryOperator{left, right, MultOperator{}}
		case currentOp.Type == TokenDivide:
			left = &BinaryOperator{left, right, DivOperator{}}
		}
	}
	return left, nil
}
func (p *Parser) parseUnary() (Expression, error) {
	notMatch, notErr := p.match(TokenNot)
	if notErr != nil {
		return nil, notErr
	}

	if notMatch {
		left, leftErr := p.parseTerm()
		if leftErr != nil {
			return nil, leftErr
		}
		return &NotCondition{&UnaryCondition{left}}, nil
	}
	return p.parseTerm()
}

func (p *Parser) parseTerm() (Expression, error) {
	eofMatch, eofErr := p.match(TokenEOF)
	if eofErr != nil {
		return nil, eofErr
	}
	if eofMatch {
		return nil, errors.New("Unexpected EOF being parsed")
	}

	openParaMatch, openParaErr := p.match(TokenOpenParenthesis)
	if openParaErr != nil {
		return nil, openParaErr
	}
	if openParaMatch {
		expr, exprErr := p.parseTernary()
		if exprErr != nil {
			return nil, exprErr
		}

		closeParaMatch, closeParaErr := p.match(TokenCloseParenthesis)
		if closeParaErr == nil {
			return nil, closeParaErr
		}
		if !closeParaMatch {
			return nil, errors.New("Missing Closed Parenthesis")
		}
		return expr, nil
	}

	numberMatch, numberErr := p.nonAdvMatch(TokenNumber)
	if numberErr != nil {
		return nil, numberErr
	}
	if numberMatch {
		current := p.advance()
		if current == nil {
			return nil, errors.New("Nil token")
		}
		intVal, err := strconv.Atoi(current.Value)
		if err != nil {
			return nil, errors.New(fmt.Sprintf(`Invalid number %v`, current.Value))
		}
		return literals.NewLiteral(intVal), nil
	}

	boolMatch, boolErr := p.nonAdvMatch(TokenBool)
	if boolErr != nil {
		return nil, boolErr
	}
	if boolMatch {
		current := p.advance()
		if current == nil {
			return nil, errors.New("Nil token")
		}
		if current.Value == "true" {
			return literals.NewLiteral(true), nil
		} else if current.Value == "false" {
			return literals.NewLiteral(false), nil
		}
		return nil, errors.New(fmt.Sprintf("Invalid bool value %v", current.Value))
	}

	stringMatch, stringErr := p.nonAdvMatch(TokenString)
	if stringErr != nil {
		return nil, stringErr
	}
	if stringMatch {
		current := p.advance()
		if current == nil {
			return nil, errors.New("Nil token")
		}
		return literals.NewLiteral(current.Value), nil
	}

	return nil, errors.New("Unknown token " + p.current().Value)

}
