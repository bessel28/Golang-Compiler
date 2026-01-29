package operators

import (
	. "compiler/components/literals"
	. "compiler/components/operators"
	"fmt"
	"testing"
)

func TestSubOperator(t *testing.T) {
	var tests = []struct {
		a, b int
		res  int
	}{
		{0, 1, -1},
		{-5, -3, -2},
		{2, -2, 4},
		{2, 5, -3},
		{5, 2, 3},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("SubOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator1 := BinaryOperator{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    SubOperator{},
			}

			expectedAnswer, _ := (&IntLiteral{Value: tt.res}).Eval()

			EvaluateTest(t, binOperator1, expectedAnswer)
		})
	}
}

func TestStringSubOperator(t *testing.T) {
	var tests = []struct {
		a, b interface{}
	}{
		{1, "test"},
		{"test", 1},
		{"test", "test"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringSubOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryOperator{
				Left:  NewLiteral(tt.a),
				Right: NewLiteral(tt.b),
				Op:    SubOperator{},
			}

			_, err := binOperator.Eval()
			if err == nil {
				t.Errorf("Expected string Subiplication to return an error")
			}
		})

	}
}
