package parsing

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	"compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/parser"
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

			CompareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserIntValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1+1", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil)}},
		{"1+1*2", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), &BinaryExpr{MultOperator{}, NewLiteral(1, nil), NewLiteral(2, nil)}}},
		{"2*2+1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2, nil), NewLiteral(2, nil)}, NewLiteral(1, nil)}},
		{"(1+1)*2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil)}, NewLiteral(2, nil)}},
		{"5>4?1:2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5, nil), NewLiteral(4, nil)}, NewLiteral(1, nil), NewLiteral(2, nil)}},
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

			CompareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserIntSpacesValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 + 1", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil)}},
		{"1 + 1 * 2", &BinaryExpr{AddOperator{}, NewLiteral(1, nil), &BinaryExpr{MultOperator{}, NewLiteral(1, nil), NewLiteral(2, nil)}}},
		{"2 * 2 + 1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2, nil), NewLiteral(2, nil)}, NewLiteral(1, nil)}},
		{"( 1 + 1 ) * 2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil)}, NewLiteral(2, nil)}},
		{"5 > 4 ? 1 : 2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5, nil), NewLiteral(4, nil)}, NewLiteral(1, nil), NewLiteral(2, nil)}},
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

			CompareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserStringValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`"a"+"b"`, &BinaryExpr{AddOperator{}, NewLiteral("a", nil), NewLiteral("b", nil)}},
		{`"ab"+"c"`, &BinaryExpr{AddOperator{}, NewLiteral("ab", nil), NewLiteral("c", nil)}},
		{`"ab"+""`, &BinaryExpr{AddOperator{}, NewLiteral("ab", nil), NewLiteral("", nil)}},
		{`""+"c"`, &BinaryExpr{AddOperator{}, NewLiteral("", nil), NewLiteral("c", nil)}},
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

			CompareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserBoolValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`true || false`, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil)}},
		{`true && false`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(false, nil)}},
		{`true && true`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(true, nil)}},
		{`(true || false) && true`, &BinaryCondition{AndCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil)}, NewLiteral(true, nil)}},
		{`true && (true || false)`, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(false, nil)}}},
		{`true || true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), NewLiteral(true, nil)}, NewLiteral(false, nil)}},
		{`true || true && false`, &BinaryCondition{OrCondition{}, NewLiteral(true, nil), &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(false, nil)}}},
		{`true && true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{AndCondition{}, NewLiteral(true, nil), NewLiteral(true, nil)}, NewLiteral(false, nil)}},
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

			CompareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_intConditions(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 && 1", &BinaryCondition{AndCondition{}, NewLiteral(1, nil), NewLiteral(1, nil)}},
		{"1 && 0", &BinaryCondition{AndCondition{}, NewLiteral(1, nil), NewLiteral(0, nil)}},
		{"0 && 0", &BinaryCondition{AndCondition{}, NewLiteral(0, nil), NewLiteral(0, nil)}},
		{"0 || 0", &BinaryCondition{OrCondition{}, NewLiteral(0, nil), NewLiteral(0, nil)}},
		{"0 || 1", &BinaryCondition{OrCondition{}, NewLiteral(0, nil), NewLiteral(1, nil)}},
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

			CompareExpressions(&expr, &tt.expectedValue, t)
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
	var intKind = IntKind
	fooFunc := functions.FunctionSignature{"foo", []ValueKind{}, &intKind}
	fooFunc1Param := functions.FunctionSignature{"foo", []ValueKind{IntKind}, &intKind}
	var tests = []struct {
		input        string
		expectedExpr Expression
		varSymbols   *VarSymbolTable
		funcSymbols  *FuncSymbolTable
	}{
		{
			"foo()",
			&functions.CallExpr{nil, -1, &fooFunc},
			&VarSymbolTable{map[string]*ValueKind{}, nil},
			&FuncSymbolTable{map[string]*functions.FunctionSignature{"foo": &fooFunc}, nil},
		}, {
			"foo(1)",
			&functions.CallExpr{nil, -1, &fooFunc1Param},
			&VarSymbolTable{map[string]*ValueKind{}, nil},
			&FuncSymbolTable{map[string]*functions.FunctionSignature{"foo": &fooFunc1Param}, nil},
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
			parser.FuncSymbols = tt.funcSymbols
			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatalf("expression error: %v", exprErr)
			}

			CompareExpressions(&expr, &tt.expectedExpr, t)
		})
	}
}
