package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"fmt"
	"testing"
)

func TestNotEqualCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1), NewLiteral(1), false},
		{NewLiteral(1), NewLiteral(2), true},
		{NewLiteral(1), NewLiteral("1"), true},
		{NewLiteral("1"), NewLiteral("1"), false},
		{NewLiteral(true), NewLiteral(true), false},
		{NewLiteral(true), NewLiteral(false), true},
		{NewLiteral(true), NewLiteral(1), true},
	}

	env := Environment{}

	for index, tt := range tests {
		testName := fmt.Sprintf("NotEqualCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := BinaryCondition{NotEqualCondition{}, tt.val1, tt.val2}

			conditionCheck, err := Eval(&condition, &env)
			if err != nil {
				t.Error(err)
			}

			boolCondition, ok := conditionCheck.(*BoolLiteral)
			if !ok {
				t.Errorf("Condition check returned non-BoolLiteral %T", conditionCheck)
			}

			if boolCondition.Value != tt.res {
				t.Errorf("Condition Expected %v but check returned %v", tt.res, boolCondition.Value)
			}
		})
	}

}
