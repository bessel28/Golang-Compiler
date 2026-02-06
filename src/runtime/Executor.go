package runtime

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/statements"
	"compiler/src/runtime/value"
	"errors"
	"fmt"
)

type Executor struct {
	*BlockStatement
	Env     *RuntimeEnvironment
	LinePtr int
}

func (ex *Executor) Execute() (value.RuntimeValue, error) {
	for ; ex.LinePtr < len(ex.Statements); ex.LinePtr++ {
		_, err := ex.executeStatement(ex.Statements[ex.LinePtr])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error executing statement %d: %v", ex.LinePtr, err))
		}
	}
	return nil, nil
}

func (ex *Executor) executeStatement(statement Statement) (value.RuntimeValue, error) {
	switch statement.(type) {
	case *BlockStatement:
		return ex.ExecuteBlock(statement.(*BlockStatement))
	case *ExpressionStatement:
		return nil, ex.ExecuteExpression(statement.(*ExpressionStatement))
	case *VariableAssignmentStatement:
		return nil, ex.ExecuteVariableAssignment(statement.(*VariableAssignmentStatement))
	case *VariableDeclarationStatement:
		return nil, ex.ExecuteVariableDeclaration(statement.(*VariableDeclarationStatement))
	case *FunctionStatement:
		return nil, ex.ExecuteFunction(statement.(*FunctionStatement))
	case *ReturnStatement:
		return ex.ExecuteReturnBlock(statement.(*ReturnStatement))
	}
	return nil, errors.New(fmt.Sprintf("Unsupported statement type: %T", statement))
}

func (ex *Executor) ExecuteBlock(block *BlockStatement) (value.RuntimeValue, error) {
	ex.Env = ex.Env.NewChild()
	var returnVal value.RuntimeValue
	for _, stmt := range block.Statements {
		statementReturn, err := ex.executeStatement(stmt)
		if err != nil {
			ex.Env = ex.Env.Parent
			return nil, err
		}
		returnVal = statementReturn
	}
	ex.Env = ex.Env.Parent
	return returnVal, nil
}

func (ex *Executor) ExecuteExpression(expression *ExpressionStatement) error {
	_, err := (&Evaluator{ex.Env, ex}).Eval(expression.Expression)
	return err
}

func (ex *Executor) ExecuteVariableAssignment(varAssignment *VariableAssignmentStatement) error {
	val, err := (&Evaluator{ex.Env, ex}).Eval(varAssignment.Value)
	if err != nil {
		return err
	}
	return ex.Env.AssignVariable(varAssignment.Name, val)
}

func (ex *Executor) ExecuteVariableDeclaration(varDeclaration *VariableDeclarationStatement) error {
	val, err := (&Evaluator{ex.Env, ex}).Eval(varDeclaration.Value)
	if err != nil {
		return err
	}
	return ex.Env.DefineVariable(varDeclaration.Name, val)
}

func (ex *Executor) ExecuteFunction(function *FunctionStatement) error {
	err := ex.Env.DefineFunction(function.FunctionSignature.FunctionIdentifier, function)
	if err != nil {
		return err
	}
	return nil
}

func (ex *Executor) ExecuteReturnBlock(returnStmt *ReturnStatement) (value.RuntimeValue, error) {
	returnLit, returnErr := (&Evaluator{ex.Env, ex}).Eval(returnStmt.ReturnExpr)
	if returnErr != nil {
		return nil, returnErr
	}
	return returnLit, returnErr
}
