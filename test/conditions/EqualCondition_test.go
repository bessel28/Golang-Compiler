package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"fmt"
	"testing"
)

func TestEqualCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1, nil), NewLiteral(1, nil), true},
		{NewLiteral(1, nil), NewLiteral(2, nil), false},
		{NewLiteral(1, nil), NewLiteral("1", nil), false},
		{NewLiteral("1", nil), NewLiteral("1", nil), true},
		{NewLiteral(true, nil), NewLiteral(true, nil), true},
		{NewLiteral(true, nil), NewLiteral(false, nil), false},
		{NewLiteral(true, nil), NewLiteral(1, nil), false},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("EqualCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), &Executor{})

			condition := BinaryCondition{EqualCondition{}, tt.val1, tt.val2, nil}

			conditionCheck, err := evaluator.Eval(&condition)
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
