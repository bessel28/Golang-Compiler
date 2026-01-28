package test

import (
	. "compiler/components"
	. "compiler/components/operators"
	"fmt"
	"testing"
)

func TestMultOperator(t *testing.T) {
	var tests = []struct {
		a, b int
		res  int
	}{
		{0, 1, 0},
		{-5, -3, 15},
		{2, -2, -4},
		{2, 5, 10},
		{5, 2, 10},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("MultOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    MultOperator{},
			}

			expectedAnswer, _ := (&IntLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestStringMultOperator(t *testing.T) {
	var tests = []struct {
		a, b Literal
	}{
		{*NewLiteral(1), *NewLiteral("test")},
		{*NewLiteral("test"), *NewLiteral(1)},
		{*NewLiteral("test"), *NewLiteral("test")},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringMultOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryOperator{
				Left:  &tt.a,
				Right: &tt.b,
				Op:    MultOperator{},
			}

			_, err := binOperator.Eval()
			if err == nil {
				t.Errorf("Expected string multiplication to return an error")
			}
		})

	}
}
