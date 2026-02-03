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
		{&BinaryCondition{EqualCondition{}, NewLiteral(1), NewLiteral(1)}, NewLiteral(0), NewLiteral(1), 0},
		{&BinaryCondition{EqualCondition{}, NewLiteral(0), NewLiteral(1)}, NewLiteral(0), NewLiteral(1), 1},
		{&BinaryCondition{EqualCondition{}, NewLiteral(0), NewLiteral(1)},
			&BinaryExpr{AddOperator{}, NewLiteral(1), NewLiteral(2)},
			&BinaryExpr{AddOperator{}, NewLiteral(2), NewLiteral(2)}, 4},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("TernaryExpr%d", index)
		t.Run(testName, func(t *testing.T) {
			ternary := TernaryExpr{
				IfElseOperator{},
				tt.condition,
				tt.trueVal,
				tt.falseVal,
			}

			expectedAnswer := &IntLiteral{Value: tt.res}
			EvaluateTest(t, &ternary, &Environment{}, expectedAnswer)
		})
	}
}
