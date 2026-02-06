package parsing

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	"compiler/src/components/statements"
	. "compiler/src/parser"
	. "compiler/src/parser/expression"
	. "compiler/src/runtime/value"
	. "compiler/test"
	"fmt"
	"testing"
)

func Test_Parser(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1+1", &BinaryExpr{Left: NewLiteral(1, nil), Right: NewLiteral(1, nil), Op: AddOperator{}}},
		{"1+1*2", &BinaryExpr{Left: NewLiteral(1, nil), Right: &BinaryExpr{Left: NewLiteral(1, nil), Right: NewLiteral(2, nil), Op: MultOperator{}}, Op: AddOperator{}}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %s", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %s", parserErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_ParserIntValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1+1", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}},
		{"1+1*2", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), &BinaryExpr{MultOperator{}, NewLiteral(1, nil), NewLiteral(2, nil), nil}, nil}},
		{"2*2+1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2, nil), NewLiteral(2, nil), nil}, NewLiteral(1, nil), nil}},
		{"(1+1)*2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}, NewLiteral(2, nil), nil}},
		{"5>4?1:2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5, nil), NewLiteral(4, nil), nil}, NewLiteral(1, nil), NewLiteral(2, nil), nil}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexIntParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_ParserIntSpacesValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 + 1", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}},
		{"1 + 1 * 2", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), &BinaryExpr{MultOperator{}, NewLiteral(1, nil), NewLiteral(2, nil), nil}, nil}},
		{"2 * 2 + 1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2, nil), NewLiteral(2, nil), nil}, NewLiteral(1, nil), nil}},
		{"( 1 + 1 ) * 2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}, NewLiteral(2, nil), nil}},
		{"5 > 4 ? 1 : 2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5, nil), NewLiteral(4, nil), nil}, NewLiteral(1, nil), NewLiteral(2, nil), nil}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexIntSpacesParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_ParserStringValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`"a"+"b"`, &BinaryExpr{AddOperator{}, NewLiteral("a", nil), NewLiteral("b", nil), nil}},
		{`"ab"+"c"`, &BinaryExpr{AddOperator{}, NewLiteral("ab", nil), NewLiteral("c", nil), nil}},
		{`"ab"+""`, &BinaryExpr{AddOperator{}, NewLiteral("ab", nil), NewLiteral("", nil), nil}},
		{`""+"c"`, &BinaryExpr{AddOperator{}, NewLiteral("", nil), NewLiteral("c", nil), nil}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexStringParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_ParserBoolValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`true || false`, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil), nil}},
		{`true && false`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(false, nil), nil}},
		{`true && true`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(true, nil), nil}},
		{`(true || false) && true`, &BinaryCondition{AndCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil), nil}, NewLiteral(true, nil), nil}},
		{`true && (true || false)`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil), nil}, nil}},
		{`true || true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(true, nil), nil}, NewLiteral(false, nil), nil}},
		{`true || true && false`, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(false, nil), nil}, nil}},
		{`true && true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(true, nil), nil}, NewLiteral(false, nil), nil}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexBoolParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_intConditions(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 && 1", &BinaryCondition{AndCondition{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}},
		{"1 && 0", &BinaryCondition{AndCondition{}, NewLiteral(1, nil), NewLiteral(0, nil), nil}},
		{"0 && 0", &BinaryCondition{AndCondition{}, NewLiteral(0, nil), NewLiteral(0, nil), nil}},
		{"0 || 0", &BinaryCondition{OrCondition{}, NewLiteral(0, nil), NewLiteral(0, nil), nil}},
		{"0 || 1", &BinaryCondition{OrCondition{}, NewLiteral(0, nil), NewLiteral(1, nil), nil}},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("InvalidParsers%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(expr, tt.expectedValue, t)
		})
	}
}

func Test_InvalidParser(t *testing.T) {
	var tests = []string{
		"(1+1",
		"1+",
		"+1",
		"1+1)",
		"1+1)+1",
	}
	for index, tt := range tests {
		testName := fmt.Sprintf("InvalidParsers%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, _ := NewParser(tt)
			_, exprErr := parser.ParseExpr()
			if exprErr == nil {
				t.Fatalf("Expected parsing error")
			}
		})
	}
}

func Test_parseFunctionCall(t *testing.T) {
	fooFunc := FunctionSignature{&FunctionIdentifier{"foo", 0}, []*Parameter{}, &IntType{}}
	fooFunc1Param := FunctionSignature{&FunctionIdentifier{"foo", 1}, []*Parameter{{"param1", &IntType{}}}, &IntType{}}
	var tests = []struct {
		input            string
		expectedExpr     Expression
		varSymbols       *VarSymbolTable
		funcParseSymbols *FuncParserSymbolTable
		funcSymbols      *FuncSymbolTable
	}{
		{
			"foo()",
			&CallExpr{&FunctionExpr{fooFunc.FunctionIdentifier, &fooFunc, 0}, []Expression{}, -1},
			&VarSymbolTable{map[string]Type{}, nil},
			&FuncParserSymbolTable{map[string]*FunctionIdentifier{"foo": fooFunc.FunctionIdentifier}, nil},
			&FuncSymbolTable{map[FunctionIdentifier]*FunctionSignature{*fooFunc.FunctionIdentifier: &fooFunc}, nil},
		}, {
			"foo(1)",
			&CallExpr{&FunctionExpr{fooFunc1Param.FunctionIdentifier, &fooFunc1Param, 0}, []Expression{NewLiteral(1, nil)}, -1},
			&VarSymbolTable{map[string]Type{}, nil},
			&FuncParserSymbolTable{map[string]*FunctionIdentifier{"foo": fooFunc1Param.FunctionIdentifier}, nil},
			&FuncSymbolTable{map[FunctionIdentifier]*FunctionSignature{*fooFunc1Param.FunctionIdentifier: &fooFunc1Param}, nil},
		},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexFunctionCall%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatalf("parser error: %v", parserErr)
			}
			parser.VarSymbols = tt.varSymbols
			parser.FuncSymbols = tt.funcParseSymbols
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}
			NewTypeChecker(&[]Statement{&statements.ExpressionStatement{expr, 0}}, tt.varSymbols, tt.funcSymbols).CheckTypes()

			CompareExpressions(expr, tt.expectedExpr, t)
		})
	}
}
