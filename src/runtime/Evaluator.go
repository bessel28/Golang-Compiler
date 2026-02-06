package runtime

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
	. "compiler/src/components/variables"
	. "compiler/src/runtime/value"
	"errors"
	"fmt"
)

type Evaluator struct {
	env    *RuntimeEnvironment
	parent *Executor
}

func NewEvaluator(env *RuntimeEnvironment, parent *Executor) *Evaluator {
	return &Evaluator{env, parent}
}

func (ev *Evaluator) Eval(node Node) (RuntimeValue, error) {
	if _, ok := node.(Expression); ok {
		nodeExpr := node.(Expression)
		eval, evalErr := ev.EvalExpr(nodeExpr)
		return eval, evalErr
	}
	return nil, errors.New(fmt.Sprintf("Unknown Node Type: %T", node))
}

func (ev *Evaluator) EvalExpr(exprNode Expression) (RuntimeValue, error) {
	// TODO Implement FunctionExpr in this switch using runtime Values instead of Literals
	switch exprNode.(type) {
	case Literal:
		litExpr := exprNode.(Literal)
		return NewValue(litExpr), nil
	case *UnaryExpr:
		unaryExpr := exprNode.(*UnaryExpr)
		value, err := evalUnary(&unaryExpr.Op, unaryExpr.Arg, ev)
		return value, err
	case *BinaryExpr:
		binExpr := exprNode.(*BinaryExpr)
		value, err := evalBinary(&binExpr.Op, binExpr.Left, binExpr.Right, ev)
		if value == nil {
			return nil, err
		}
		return value, err
	case *TernaryExpr:
		ternaryExpr := exprNode.(*TernaryExpr)
		value, err := evalTernary(&ternaryExpr.Op, ternaryExpr.Condition, ternaryExpr.Left, ternaryExpr.Right, ev)
		if value == nil {
			return nil, err
		}
		return value, err
	case *VariableExpr:
		varExpr := exprNode.(*VariableExpr)
		varLit, err := ev.env.GetVariable(varExpr.Identifier)
		return varLit, err
	case *CallExpr:
		callExpr := exprNode.(*CallExpr)
		funcIdentifier, identifierErr := ev.EvalCallable(callExpr.Callee)
		if identifierErr != nil {
			return nil, identifierErr
		}

		funcStmt, funcStmtErr := ev.parent.Env.GetFunction(funcIdentifier)
		if funcStmtErr != nil {
			return nil, funcStmtErr
		}

		env := NewRuntimeEnvironment(nil, nil, nil)
		funcExecutor := Executor{[]Statement{funcStmt.Block}, env, 0}

		for paramIndex, param := range callExpr.Params {
			evaluator := NewEvaluator(env, &funcExecutor)
			value, evalErr := evaluator.Eval(param)
			if evalErr != nil {
				return nil, evalErr
			}

			paramSignature := callExpr.Callee.GetFunctionSignature().Params[paramIndex]
			paramDefErr := env.DefineVariable(paramSignature.Name, value)
			if paramDefErr != nil {
				return nil, paramDefErr
			}
		}

		returnLit, funcErr := funcExecutor.ExecuteBlock(funcStmt.Block)
		return returnLit, funcErr
	case *FunctionExpr:
		funcExpr := exprNode.(*FunctionExpr)
		return &FuncValue{funcExpr.Identifier}, nil
	}
	if condExpr, ok := exprNode.(Condition); ok {
		boolLit, err := ev.EvalCondition(condExpr)
		return boolLit, err
	}

	return nil, errors.New(fmt.Sprintf("Unknown Expression Type: %T", exprNode))
}

func (ev *Evaluator) EvalCondition(condNode Condition) (*BoolValue, error) {
	switch condNode.(type) {
	case *BoolLiteral:
		return NewValue(condNode.(*BoolLiteral)).(*BoolValue), nil
	case *BinaryCondition:
		binCondition := condNode.(*BinaryCondition)
		binaryLit, err := evalBinary(&binCondition.Op, binCondition.Left, binCondition.Right, ev)
		if binaryLit == nil {
			return nil, err
		}
		return binaryLit, err
	case *UnaryCondition:
		unaryCondition := condNode.(*UnaryCondition)
		unaryLit, err := evalUnary(&unaryCondition.Op, unaryCondition.Arg, ev)
		if unaryLit == nil {
			return nil, err
		}
		return unaryLit, err
	}
	return nil, errors.New(fmt.Sprintf("Unknown Condition Type: %T", condNode))
}

func evalUnary[T RuntimeValue](op *UnaryOperator[T], arg Expression, evaluator *Evaluator) (T, error) {
	var zeroVal T
	argExpr, err := evaluator.EvalExpr(arg)
	if err != nil {
		return zeroVal, err
	}
	evalVal, evalErr := (*op).Apply(argExpr)
	if evalErr != nil {
		return zeroVal, evalErr
	}
	return evalVal, nil
}

func evalBinary[T RuntimeValue](op *BinaryOperator[T], left Expression, right Expression, evaluator *Evaluator) (T, error) {
	var zeroVal T
	leftEval, err := evaluator.EvalExpr(left)
	if err != nil {
		return zeroVal, err
	}
	rightEval, err2 := evaluator.EvalExpr(right)
	if err2 != nil {
		return zeroVal, err2
	}
	evalVal, evalErr := (*op).Apply(leftEval, rightEval)
	if evalErr != nil {
		return zeroVal, evalErr
	}
	return evalVal, nil
}

func evalTernary[T RuntimeValue](op *TernaryOperator[T], condition Condition, left Expression, right Expression, evaluator *Evaluator) (T, error) {
	var zeroVal T
	condEval, err := evaluator.EvalCondition(condition)
	if err != nil {
		return zeroVal, err
	}
	leftEval, err2 := evaluator.EvalExpr(left)
	if err2 != nil {
		return zeroVal, err2
	}
	rightEval, err3 := evaluator.EvalExpr(right)
	if err3 != nil {
		return zeroVal, err3
	}
	evalVal, evalErr := (*op).Apply(condEval, leftEval, rightEval)
	if evalErr != nil {
		return zeroVal, evalErr
	}
	return evalVal, nil
}

func (ev *Evaluator) EvalCallable(callable Callable) (*FunctionIdentifier, error) {
	switch callable.(type) {
	case *FunctionExpr:
		return callable.(*FunctionExpr).Identifier, nil
	case *VariableExpr:
		varEval, evalErr := ev.EvalExpr(callable.(*VariableExpr))
		if evalErr != nil {
			return nil, evalErr
		}
		funcEval, ok := varEval.(*FuncValue)
		if !ok {
			return nil, errors.New("Variable RuntimeValue was not callable")
		}
		return funcEval.Val, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown callable type: %T", callable))
	}
}
