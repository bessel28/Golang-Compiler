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
			binOperator := BinaryExpr{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    DivOperator{},
			}

			expectedAnswer := &IntLiteral{Value: tt.res}
			evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), &Executor{})
			EvaluateTest(t, &binOperator, evaluator, expectedAnswer)
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
			binOperator := BinaryExpr{
				Left:  &IntLiteral{Value: tt.a},
				Right: &IntLiteral{Value: tt.b},
				Op:    DivOperator{},
			}

			evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), &Executor{})
			_, err := evaluator.Eval(&binOperator)
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
			binOperator := BinaryExpr{
				Left:  NewLiteral(tt.a, nil),
				Right: NewLiteral(tt.b, nil),
				Op:    DivOperator{},
			}

			evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), &Executor{})
			_, err := evaluator.Eval(&binOperator)
			if err == nil {
				t.Errorf("Expected string division to return an error")
			}
		})

	}
}
