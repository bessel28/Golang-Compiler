package parsing

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/parser"
	"fmt"
	"reflect"
	"testing"
)

func Test_Parser(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1+1", &BinaryExpr{Left: NewLiteral(1), Right: NewLiteral(1), Op: AddOperator{}}},
		{"1+1*2", &BinaryExpr{Left: NewLiteral(1), Right: &BinaryExpr{Left: NewLiteral(1), Right: NewLiteral(2), Op: MultOperator{}}, Op: AddOperator{}}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserIntValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1+1", &BinaryExpr{AddOperator{}, NewLiteral(1), NewLiteral(1)}},
		{"1+1*2", &BinaryExpr{AddOperator{}, NewLiteral(1), &BinaryExpr{MultOperator{}, NewLiteral(1), NewLiteral(2)}}},
		{"2*2+1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2), NewLiteral(2)}, NewLiteral(1)}},
		{"(1+1)*2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1), NewLiteral(1)}, NewLiteral(2)}},
		{"5>4?1:2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5), NewLiteral(4)}, NewLiteral(1), NewLiteral(2)}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserIntSpacesValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 + 1", &BinaryExpr{AddOperator{}, NewLiteral(1), NewLiteral(1)}},
		{"1 + 1 * 2", &BinaryExpr{AddOperator{}, NewLiteral(1), &BinaryExpr{MultOperator{}, NewLiteral(1), NewLiteral(2)}}},
		{"2 * 2 + 1", &BinaryExpr{AddOperator{}, &BinaryExpr{MultOperator{}, NewLiteral(2), NewLiteral(2)}, NewLiteral(1)}},
		{"( 1 + 1 ) * 2", &BinaryExpr{MultOperator{}, &BinaryExpr{AddOperator{}, NewLiteral(1), NewLiteral(1)}, NewLiteral(2)}},
		{"5 > 4 ? 1 : 2", &TernaryExpr{IfElseOperator{}, &BinaryCondition{GreaterThanCondition{}, NewLiteral(5), NewLiteral(4)}, NewLiteral(1), NewLiteral(2)}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserStringValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`"a"+"b"`, &BinaryExpr{AddOperator{}, NewLiteral("a"), NewLiteral("b")}},
		{`"ab"+"c"`, &BinaryExpr{AddOperator{}, NewLiteral("ab"), NewLiteral("c")}},
		{`"ab"+""`, &BinaryExpr{AddOperator{}, NewLiteral("ab"), NewLiteral("")}},
		{`""+"c"`, &BinaryExpr{AddOperator{}, NewLiteral(""), NewLiteral("c")}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_ParserBoolValue(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{`true || false`, &BinaryCondition{OrCondition{}, NewLiteral(true), NewLiteral(false)}},
		{`true && false`, &BinaryCondition{AndCondition{}, NewLiteral(true), NewLiteral(false)}},
		{`true && true`, &BinaryCondition{AndCondition{}, NewLiteral(true), NewLiteral(true)}},
		{`(true || false) && true`, &BinaryCondition{AndCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true), NewLiteral(false)}, NewLiteral(true)}},
		{`true && (true || false)`, &BinaryCondition{AndCondition{}, NewLiteral(true), &BinaryCondition{OrCondition{}, NewLiteral(true), NewLiteral(false)}}},
		{`true || true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{OrCondition{}, NewLiteral(true), NewLiteral(true)}, NewLiteral(false)}},
		{`true || true && false`, &BinaryCondition{OrCondition{}, NewLiteral(true), &BinaryCondition{AndCondition{}, NewLiteral(true), NewLiteral(false)}}},
		{`true && true || false`, &BinaryCondition{OrCondition{}, &BinaryCondition{AndCondition{}, NewLiteral(true), NewLiteral(true)}, NewLiteral(false)}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
		})
	}
}

func Test_intConditions(t *testing.T) {
	var tests = []struct {
		input         string
		expectedValue Expression
	}{
		{"1 && 1", &BinaryCondition{AndCondition{}, NewLiteral(1), NewLiteral(1)}},
		{"1 && 0", &BinaryCondition{AndCondition{}, NewLiteral(1), NewLiteral(0)}},
		{"0 && 0", &BinaryCondition{AndCondition{}, NewLiteral(0), NewLiteral(0)}},
		{"0 || 0", &BinaryCondition{OrCondition{}, NewLiteral(0), NewLiteral(0)}},
		{"0 || 1", &BinaryCondition{OrCondition{}, NewLiteral(0), NewLiteral(1)}},
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

			compareExpressions(&expr, &tt.expectedValue, t)
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

func compareExpressions(actualExpr *Expression, expectedExpr *Expression, t *testing.T) {
	if reflect.TypeOf(actualExpr) != reflect.TypeOf(expectedExpr) {
		t.Fatalf("(sub)Expression type was not expected. Expected: %T, Actual: %T", expectedExpr, actualExpr)
	}

	switch (*actualExpr).(type) {
	case Literal:
		actualLit := (*actualExpr).(Literal)
		expectedLit := (*expectedExpr).(Literal)
		if !actualLit.Equals(expectedLit) {
			t.Fatalf("(sub)Literal was unexpected value. Expected: %v, Actual: %v", expectedLit.String(), actualLit.String())
		}
		return
	case *UnaryExpr:
		actualUnary := (*actualExpr).(*UnaryExpr)
		expectedUnary := (*expectedExpr).(*UnaryExpr)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		compareExpressions(&actualUnary.Arg, &expectedUnary.Arg, t)
		return
	case *BinaryExpr:
		actualBinary := (*actualExpr).(*BinaryExpr)
		expectedBinary := (*expectedExpr).(*BinaryExpr)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		compareExpressions(&actualBinary.Left, &expectedBinary.Left, t)
		compareExpressions(&actualBinary.Right, &expectedBinary.Right, t)
		return
	case *TernaryExpr:
		actualTernary := (*actualExpr).(*TernaryExpr)
		expectedTernary := (*expectedExpr).(*TernaryExpr)
		switch {
		case reflect.TypeOf(actualTernary.Op) != reflect.TypeOf(expectedTernary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedTernary.Op, actualTernary.Op)
		case reflect.TypeOf(actualTernary.Condition) != reflect.TypeOf(expectedTernary.Condition):
			t.Fatalf("Unary Condition did not match. Expected: %T, Actual: %T", expectedTernary.Condition, actualTernary.Condition)
		case reflect.TypeOf(actualTernary.Left) != reflect.TypeOf(expectedTernary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedTernary.Left, actualTernary.Left)
		case reflect.TypeOf(actualTernary.Right) != reflect.TypeOf(expectedTernary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedTernary.Right, actualTernary.Right)
		}
		actualCondExpr := actualTernary.Condition.(Expression)
		expectedCondExpr := expectedTernary.Condition.(Expression)
		compareExpressions(&actualCondExpr, &expectedCondExpr, t)
		compareExpressions(&actualTernary.Left, &expectedTernary.Left, t)
		compareExpressions(&actualTernary.Right, &expectedTernary.Right, t)
		return
	case *UnaryCondition:
		actualUnary := (*actualExpr).(*UnaryCondition)
		expectedUnary := (*expectedExpr).(*UnaryCondition)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		compareExpressions(&actualUnary.Arg, &expectedUnary.Arg, t)
		return
	case *BinaryCondition:
		actualBinary := (*actualExpr).(*BinaryCondition)
		expectedBinary := (*expectedExpr).(*BinaryCondition)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		compareExpressions(&actualBinary.Left, &expectedBinary.Left, t)
		compareExpressions(&actualBinary.Right, &expectedBinary.Right, t)
		return
	}
}
