package runtime

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/components/statements"
	. "compiler/src/runtime"
	"compiler/test"
	"fmt"
	"testing"
)

func Test_VariableStatements(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{
			[]Statement{
				&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
			},
			NewRuntimeEnvironment(map[string]Literal{"var1": NewLiteral(1, nil)}, nil, nil),
		},
		{
			[]Statement{
				&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
				&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 1},
			},
			NewRuntimeEnvironment(map[string]Literal{"var1": NewLiteral(2, nil)}, nil, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_VariableStatements%d", i)
		t.Run(testName, func(t *testing.T) {
			executor := Executor{tt.input, NewRuntimeEnvironment(nil, nil, nil), 0}
			_, err := executor.Execute()
			if err != nil {
				t.Error(err)
			}

			test.CompareRuntimeEnvironments(executor.Env, tt.expectedEnv, t)
		})
	}
}

func Test_BlockStatements(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{
			[]Statement{
				&BlockStatement{[]Statement{
					&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
				}, 0},
			},
			NewRuntimeEnvironment(nil, nil, nil),
		},
		{
			[]Statement{
				&BlockStatement{[]Statement{
					&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
					&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
				}, 0},
			},
			NewRuntimeEnvironment(nil, nil, nil),
		},
		{
			[]Statement{
				&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
				&BlockStatement{[]Statement{
					&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
					&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
				}, 0},
			},
			NewRuntimeEnvironment(map[string]Literal{"var1": NewLiteral(1, nil)}, nil, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_BlockStatements%d", i)
		t.Run(testName, func(t *testing.T) {
			executor := Executor{tt.input, NewRuntimeEnvironment(nil, nil, nil), 0}
			_, err := executor.Execute()
			if err != nil {
				t.Error(err)
			}

			test.CompareRuntimeEnvironments(executor.Env, tt.expectedEnv, t)
		})
	}
}

func Test_FunctionStatements(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{
			[]Statement{
				&FunctionStatement{"func1", make(map[string]Type), &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{}, 0},
			},
			NewRuntimeEnvironment(nil, map[string]*FunctionStatement{
				"func1": {"func1", make(map[string]Type), &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{}, 0},
			}, nil),
		},
		{
			[]Statement{
				&FunctionStatement{"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{}, 0},
			},
			NewRuntimeEnvironment(nil, map[string]*FunctionStatement{
				"func1": {"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{}, 0},
			}, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_FunctionStatements%d", i)
		t.Run(testName, func(t *testing.T) {
			executor := Executor{tt.input, NewRuntimeEnvironment(nil, nil, nil), 0}
			_, err := executor.Execute()
			if err != nil {
				t.Error(err)
			}

			test.CompareRuntimeEnvironments(executor.Env, tt.expectedEnv, t)
		})
	}
}

func Test_ReturnStatements(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{
			[]Statement{
				&FunctionStatement{"func1", make(map[string]Type), &BlockStatement{
					[]Statement{
						&ReturnStatement{
							Position:   0,
							ReturnExpr: NewLiteral(1, nil),
						},
					}, 0}, &IntType{}, 0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{nil, 0,
						make([]Expression, 0),
						&FunctionSignature{"func1", []*Parameter{}, &IntType{}},
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]Literal{
					"var1": NewLiteral(1, nil),
				},
				map[string]*FunctionStatement{
					"func1": {"func1", make(map[string]Type), &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: NewLiteral(1, nil),
							},
						}, 0}, &IntType{}, 0},
				}, nil),
		},
		{
			[]Statement{
				&FunctionStatement{"func1", make(map[string]Type), &BlockStatement{
					[]Statement{
						&ReturnStatement{
							Position:   0,
							ReturnExpr: &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), &IntType{}},
						},
					}, 0}, &IntType{}, 0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{nil, 0,
						make([]Expression, 0),
						&FunctionSignature{"func1", []*Parameter{}, &IntType{}},
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]Literal{
					"var1": NewLiteral(2, nil),
				},
				map[string]*FunctionStatement{
					"func1": {"func1", make(map[string]Type), &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), &IntType{}},
							},
						}, 0}, &IntType{}, 0},
				}, nil),
		},
		{
			[]Statement{
				&FunctionStatement{"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
					[]Statement{
						&ReturnStatement{
							Position:   0,
							ReturnExpr: &VariableExpr{"param1", &IntType{}, 0},
						},
					}, 0}, &IntType{}, 0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{nil, 0,
						[]Expression{NewLiteral(2, nil)},
						&FunctionSignature{"func1", []*Parameter{{&IntType{}, "param1"}}, &IntType{}},
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]Literal{
					"var1": NewLiteral(2, nil),
				},
				map[string]*FunctionStatement{
					"func1": {"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: &VariableExpr{"param1", &IntType{}, 0},
							},
						}, 0}, &IntType{}, 0},
				}, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_ReturnStatements%d", i)
		t.Run(testName, func(t *testing.T) {
			executor := Executor{tt.input, NewRuntimeEnvironment(nil, nil, nil), 0}
			_, err := executor.Execute()
			if err != nil {
				t.Error(err)
			}

			test.CompareRuntimeEnvironments(executor.Env, tt.expectedEnv, t)
		})
	}
}

func Test_ExpressionStatements(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{
			[]Statement{
				&ExpressionStatement{NewLiteral(1, nil), 0},
			},
			NewRuntimeEnvironment(nil, nil, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_ExpressionStatements%d", i)
		t.Run(testName, func(t *testing.T) {
			executor := Executor{tt.input, NewRuntimeEnvironment(nil, nil, nil), 0}
			_, err := executor.Execute()
			if err != nil {
				t.Error(err)
			}

			test.CompareRuntimeEnvironments(executor.Env, tt.expectedEnv, t)
		})
	}
}
