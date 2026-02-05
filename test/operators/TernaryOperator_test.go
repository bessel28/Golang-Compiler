package operators

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/runtime"
	. "compiler/test"
	"fmt"
	"testing"
)

func TestTernaryOperator(t *testing.T) {
	var tests = []struct {
		condition         Condition
		trueVal, falseVal Expression
		res               int
	}{
		{&BinaryCondition{EqualCondition{}, NewLiteral(1, nil), NewLiteral(1, nil), nil}, NewLiteral(0, nil), NewLiteral(1, nil), 0},
		{&BinaryCondition{EqualCondition{}, NewLiteral(0, nil), NewLiteral(1, nil), nil}, NewLiteral(0, nil), NewLiteral(1, nil), 1},
		{&BinaryCondition{EqualCondition{}, NewLiteral(0, nil), NewLiteral(1, nil), nil},
			&BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(2, nil), nil},
			&BinaryExpr{AddOperator{}, NewLiteral(2, nil), NewLiteral(2, nil), nil}, 4},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("TernaryExpr%d", index)
		t.Run(testName, func(t *testing.T) {
			ternary := TernaryExpr{
				IfElseOperator{},
				tt.condition,
				tt.trueVal,
				tt.falseVal,
				nil,
			}

			expectedAnswer := &IntLiteral{Value: tt.res}
			evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), &Executor{})
			EvaluateTest(t, &ternary, evaluator, expectedAnswer)
		})
	}
}
