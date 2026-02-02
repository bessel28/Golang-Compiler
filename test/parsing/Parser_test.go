package parsing

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/parser"
	"fmt"
	"reflect"
	"testing"
)

func Test_Parser(t *testing.T) {
	var tests = []struct {
		input              string
		expectedExpression Expression
	}{
		{"1+1", &BinaryOperator{Left: NewLiteral(1), Right: NewLiteral(1), Op: AddOperator{}}},
		{"1+1*2", &BinaryOperator{Left: NewLiteral(1), Right: &BinaryOperator{Left: NewLiteral(1), Right: NewLiteral(2), Op: MultOperator{}}, Op: AddOperator{}}},
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

			if reflect.TypeOf(expr) != reflect.TypeOf(tt.expectedExpression) {
				t.Fatalf("Incorrect Expression Type. Expected %T got %T", tt.expectedExpression, expr)
			}

			exprVal, _ := expr.Eval()
			expectedVal, _ := tt.expectedExpression.Eval()

			if !exprVal.Equals(expectedVal) {
				t.Fatalf("Incorrect eval. Expected %v got %v", expectedVal, exprVal)
			}

		})
	}
}

func Test_ParserIntValue(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"1+1", 2},
		{"1+1*2", 3},
		{"2*2+1", 5},
		{"(1+1)*2", 4},
		{"5>4?1:2", 1},
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

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.value)) {
				t.Fatalf("Incorrect eval. Expected %v got %v", tt.value, exprVal)
			}

		})
	}
}

func Test_ParserIntSpacesValue(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"1 + 1", 2},
		{"1 + 1 * 2", 3},
		{"2 * 2 + 1", 5},
		{"( 1 + 1 ) * 2", 4},
		{"5 > 4 ? 1 : 2", 1},
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

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.value)) {
				t.Fatalf("Incorrect eval. Expected %v got %v", tt.value, exprVal)
			}

		})
	}
}

func Test_ParserStringValue(t *testing.T) {
	var tests = []struct {
		input string
		value string
	}{
		{`"a"+"b"`, "ab"},
		{`"ab"+"c"`, "abc"},
		{`"ab"+""`, "ab"},
		{`""+"c"`, "c"},
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

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.value)) {
				t.Fatalf("Incorrect eval. Expected %v got %v", tt.value, exprVal)
			}

		})
	}
}

func Test_ParserBoolValue(t *testing.T) {
	var tests = []struct {
		input string
		value bool
	}{
		{`true || false`, true},
		{`true && false`, false},
		{`true && true`, true},
		{`(true || false) && true`, true},
		{`true && (true || false)`, true},
		{`true || true || false`, true},
		{`true || true && false`, true},
		{`true && true || false`, true},
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
				t.Fatalf("expression error: %v", parserErr)
			}

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.value)) {
				t.Fatalf("Incorrect eval. Expected %v got %v", tt.value, exprVal)
			}

		})
	}
}

func Test_intConditions(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue bool
	}{
		{"1 && 1", true},
		{"1 && 0", false},
		{"0 && 0", false},
		{"0 || 0", false},
		{"0 || 1", true},
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
				t.Fatalf("expression error: %v", parserErr)
			}

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.expectedValue)) {
				t.Fatalf("Incorrect eval. Expected %v got %v", tt.expectedValue, exprVal)
			}
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
			expr, exprErr := parser.ParseExpr()
			if exprErr == nil {
				t.Fatalf("Expected parsing error")
			}
			fmt.Println(expr)
		})
	}
}
