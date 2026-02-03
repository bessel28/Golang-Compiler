package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"fmt"
	"testing"
)

func TestLessThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1), NewLiteral(1), false},
		{NewLiteral(1), NewLiteral(2), true},
		{NewLiteral(2), NewLiteral(1), false},
		{NewLiteral("A"), NewLiteral("A"), false},
		{NewLiteral("B"), NewLiteral("A"), false},
		{NewLiteral("A"), NewLiteral("B"), true},
	}

	env := Environment{}

	for index, tt := range tests {
		testName := fmt.Sprintf("LessThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := BinaryCondition{LessThanCondition{}, tt.val1, tt.val2}

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

func TestInvalidLessThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
	}{

		{NewLiteral(1), NewLiteral("1")},
		{NewLiteral(true), NewLiteral(true)},
		{NewLiteral(true), NewLiteral(false)},
		{NewLiteral(true), NewLiteral(1)},
	}

	env := Environment{}

	for index, tt := range tests {
		testName := fmt.Sprintf("InvalidLessThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			condition := BinaryCondition{LessThanCondition{}, tt.val1, tt.val2}

			_, err := Eval(&condition, &env)
			if err == nil {
				t.Error("Expected error but check returned nil")
			}
		})
	}
}
