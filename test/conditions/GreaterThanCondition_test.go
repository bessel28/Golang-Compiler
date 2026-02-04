package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/conditions"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"fmt"
	"testing"
)

func TestGreaterThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
		res        bool
	}{
		{NewLiteral(1, nil), NewLiteral(1, nil), false},
		{NewLiteral(1, nil), NewLiteral(2, nil), false},
		{NewLiteral(2, nil), NewLiteral(1, nil), true},
		{NewLiteral("A", nil), NewLiteral("A", nil), false},
		{NewLiteral("B", nil), NewLiteral("A", nil), true},
		{NewLiteral("A", nil), NewLiteral("B", nil), false},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("GreaterThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			evaluator := NewEvaluator(&RuntimeEnvironment{}, &Executor{})

			condition := BinaryCondition{GreaterThanCondition{}, tt.val1, tt.val2, nil}

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

func TestInvalidGreaterThanCondition(t *testing.T) {
	var tests = []struct {
		val1, val2 Literal
	}{

		{NewLiteral(1, nil), NewLiteral("1", nil)},
		{NewLiteral(true, nil), NewLiteral(true, nil)},
		{NewLiteral(true, nil), NewLiteral(false, nil)},
		{NewLiteral(true, nil), NewLiteral(1, nil)},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("InvalidGreaterThanCondition%d", index)
		t.Run(testName, func(t *testing.T) {
			evaluator := NewEvaluator(&RuntimeEnvironment{}, &Executor{})

			condition := BinaryCondition{GreaterThanCondition{}, tt.val1, tt.val2, nil}

			_, err := evaluator.Eval(&condition)

			if err == nil {
				t.Error("Expected error but check returned nil")
			}
		})
	}
}
