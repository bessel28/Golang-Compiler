package test

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"testing"
)

func EvaluateTest(t *testing.T, expr Expression, env *Environment, expectedLiteral Literal) {
	literal, err := Eval(expr, env)
	if err != nil {
		t.Errorf("Error evaluating literal: %v", err)
	}

	if !expectedLiteral.Equals(literal) {
		t.Errorf("Incorrect literal. Expected: %v, got: %v", expectedLiteral, literal)
	}
}
