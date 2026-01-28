package operators

import (
	. "compiler/components"
	. "compiler/components/conditions"
	. "compiler/components/literals"
	. "compiler/components/operators"
	"fmt"
	"testing"
)

func TestTernaryOperator(t *testing.T) {
	var tests = []struct {
		condition         Condition
		trueVal, falseVal Expression
		res               interface{}
	}{
		{&EqualCondition{&BinaryCondition{Left: NewLiteral(1), Right: NewLiteral(1)}}, NewLiteral(0), NewLiteral(1), 0},
		{&EqualCondition{&BinaryCondition{Left: NewLiteral(0), Right: NewLiteral(1)}}, NewLiteral(0), NewLiteral(1), 1},
		{&EqualCondition{&BinaryCondition{Left: NewLiteral(0), Right: NewLiteral(1)}},
			BinaryOperator{Left: NewLiteral(1), Right: NewLiteral(2), Op: AddOperator{}},
			BinaryOperator{Left: NewLiteral(2), Right: NewLiteral(2), Op: AddOperator{}}, 4},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("TernaryOperator%d", index)
		t.Run(testName, func(t *testing.T) {
			ternary := TernaryOperator{
				Condition: tt.condition,
				Left:      tt.trueVal,
				Right:     tt.falseVal,
			}

			EvaluateTest(t, &ternary, NewLiteral(tt.res))
		})
	}
}
