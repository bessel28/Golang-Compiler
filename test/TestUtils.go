package test

import (
	. "compiler/src/components/ast"
	"compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/runtime"
	"reflect"
	"testing"
)

func EvaluateTest(t *testing.T, expr Expression, env *RuntimeEnvironment, expectedLiteral Literal) {
	literal, err := Eval(expr, env)
	if err != nil {
		t.Errorf("Error evaluating literal: %v", err)
	}

	if !expectedLiteral.Equals(literal) {
		t.Errorf("Incorrect literal. Expected: %v, got: %v", expectedLiteral, literal)
	}
}

func CompareExpressions(actualExpr *Expression, expectedExpr *Expression, t *testing.T) {
	if reflect.TypeOf(actualExpr) != reflect.TypeOf(expectedExpr) {
		t.Fatalf("(sub)Expression type was not expected. Expected: %T, Actual: %T", expectedExpr, actualExpr)
	}

	switch (*actualExpr).(type) {
	case Literal:
		actualLit := (*actualExpr).(Literal)
		expectedLit := (*expectedExpr).(Literal)
		if !actualLit.Equals(expectedLit) {
			t.Fatalf("(sub)Literal was unexpected value. Expected: %v, Actual: %v", expectedLit.String(), actualLit.String())
		}
	case *UnaryExpr:
		actualUnary := (*actualExpr).(*UnaryExpr)
		expectedUnary := (*expectedExpr).(*UnaryExpr)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		CompareExpressions(&actualUnary.Arg, &expectedUnary.Arg, t)
	case *BinaryExpr:
		actualBinary := (*actualExpr).(*BinaryExpr)
		expectedBinary := (*expectedExpr).(*BinaryExpr)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		CompareExpressions(&actualBinary.Left, &expectedBinary.Left, t)
		CompareExpressions(&actualBinary.Right, &expectedBinary.Right, t)
	case *TernaryExpr:
		actualTernary := (*actualExpr).(*TernaryExpr)
		expectedTernary := (*expectedExpr).(*TernaryExpr)
		switch {
		case reflect.TypeOf(actualTernary.Op) != reflect.TypeOf(expectedTernary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedTernary.Op, actualTernary.Op)
		case reflect.TypeOf(actualTernary.Condition) != reflect.TypeOf(expectedTernary.Condition):
			t.Fatalf("Unary Condition did not match. Expected: %T, Actual: %T", expectedTernary.Condition, actualTernary.Condition)
		case reflect.TypeOf(actualTernary.Left) != reflect.TypeOf(expectedTernary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedTernary.Left, actualTernary.Left)
		case reflect.TypeOf(actualTernary.Right) != reflect.TypeOf(expectedTernary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedTernary.Right, actualTernary.Right)
		}
		actualCondExpr := actualTernary.Condition.(Expression)
		expectedCondExpr := expectedTernary.Condition.(Expression)
		CompareExpressions(&actualCondExpr, &expectedCondExpr, t)
		CompareExpressions(&actualTernary.Left, &expectedTernary.Left, t)
		CompareExpressions(&actualTernary.Right, &expectedTernary.Right, t)
	case *UnaryCondition:
		actualUnary := (*actualExpr).(*UnaryCondition)
		expectedUnary := (*expectedExpr).(*UnaryCondition)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		CompareExpressions(&actualUnary.Arg, &expectedUnary.Arg, t)
	case *BinaryCondition:
		actualBinary := (*actualExpr).(*BinaryCondition)
		expectedBinary := (*expectedExpr).(*BinaryCondition)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		CompareExpressions(&actualBinary.Left, &expectedBinary.Left, t)
		CompareExpressions(&actualBinary.Right, &expectedBinary.Right, t)
	case *functions.CallExpr:
		actualFunction := (*actualExpr).(*functions.CallExpr)
		expectedFunction := (*expectedExpr).(*functions.CallExpr)
		switch {
		case actualFunction.FunctionSignature.Name != expectedFunction.FunctionSignature.Name:
			t.Fatalf("Function Signature names did not match. Expected %v, Actual %v", actualFunction.FunctionSignature.Name, expectedFunction.FunctionSignature.Name)
		case actualFunction.FunctionSignature.ReturnType != expectedFunction.FunctionSignature.ReturnType:
			t.Fatalf("Function Signature return types did not match. Expected %v, Actual %v", actualFunction.FunctionSignature.ReturnType, expectedFunction.FunctionSignature.ReturnType)
		case len(actualFunction.FunctionSignature.Params) != len(expectedFunction.FunctionSignature.Params):
			t.Fatalf("Incorrect number of params in function signature. Expected %v, Actual %v", len(expectedFunction.Params), len(actualFunction.Params))

		}
		for i, expectedParam := range expectedFunction.Params {
			if actualFunction.FunctionSignature.Params[i] != expectedParam {
				t.Fatalf("Function Signature param types don't match at %d. Expected %v, Actual %v", i, expectedParam, actualFunction.FunctionSignature.Params[i])
			}
		}
	default:
		t.Skipf("Undefined Expression Comparison for type %T", *actualExpr)
	}
}
