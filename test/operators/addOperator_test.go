package operators

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	"fmt"
	"testing"
)

func TestIntAddOperator(t *testing.T) {
	var tests = []struct {
		a, b int
		res  int
	}{
		{0, 1, 1},
		{-5, -3, -8},
		{2, -2, 0},
		{2, 5, 7},
		{5, 2, 7},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("IntAddOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    AddOperator{},
			}

			expectedAnswer, _ := (&IntLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestNestedIntAddOperator(t *testing.T) {
	var tests = []struct {
		a, b, c int
		res     int
	}{
		{1, 1, 5, 7},
		{5, 5, 5, 15},
		{-5, -5, -5, -15},
		{-5, 5, -5, -5},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("NestedIntAddOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryOperator{
				Left: &IntLiteral{Value: tt.a},
				Right: &BinaryOperator{
					Left:  &IntLiteral{Value: tt.b},
					Right: &IntLiteral{Value: tt.c},
					Op:    AddOperator{},
				},
				Op: AddOperator{},
			}

			expectedAnswer, _ := (&IntLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator, expectedAnswer)
		})
	}
}

func TestStringAddOperator(t *testing.T) {
	var tests = []struct {
		a, b string
		res  string
	}{
		{"test", "text", "testtext"},
		{"text", "test", "texttest"},
		{"a", "b", "ab"},
		{"b", "", "b"},
		{"", "b", "b"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringAddOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &StringLiteral{Value: tt.a},
				Right: &StringLiteral{Value: tt.b},
				Op:    AddOperator{},
			}

			expectedAnswer, _ := (&StringLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestStringIntAddOperator(t *testing.T) {
	var tests = []struct {
		a   string
		b   int
		res string
	}{
		{"test", 1, "test1"},
		{"text", 2, "text2"},
		{"text", -2, "text-2"},
		{"", 1, "1"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringIntAddOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &StringLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    AddOperator{},
			}

			expectedAnswer, _ := (&StringLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestIntStringAddOperator(t *testing.T) {
	var tests = []struct {
		a   int
		b   string
		res string
	}{
		{1, "test", "1test"},
		{2, "text", "2text"},
		{-2, "text", "-2text"},
		{1, "", "1"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("IntStringAddOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &StringLiteral{Value: tt.b},
				Op:    AddOperator{},
			}

			expectedAnswer, _ := (&StringLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func EvaluateTest(t *testing.T, expr Expression, expectedLiteral Literal) {
	literal, err := expr.Eval()
	if err != nil {
		t.Errorf("Error evaluating literal: %v", err)
	}

	if !expectedLiteral.Equals(literal) {
		t.Errorf("Incorrect literal. Expected: %v, got: %v", expectedLiteral, literal)
	}
}
