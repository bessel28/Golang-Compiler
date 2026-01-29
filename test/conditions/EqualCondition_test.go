package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/conditions"
	. "compiler/components/literals"
	"fmt"
	"testing"
)

func TestEqualCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1), NewLiteral(1), true},
		{NewLiteral(1), NewLiteral(2), false},
		{NewLiteral(1), NewLiteral("1"), false},
		{NewLiteral("1"), NewLiteral("1"), true},
		{NewLiteral(true), NewLiteral(true), true},
		{NewLiteral(true), NewLiteral(false), false},
		{NewLiteral(true), NewLiteral(1), false},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("EqualCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := EqualCondition{&BinaryCondition{Left: tt.val1, Right: tt.val2}}

			conditionCheck, err := condition.Eval()
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
