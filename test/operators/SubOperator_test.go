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
			binOperator1 := BinaryExpr{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    SubOperator{},
			}

			expectedAnswer := &IntLiteral{Value: tt.res}
			EvaluateTest(t, &binOperator1, &Environment{}, expectedAnswer)
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
			binOperator := BinaryExpr{
				Left:  NewLiteral(tt.a),
				Right: NewLiteral(tt.b),
				Op:    SubOperator{},
			}

			_, err := Eval(&binOperator, &Environment{})
			if err == nil {
				t.Errorf("Expected string Subiplication to return an error")
			}
		})

	}
}
