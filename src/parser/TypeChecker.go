package parser

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/statements"
	. "compiler/src/components/variables"
	"compiler/src/runtime/value"
	"errors"
	"fmt"
)

type TypeChecker struct {
	*BlockStatement
	*VarSymbolTable
	*FuncSymbolTable
}

func NewTypeChecker(block *BlockStatement, varSymbols *VarSymbolTable, funcSymbols *FuncSymbolTable) *TypeChecker {
	if varSymbols == nil {
		varSymbols = &VarSymbolTable{}
	}
	if funcSymbols == nil {
		funcSymbols = &FuncSymbolTable{}
	}
	return &TypeChecker{block, varSymbols, funcSymbols}
}

func (tc *TypeChecker) CheckTypes() error {
	for _, stmt := range tc.BlockStatement.Statements {
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
	args := make([]*Parameter, len(statement.Params))
	for name, paramType := range statement.Params {
		args = append(args, &Parameter{name, paramType})
	}

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

	statement.FunctionIdentifier = value.NewFunctionIdentifier(statement.Name)

	signature := &FunctionSignature{statement.FunctionIdentifier, args, statement.ReturnType}
	tc.FuncSymbolTable.Declare(statement.FunctionIdentifier, signature)
	return statement.ReturnType, nil
}
func (tc *TypeChecker) checkReturnStatement(statement *ReturnStatement) (Type, error) {
	return tc.checkExpression(statement.ReturnExpr)
}
func (tc *TypeChecker) checkVariableAssignmentStatement(statement *VariableAssignmentStatement) error {
	expectedType := tc.VarSymbolTable.Get(statement.Name)
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
	switch expr.(type) {
	case *VariableExpr:
		return tc.VarSymbolTable.Get(expr.(*VariableExpr).Identifier), nil
	case *FunctionExpr:
		signature := tc.FuncSymbolTable.Get(expr.(*FunctionExpr).Identifier)
		if signature == nil {
			return nil, errors.New(fmt.Sprintf("Function Signature '%v' not found", expr.(*FunctionExpr).Identifier))
		}
		expr.(*FunctionExpr).SetFunctionSignature(signature)
		return signature, nil
	case *CallExpr:
		callExpr := expr.(*CallExpr)

		signatureType, signatureErr := tc.checkExpression(callExpr.Callee)
		if signatureErr != nil {
			return nil, signatureErr
		}

		signature := signatureType.(*FunctionSignature)
		if len(callExpr.Params) != len(signature.Params) {
			return nil, errors.New(fmt.Sprintf("Incorrect arg count for %v. Expected %d, Actual: %d",
				signature.Name, len(signature.Params), len(callExpr.Params)))
		}

		for i, param := range callExpr.Params {
			paramType, err := tc.checkExpression(param)
			if err != nil {
				return nil, err
			}

			if !paramType.Compare(signature.Params[i].ValType) {
				return nil, errors.New(fmt.Sprintf("Incorrect param type for %v. Expected %T, Actual: %T",
					signature.Name, signature.Params[i].ValType, paramType))
			}
		}

		return callExpr.Type()
	default:
		return expr.Type()
	}
}
