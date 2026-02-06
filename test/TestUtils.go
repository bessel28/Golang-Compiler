package test

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/statements"
	. "compiler/src/components/variables"
	. "compiler/src/runtime"
	"compiler/src/runtime/value"
	"reflect"
	"testing"
)

func EvaluateTest(t *testing.T, expr Expression, evaluator *Evaluator, expectedValue value.RuntimeValue) {
	literal, err := evaluator.Eval(expr)
	if err != nil {
		t.Errorf("Error evaluating literal: %v", err)
	}

	if expectedValue.Value() != literal.Value() {
		t.Errorf("Incorrect literal. Expected: %v, got: %v", expectedValue, literal)
	}
}

func CompareExpressions(actualExpr Expression, expectedExpr Expression, t *testing.T) {
	if reflect.TypeOf(actualExpr) != reflect.TypeOf(expectedExpr) {
		t.Fatalf("(sub)Expression type was not expected. Expected: %T, Actual: %T", expectedExpr, actualExpr)
	}

	switch actualExpr.(type) {
	case Literal:
		actualLit := actualExpr.(Literal)
		expectedLit := expectedExpr.(Literal)
		if !actualLit.Equals(expectedLit) {
			t.Fatalf("(sub)Literal was unexpected value. Expected: %v, Actual: %v", expectedLit.String(), actualLit.String())
		}
	case *UnaryExpr:
		actualUnary := actualExpr.(*UnaryExpr)
		expectedUnary := expectedExpr.(*UnaryExpr)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		CompareExpressions(actualUnary.Arg, expectedUnary.Arg, t)
	case *BinaryExpr:
		actualBinary := actualExpr.(*BinaryExpr)
		expectedBinary := expectedExpr.(*BinaryExpr)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		CompareExpressions(actualBinary.Left, expectedBinary.Left, t)
		CompareExpressions(actualBinary.Right, expectedBinary.Right, t)
	case *TernaryExpr:
		actualTernary := actualExpr.(*TernaryExpr)
		expectedTernary := expectedExpr.(*TernaryExpr)
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
		CompareExpressions(actualCondExpr, expectedCondExpr, t)
		CompareExpressions(actualTernary.Left, expectedTernary.Left, t)
		CompareExpressions(actualTernary.Right, expectedTernary.Right, t)
	case *UnaryCondition:
		actualUnary := actualExpr.(*UnaryCondition)
		expectedUnary := expectedExpr.(*UnaryCondition)
		switch {
		case reflect.TypeOf(actualUnary.Op) != reflect.TypeOf(expectedUnary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedUnary.Op, actualUnary.Op)
		case reflect.TypeOf(actualUnary.Arg) != reflect.TypeOf(expectedUnary.Arg):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedUnary.Arg, actualUnary.Arg)
		}
		CompareExpressions(actualUnary.Arg, expectedUnary.Arg, t)
	case *BinaryCondition:
		actualBinary := actualExpr.(*BinaryCondition)
		expectedBinary := expectedExpr.(*BinaryCondition)
		switch {
		case reflect.TypeOf(actualBinary.Op) != reflect.TypeOf(expectedBinary.Op):
			t.Fatalf("Unary Operators did not match. Expected: %T, Actual: %T", expectedBinary.Op, actualBinary.Op)
		case reflect.TypeOf(actualBinary.Left) != reflect.TypeOf(expectedBinary.Left):
			t.Fatalf("Unary Arg did not match. Expected: %T, Actual: %T", expectedBinary.Left, actualBinary.Left)
		case reflect.TypeOf(actualBinary.Right) != reflect.TypeOf(expectedBinary.Right):
			t.Fatalf("Unary Right did not match. Expected: %T, Actual: %T", expectedBinary.Right, actualBinary.Right)
		}
		CompareExpressions(actualBinary.Left, expectedBinary.Left, t)
		CompareExpressions(actualBinary.Right, expectedBinary.Right, t)
	case *CallExpr:
		actualFunction := actualExpr.(*CallExpr)
		expectedFunction := expectedExpr.(*CallExpr)

		if actualFunction.Callee.GetFunctionSignature().FunctionIdentifier != expectedFunction.Callee.GetFunctionSignature().FunctionIdentifier {
			t.Fatalf("Function Identifiers did not match. Expected %v, Actual %v",
				expectedFunction.Callee.GetFunctionSignature().FunctionIdentifier, actualFunction.Callee.GetFunctionSignature().FunctionIdentifier)
		}

		if actualFunction.Callee.GetFunctionSignature() == nil {
			if expectedFunction.Callee.GetFunctionSignature() != nil {
				t.Fatalf("Expected function '%v' was not found", expectedFunction.Callee.GetFunctionSignature().FunctionIdentifier)
			}
		} else {
			switch {
			case actualFunction.Callee.GetFunctionSignature().ReturnType != expectedFunction.Callee.GetFunctionSignature().ReturnType:
				t.Fatalf("Function Signature return types did not match. Expected %v, Actual %v",
					actualFunction.Callee.GetFunctionSignature().ReturnType, expectedFunction.Callee.GetFunctionSignature().ReturnType)
			case len(actualFunction.Callee.GetFunctionSignature().Params) != len(expectedFunction.Callee.GetFunctionSignature().Params):
				t.Fatalf("Incorrect number of params in function signature. Expected %v, Actual %v",
					len(expectedFunction.Params), len(actualFunction.Params))

			}
			for i, expectedParam := range expectedFunction.Callee.GetFunctionSignature().Params {
				if actualFunction.Callee.GetFunctionSignature().Params[i] != expectedParam {
					t.Fatalf("Function Signature param types don't match at %d. Expected %v, Actual %v",
						i, expectedParam, actualFunction.Callee.GetFunctionSignature().Params[i])
				}
			}
		}
		for i, expectedParam := range expectedFunction.Params {
			CompareExpressions(actualFunction.Params[i], expectedParam, t)
		}
	case *VariableExpr:
		actualVar := actualExpr.(*VariableExpr)
		expectedVar := expectedExpr.(*VariableExpr)
		if expectedVar.Identifier != actualVar.Identifier {
			t.Fatalf("Variable names did not match. Expected %v, Actual: %v", expectedVar.Identifier, actualVar.Identifier)
		}
		if !expectedVar.ReturnType.Compare(actualVar.ReturnType) {
			t.Fatalf("Variable Return Types did not match. Expected %T, Actual: %T", expectedVar.ReturnType, actualVar.ReturnType)
		}
	default:
		t.Skipf("Undefined Expression Comparison for type %T", actualExpr)
	}
}

func CompareRuntimeEnvironments(actualEnvironment *RuntimeEnvironment, expectedEnvironment *RuntimeEnvironment, t *testing.T) {
	if expectedEnvironment == nil {
		return
	}
	if actualEnvironment == nil {
		t.Fatalf("RuntimeEnvironment is nil")
	}

	if len(expectedEnvironment.Variables) != len(actualEnvironment.Variables) {
		t.Fatalf("Environments have different number of variables declared")
	}
	for varName, exprVar := range expectedEnvironment.Variables {
		actVar, exists := actualEnvironment.Variables[varName]
		if !exists {
			t.Fatalf("Environment variable (%s) not declared in env2", varName)
		}

		switch exprVar.(type) {
		case *value.FuncValue:
			if exprVar.Value().(*value.FunctionIdentifier).Name != actVar.Value().(*value.FunctionIdentifier).Name {
				t.Fatalf("Incorrect variable value for (%s). Expected %v, Actual %v", varName, exprVar, actVar)
			}
			if exprVar.Value().(*value.FunctionIdentifier).UUID != actVar.Value().(*value.FunctionIdentifier).UUID {
				t.Fatalf("Incorrect variable value for (%s). Expected %v, Actual %v", varName, exprVar, actVar)
			}
		default:
			if exprVar.Value() != actVar.Value() {
				t.Fatalf("Incorrect variable value for (%s). Expected %v, Actual %v", varName, exprVar, actVar)
			}
		}

	}
	for funcName, expFunc := range expectedEnvironment.Functions {
		actFunc, exists := actualEnvironment.Functions[funcName]
		if !exists {
			t.Fatalf("Environment function (%s) not declared in env2", funcName)
		}
		if len(actFunc.Params) != len(expFunc.Params) {
			t.Fatalf("Environment function (%s) has different number of parameters declared", funcName)
		}
		for paramIndex, expParam := range expFunc.Params {
			actParam := actFunc.Params[paramIndex]
			if !expParam.Compare(actParam) {
				t.Fatalf("Environment function (%s) parameter %v's type is different than expected. Expected: %T, Actual :%T", funcName, paramIndex, expParam, actParam)
			}
		}
		if !expFunc.ReturnType.Compare(actFunc.ReturnType) {
			t.Fatalf("Environment function (%s) has different return type than expected. Expected %T, Actual %T", funcName, expFunc.ReturnType, actFunc.ReturnType)
		}
		CompareStatements(actFunc.Block, expFunc.Block, t)
	}
	if expectedEnvironment.Parent != nil {
		CompareRuntimeEnvironments(actualEnvironment.Parent, expectedEnvironment.Parent, t)
	}
}

func CompareStatements(actualStatement Statement, expectedStatement Statement, t *testing.T) {
	if reflect.TypeOf(actualStatement) != reflect.TypeOf(expectedStatement) {
		t.Fatalf("(sub)Expression type was not expected. Expected: %T, Actual: %T", expectedStatement, actualStatement)
	}
	switch actualStatement.(type) {
	case *BlockStatement:
		expectedStmt := expectedStatement.(*BlockStatement)
		actualStmt := actualStatement.(*BlockStatement)
		if len(expectedStmt.Statements) != len(actualStmt.Statements) {
			t.Fatalf("Block Statement sizes do not match. Expected: %d, Actual: %d", len(expectedStmt.Statements), len(actualStmt.Statements))
		}
		for i, stmt := range expectedStmt.Statements {
			CompareStatements(stmt, actualStmt.Statements[i], t)
		}
	case *ExpressionStatement:
		expectedStmt := expectedStatement.(*ExpressionStatement)
		actualStmt := actualStatement.(*ExpressionStatement)
		CompareExpressions(actualStmt.Expression, expectedStmt.Expression, t)
	case *VariableAssignmentStatement:
		expectedStmt := expectedStatement.(*VariableAssignmentStatement)
		actualStmt := actualStatement.(*VariableAssignmentStatement)
		if expectedStmt.Name != actualStmt.Name {
			t.Fatalf("Variable names did not match. Expected %v, Actual: %v", expectedStmt.Name, actualStmt.Name)
		}
		CompareExpressions(actualStmt.Value, expectedStmt.Value, t)
	case *VariableDeclarationStatement:
		expectedStmt := expectedStatement.(*VariableDeclarationStatement)
		actualStmt := actualStatement.(*VariableDeclarationStatement)
		if expectedStmt.Name != actualStmt.Name {
			t.Fatalf("Variable names did not match. Expected %v, Actual: %v", expectedStmt.Name, actualStmt.Name)
		}
		CompareExpressions(actualStmt.Value, expectedStmt.Value, t)
	case *FunctionStatement:
		expectedStmt := expectedStatement.(*FunctionStatement)
		actualStmt := actualStatement.(*FunctionStatement)
		if expectedStmt.Name != actualStmt.Name {
			t.Fatalf("Function names did not match. Expected %v, Actual: %v", expectedStmt.Name, actualStmt.Name)
		}
		if expectedStmt.ReturnType.Compare(actualStmt.ReturnType) {
			t.Fatalf("Function Return Types did not match. Expected %T, Actual: %T", expectedStmt.ReturnType, actualStmt.ReturnType)
		}
		CompareStatements(actualStmt.Block, expectedStmt.Block, t)
	case *ReturnStatement:
		expectedStmt := expectedStatement.(*ReturnStatement)
		actualStmt := actualStatement.(*ReturnStatement)
		CompareExpressions(actualStmt.ReturnExpr, expectedStmt.ReturnExpr, t)
	default:
		t.Skipf("Undefined comparison for statement type %T", actualStatement)
	}

}
