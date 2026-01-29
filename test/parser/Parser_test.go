package parser

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
	. "compiler/components/operators"
	parser2 "compiler/parser"
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
			parser := parser2.NewParser(tt.input)
			expr := parser.ParseExpr()

			fmt.Printf("Expected: %+v\n", tt.expectedExpression)
			fmt.Printf("Actual: %+v\n", expr)

			if reflect.TypeOf(expr) != reflect.TypeOf(tt.expectedExpression) {
				t.Errorf("Incorrect Expression Type. Expected %T got %T", tt.expectedExpression, expr)
			}

			exprVal, _ := expr.Eval()
			expectedVal, _ := tt.expectedExpression.Eval()

			if !exprVal.Equals(expectedVal) {
				t.Errorf("Incorrect eval. Expected %v got %v", expectedVal, exprVal)
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
		{"(1+1)*2)", 4},
		{"5>4?1:2", 1},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("LexParser%d", index)
		t.Run(testName, func(t *testing.T) {
			parser := parser2.NewParser(tt.input)
			expr := parser.ParseExpr()

			exprVal, _ := expr.Eval()

			if !exprVal.Equals(NewLiteral(tt.value)) {
				t.Errorf("Incorrect eval. Expected %v got %v", tt.value, exprVal)
			}

		})
	}
}

//TODO test logical operators and string & bool parsing
