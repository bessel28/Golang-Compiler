package test

import (
	. "compiler/components"
	. "compiler/components/literals"
	. "compiler/components/operators"
	"fmt"
	"testing"
)

func TestDivOperator(t *testing.T) {
	var tests = []struct {
		a, b int
		res  int
	}{
		{0, 1, 0 / 1},
		{-5, -3, -5 / -3},
		{2, -2, -1},
		{2, 5, 2 / 5},
		{5, 2, 5 / 2},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("DivOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    DivOperator{},
			}

			expectedAnswer, _ := (&IntLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestDivByZero(t *testing.T) {
	var tests = []struct {
		a, b int
	}{
		{1, 0},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("DivByZero%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    DivOperator{},
			}

			_, err := binOperator.Eval()
			if err == nil {
				t.Errorf("Expected division by zero to return an error")
			}
		})
	}
}

func TestStringDivOperator(t *testing.T) {
	var tests = []struct {
		a, b interface{}
	}{
		{1, "test"},
		{"test", 1},
		{"test", "test"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringDivOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryOperator{
				Left:  NewLiteral(tt.a),
				Right: NewLiteral(tt.b),
				Op:    DivOperator{},
			}

			_, err := binOperator.Eval()
			if err == nil {
				t.Errorf("Expected string division to return an error")
			}
		})

	}
}
