package parser

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/statements"
	"errors"
	"fmt"
)

type TypeChecker struct {
	Statements *[]Statement
	*VarSymbolTable
}

func NewTypeChecker(statements *[]Statement) *TypeChecker {
	return &TypeChecker{statements, &VarSymbolTable{}}
}

func (tc *TypeChecker) CheckTypes() error {
	for _, stmt := range *tc.Statements {
		_, err := tc.CheckType(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tc *TypeChecker) CheckType(statement Statement) (Type, error) {
	stmt := statement
	switch stmt.(type) {
	case *BlockStatement:
		return nil, tc.checkBlockStatement(stmt.(*BlockStatement))
	case *ExpressionStatement:
		return nil, tc.checkExpressionStatement(stmt.(*ExpressionStatement))
	case *FunctionStatement:
		return tc.checkFunctionStatement(stmt.(*FunctionStatement))
	case *ReturnStatement:
		return tc.checkReturnStatement(stmt.(*ReturnStatement))
	case *VariableAssignmentStatement:
		return nil, tc.checkVariableAssignmentStatement(stmt.(*VariableAssignmentStatement))
	case *VariableDeclarationStatement:
		return nil, tc.checkVariableDeclarationStatement(stmt.(*VariableDeclarationStatement))
	default:
		return nil, errors.New(fmt.Sprintf("Undefined typechecking for: %T", statement))
	}
}

func (tc *TypeChecker) checkBlockStatement(statement *BlockStatement) error {
	tc.VarSymbolTable = tc.VarSymbolTable.Child()
	for _, stmt := range statement.Statements {
		_, err := tc.CheckType(stmt)
		if err != nil {
			tc.VarSymbolTable = tc.VarSymbolTable.Parent
			return err
		}
	}
	tc.VarSymbolTable = tc.VarSymbolTable.Parent
	return nil
}

func (tc *TypeChecker) checkExpressionStatement(statement *ExpressionStatement) error {
	_, err := tc.checkExpression(statement.Expression)
	if err != nil {
		return err
	}
	return nil
}
func (tc *TypeChecker) checkFunctionStatement(statement *FunctionStatement) (Type, error) {
	varSymbols := tc.VarSymbolTable
	tc.VarSymbolTable = &VarSymbolTable{}
	for _, stmt := range statement.Block.Statements {
		switch stmt.(type) {
		case *ReturnStatement:
			rtnType, rtnErr := tc.checkReturnStatement(stmt.(*ReturnStatement))
			if rtnErr != nil {
				tc.VarSymbolTable = varSymbols
				return nil, rtnErr
			}
			if rtnType != statement.ReturnType {
				tc.VarSymbolTable = varSymbols
				return nil, errors.New(fmt.Sprintf("Returned value's type didn't match expected type. Expected: %T, Actual: %T", statement.ReturnType, rtnType))
			}
		default:
			_, err := tc.CheckType(stmt)
			if err != nil {
				tc.VarSymbolTable = varSymbols
				return nil, err
			}
		}
	}
	tc.VarSymbolTable = varSymbols
	return statement.ReturnType, nil
}
func (tc *TypeChecker) checkReturnStatement(statement *ReturnStatement) (Type, error) {
	return tc.checkExpression(statement.ReturnExpr)
}
func (tc *TypeChecker) checkVariableAssignmentStatement(statement *VariableAssignmentStatement) error {
	expectedType := *tc.VarSymbolTable.GetType(statement.Name)
	exprType, exprErr := tc.checkExpression(statement.Value)
	if exprErr != nil {
		return exprErr
	}
	if expectedType != exprType {
		return errors.New(fmt.Sprintf("%v (T%) cannot be assigned to %v (%T)", statement.Value, exprType, statement.Name, expectedType))
	}
	return nil
}
func (tc *TypeChecker) checkVariableDeclarationStatement(statement *VariableDeclarationStatement) error {
	exprType, exprErr := tc.checkExpression(statement.Value)
	if exprErr != nil {
		return exprErr
	}

	tc.VarSymbolTable.Declare(statement.Name, exprType)
	return nil
}

func (tc *TypeChecker) checkExpression(expr Expression) (Type, error) {
	return expr.Type()
	//switch expr.(type) {
	//case Literal:
	//	return expr.Type()
	//case *UnaryExpr:
	//	return expr.(*UnaryExpr).Type()
	//case *BinaryExpr:
	//	return expr.(*BinaryExpr).Type()
	//case *TernaryExpr:
	//	return expr.(*TernaryExpr).Type()
	//case *BinaryCondition:
	//	return expr.(*BinaryCondition).Type()
	//case *UnaryCondition:
	//	return expr.(*UnaryCondition).Type()
	//case *functions.CallExpr:
	//	return expr.(*functions.CallExpr).Type()
	//}
	//return nil, errors.New("Unsupported expression during typecheck")
}
