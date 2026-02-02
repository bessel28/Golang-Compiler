package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	"fmt"
	"testing"
)

func TestGreaterThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1), NewLiteral(1), false},
		{NewLiteral(1), NewLiteral(2), false},
		{NewLiteral(2), NewLiteral(1), true},
		{NewLiteral("A"), NewLiteral("A"), false},
		{NewLiteral("B"), NewLiteral("A"), true},
		{NewLiteral("A"), NewLiteral("B"), false},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("GreaterThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := GreaterThanCondition{&BinaryCondition{Left: tt.val1, Right: tt.val2}}

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

func TestInvalidGreaterThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
	}{

		{NewLiteral(1), NewLiteral("1")},
		{NewLiteral(true), NewLiteral(true)},
		{NewLiteral(true), NewLiteral(false)},
		{NewLiteral(true), NewLiteral(1)},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("InvalidGreaterThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := GreaterThanCondition{&BinaryCondition{Left: tt.val1, Right: tt.val2}}

			_, err := condition.Eval()
			if err == nil {
				t.Error("Expected error but check returned nil")
			}
		})
	}
}
