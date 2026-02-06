package expression

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/components/variables"
	"compiler/src/parser"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Token struct {
	Type  TokenType
	Value string
	Pos   int
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
	Tokens      []Token
	pos         int
	VarSymbols  *parser.VarSymbolTable
	FuncSymbols *parser.FuncParserSymbolTable
}

func (p *Parser) current() *Token {
	if p.pos >= len(p.Tokens) {
		return nil
	}
	return &p.Tokens[p.pos]
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
		return false, errors.New("No Tokens left to match")
	}
	for _, t := range types {
		if token.Type == t {
			return true, nil
		}
	}
	return false, nil
}

func NewParser(input string) (*Parser, error) {
	tokens, err := Lex(input)
	return &Parser{tokens, 0, &parser.VarSymbolTable{}, &parser.FuncParserSymbolTable{}}, err
}

// ParseExpr Attempts to parse all the Tokens in a Parser's Tokens field.
// NOTE : This function expects ALL Tokens to be parsed and will return an error if there are unparsed Tokens remaining
func (p *Parser) ParseExpr() (Expression, error) {
	expr, exprErr := p.parseTernary()
	if exprErr != nil {
		return nil, exprErr
	}
	if current := p.current(); current != nil && current.Type != TokenEOF {
		unparsedTokenDebug := fmt.Sprintf("Current Token: %+v\nTotal Tokens remaining: %v\n%v",
			*current, len(p.Tokens)-p.pos, strings.Join(mapTokens(p.Tokens[p.pos:]), ", "))
		return expr, errors.New(fmt.Sprintf("Unparsed Tokens remaining: %v", unparsedTokenDebug))
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
		return &TernaryExpr{IfElseOperator{}, condition, trueExpr, falseExpr, nil}, nil
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
		left = &BinaryCondition{OrCondition{}, left, right, nil}
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
		left = &BinaryCondition{AndCondition{}, left, right, nil}
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
		return &BinaryCondition{EqualCondition{}, left, right, nil}, nil
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
		return &BinaryCondition{NotEqualCondition{}, left, right, nil}, nil
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

		return &BinaryCondition{GreaterThanCondition{}, left, right, nil}, nil
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

		return &BinaryCondition{LessThanCondition{}, left, right, nil}, nil
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
			left = &BinaryExpr{AddOperator{}, left, right, nil}
		case currentOp.Type == TokenMinus:
			left = &BinaryExpr{SubOperator{}, left, right, nil}
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
			left = &BinaryExpr{MultOperator{}, left, right, nil}
		case currentOp.Type == TokenDivide:
			left = &BinaryExpr{DivOperator{}, left, right, nil}
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
		left, leftErr := p.parseCall()
		if leftErr != nil {
			return nil, leftErr
		}
		return &UnaryCondition{NotCondition{}, left, nil}, nil
	}
	return p.parseCall()
}

func (p *Parser) parseCall() (Expression, error) {
	start := p.pos
	expr, exprErr := p.parseTerm()
	if exprErr != nil {
		return nil, exprErr
	}

	openParenthesis, openParaErr := p.match(TokenOpenParenthesis)
	if openParaErr != nil {
		return nil, openParaErr
	}
	if openParenthesis {
		variableExpr, variableType := expr.(*VariableExpr)
		if variableType && reflect.TypeOf(variableExpr.Type) != reflect.TypeOf(FunctionSignature{}) {
			return nil, errors.New(fmt.Sprintf("Variable %v is not a callable type", variableExpr.Identifier))
		}
		functionExpr, functionType := expr.(*FunctionExpr)
		if !functionType && !variableType {
			return nil, errors.New("Uncallable Expression")
		}

		var args []Expression

		closeParenthesis, closeParaErr := p.nonAdvMatch(TokenCloseParenthesis)
		if closeParaErr != nil {
			return nil, closeParaErr
		}
		if !closeParenthesis {
			for {
				param, paramErr := p.parseTernary()
				if paramErr != nil {
					return nil, paramErr
				}
				args = append(args, param)

				commaMatch, commaErr := p.match(TokenComma)
				if commaErr != nil {
					return nil, commaErr
				}
				if !commaMatch {
					break
				}
			}
		}

		closeParenthesis, closeParaErr = p.match(TokenCloseParenthesis)

		if closeParaErr != nil {
			return nil, closeParaErr
		}
		if !closeParenthesis {
			return nil, errors.New("Missing closing Parenthesis in function call")
		}

		return &CallExpr{functionExpr, args, start}, nil
	}
	return expr, nil
}

func (p *Parser) parseTerm() (Expression, error) {
	start := p.pos
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
		if closeParaErr != nil {
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
		return NewLiteral(intVal, &start), nil
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
			return NewLiteral(true, &start), nil
		} else if current.Value == "false" {
			return NewLiteral(false, &start), nil
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
		return NewLiteral(current.Value, &start), nil
	}

	identifierMatch, identifierErr := p.nonAdvMatch(TokenIdentifier)
	if identifierErr != nil {
		return nil, identifierErr
	}
	if identifierMatch {
		identifier := p.advance()
		if identifier == nil {
			return nil, errors.New("Nil token")
		}

		if p.VarSymbols.IsDeclared(identifier.Value) {
			varType := p.VarSymbols.Get(identifier.Value)
			if varType == nil {
				return nil, errors.New("Variable declaration not found")
			}
			return &VariableExpr{identifier.Value, varType, start}, nil
		}

		funcIdentifier := p.FuncSymbols.Get(identifier.Value)
		return &FunctionExpr{funcIdentifier, nil, start}, nil
	}

	return nil, errors.New("Unknown term")
}
