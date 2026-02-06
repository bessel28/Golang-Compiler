package runtime

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	. "compiler/src/components/statements"
	. "compiler/src/components/variables"
	. "compiler/src/runtime"
	. "compiler/src/runtime/value"
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
			NewRuntimeEnvironment(map[string]RuntimeValue{"var1": &IntValue{1}}, nil, nil),
		},
		{
			[]Statement{
				&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
				&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 1},
			},
			NewRuntimeEnvironment(map[string]RuntimeValue{"var1": &IntValue{2}}, nil, nil),
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
			NewRuntimeEnvironment(map[string]RuntimeValue{"var1": &IntValue{1}}, nil, nil),
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
					}, 0}, &NilType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &NilType{}},
					0},
			},
			NewRuntimeEnvironment(nil, map[FunctionIdentifier]*FunctionStatement{
				{"func1", 0}: {"func1", make(map[string]Type), &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &NilType{}},
					0},
			}, nil),
		},
		{
			[]Statement{
				&FunctionStatement{"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0},
						[]*Parameter{{"param1", &IntType{}}},
						&NilType{}},
					0},
			},
			NewRuntimeEnvironment(nil, map[FunctionIdentifier]*FunctionStatement{
				{"func1", 0}: {"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
					[]Statement{
						&VariableDeclarationStatement{"var1", NewLiteral(1, nil), 0},
						&VariableAssignmentStatement{"var1", NewLiteral(2, nil), 0},
					}, 0}, &NilType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0},
						[]*Parameter{{"param1", &IntType{}}},
						&NilType{}},
					0},
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
					}, 0}, &IntType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0},
						[]*Parameter{},
						&IntType{}},
					0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{
						&FunctionExpr{&FunctionIdentifier{"func1", 0}, &FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &IntType{}}, 0},
						make([]Expression, 0),
						0,
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]RuntimeValue{
					"var1": &IntValue{1},
				},
				map[FunctionIdentifier]*FunctionStatement{
					{"func1", 0}: {"func1", make(map[string]Type), &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: NewLiteral(1, nil),
							},
						}, 0}, &IntType{},
						&FunctionSignature{&FunctionIdentifier{"func1", 0},
							[]*Parameter{},
							&IntType{}},
						0},
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
					}, 0}, &IntType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0},
						[]*Parameter{},
						&IntType{}},
					0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{
						&FunctionExpr{&FunctionIdentifier{"func1", 0}, &FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &IntType{}}, 0},
						make([]Expression, 0),
						0,
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]RuntimeValue{
					"var1": &IntValue{2},
				},
				map[FunctionIdentifier]*FunctionStatement{
					{"func1", 0}: {"func1", make(map[string]Type), &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: &BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), &IntType{}},
							},
						}, 0}, &IntType{},
						&FunctionSignature{&FunctionIdentifier{"func1", 0},
							[]*Parameter{},
							&IntType{}},
						0},
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
					}, 0}, &IntType{},
					&FunctionSignature{&FunctionIdentifier{"func1", 0},
						[]*Parameter{{"param1", &IntType{}}},
						&IntType{}},
					0},
				&VariableDeclarationStatement{
					"var1", &CallExpr{
						&FunctionExpr{&FunctionIdentifier{"func1", 0}, &FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{{"param1", &IntType{}}}, &IntType{}}, 0},
						[]Expression{NewLiteral(2, nil)},
						0,
					}, 0},
			},
			NewRuntimeEnvironment(
				map[string]RuntimeValue{
					"var1": &IntValue{2},
				},
				map[FunctionIdentifier]*FunctionStatement{
					{"func1", 0}: {"func1", map[string]Type{"param1": &IntType{}}, &BlockStatement{
						[]Statement{
							&ReturnStatement{
								Position:   0,
								ReturnExpr: &VariableExpr{"param1", &IntType{}, 0},
							},
						}, 0}, &IntType{},
						&FunctionSignature{&FunctionIdentifier{"func1", 0},
							[]*Parameter{{"param1", &IntType{}}},
							&IntType{}},
						0},
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

func Test_FuncVariables(t *testing.T) {
	var tests = []struct {
		input       []Statement
		expectedEnv *RuntimeEnvironment
	}{
		{[]Statement{
			&FunctionStatement{"func1", make(map[string]Type), &BlockStatement{
				[]Statement{
					&ReturnStatement{0, NewLiteral(1, nil)},
				}, 0}, &IntType{},
				&FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &IntType{}},
				0},
			&VariableDeclarationStatement{
				"var1",
				&FunctionExpr{&FunctionIdentifier{"func1", 0}, &FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &IntType{}}, 0},
				1,
			},
			&VariableDeclarationStatement{
				"var2",
				&CallExpr{
					&FunctionExpr{&FunctionIdentifier{"func1", 0}, &FunctionSignature{&FunctionIdentifier{"func", 0}, []*Parameter{}, &IntType{}}, 0},
					[]Expression{},
					2,
				},
				2,
			},
		},
			NewRuntimeEnvironment(
				map[string]RuntimeValue{
					"var1": &FuncValue{&FunctionIdentifier{"func1", 0}},
					"var2": &IntValue{1},
				},
				map[FunctionIdentifier]*FunctionStatement{
					{"func1", 0}: {"func1", make(map[string]Type), &BlockStatement{
						[]Statement{
							&ReturnStatement{0, NewLiteral(1, nil)},
						}, 0}, &IntType{},
						&FunctionSignature{&FunctionIdentifier{"func1", 0}, []*Parameter{}, &IntType{}},
						0},
				}, nil),
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test_FuncVariables%d", i)
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
