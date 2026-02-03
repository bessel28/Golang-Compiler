package operators

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/runtime"
	. "compiler/test"
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
			binOperator1 := BinaryExpr{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    MultOperator{},
			}

			expectedAnswer := &IntLiteral{Value: tt.res}
			EvaluateTest(t, &binOperator1, &Environment{}, expectedAnswer)
		})
	}
}

func TestStringMultOperator(t *testing.T) {
	var tests = []struct {
		a, b interface{}
	}{
		{1, "test"},
		{"test", 1},
		{"test", "test"},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("StringMultOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			binOperator := BinaryExpr{
				Left:  NewLiteral(tt.a),
				Right: NewLiteral(tt.b),
				Op:    MultOperator{},
			}

			_, err := Eval(&binOperator, &Environment{})
			if err == nil {
				t.Errorf("Expected string multiplication to return an error")
			}
		})

	}
}
