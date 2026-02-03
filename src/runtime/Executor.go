package runtime

import (
	"compiler/src/components/ast"
	. "compiler/src/components/statements"
	"errors"
	"fmt"
)

type Executor struct {
	Statements []ast.Statement
	Env        *RuntimeEnvironment
	LinePtr    int
}

func (e *Executor) Execute() (interface{}, error) {
	for ; e.LinePtr < len(e.Statements); e.LinePtr++ {
		_, err := executeStatement(e.Statements[e.LinePtr], e.Env)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error executing statement %d: %v", e.LinePtr, err))
		}
	}
	return nil, nil
}

func executeStatement(statement ast.Statement, env *RuntimeEnvironment) (interface{}, error) {
	return nil, nil
}

func ExecuteBlock(block ast.BlockStatement, env *RuntimeEnvironment) error {
	local := env.NewChild()
	for _, stmt := range block.Statements {
		if _, err := executeStatement(stmt, local); err != nil {
			return err
		}
	}
	return nil
}

func ExecuteExpression(expression ExpressionStatement, env *RuntimeEnvironment) error {
	_, err := Eval(expression.Expression, env)
	return err
}

func ExecuteVariableAssignment(varAssignment VariableAssignmentStatement, env *RuntimeEnvironment) error {
	val, err := Eval(varAssignment.Value, env)
	if err != nil {
		return err
	}
	return env.Assign(varAssignment.Name, &val)
}

func ExecuteVariableDeclaration(varDeclaration VariableDeclarationStatement, env *RuntimeEnvironment) error {
	val, err := Eval(varDeclaration.Value, env)
	if err != nil {
		return err
	}
	return env.Define(varDeclaration.Name, &val)
}
