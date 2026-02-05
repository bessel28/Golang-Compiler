package runtime

import (
	. "compiler/src/components/literals"
	. "compiler/src/parser/expression"
	. "compiler/src/runtime"
	"fmt"
	"testing"
)

var executor = &Executor{}

func TestEvalInt(t *testing.T) {
	var tests = []struct {
		input string
		value int
	}{
		{"1+1", 2},
		{"2*2", 4},
		{"(2+3)*2", 10},
		{"5/5", 1},
		{"35/(2+3)", 7},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("TestEvalInt%v", i)
		evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), executor)
		evalTest(testName, &tt.input, tt.value, evaluator, t)
	}
}

func TestEvalBool(t *testing.T) {
	var tests = []struct {
		input string
		value bool
	}{
		{"true||false", true},
		{"true&&false", false},
		{"true&&true", true},
		{"!true&&true", false},
		{"!true||false", false},
		{"!true||!true", false},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("TestEvalBool%v", i)
		evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), executor)
		evalTest(testName, &tt.input, tt.value, evaluator, t)
	}
}

func TestEvalString(t *testing.T) {
	var tests = []struct {
		input string
		value string
	}{
		{`"a"+"b"`, "ab"},
		{`"ab"+"c"`, "abc"},
		{`"a"+"b"+"c"`, "abc"},
		{`"a"+""+"c"`, "ac"},
	}

	for i, tt := range tests {
		evaluator := NewEvaluator(NewRuntimeEnvironment(nil, nil, nil), executor)
		testName := fmt.Sprintf("TestEvalString%v", i)
		evalTest(testName, &tt.input, tt.value, evaluator, t)
	}
}

func evalTest(testName string, input *string, expectedValue any, evaluator *Evaluator, t *testing.T) {

	t.Run(testName, func(t *testing.T) {
		parser, parserErr := NewParser(*input)
		if parserErr != nil {
			t.Fatalf("Parser returned error: %v", parserErr)
		}
		expr, exprErr := parser.ParseExpr()
		if exprErr != nil {
			t.Fatalf("Error parsing expression: %v", exprErr)
		}
		evalVal, evalErr := evaluator.Eval(expr)
		if evalErr != nil {
			t.Fatalf("Error evaluating expression: %v", evalErr)
		}

		if !evalVal.Equals(NewLiteral(expectedValue, nil)) {
			t.Fatalf("Eval returned wrong value. Expected: %v, Actual: %v", expectedValue, evalVal)
		}
	})
}
