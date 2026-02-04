package runtime

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/functions"
	. "compiler/src/components/literals"
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

func (ev *Evaluator) Eval(node Node) (Literal, error) {
	if _, ok := node.(Expression); ok {
		nodeExpr := node.(Expression)
		eval, evalErr := ev.EvalExpr(&nodeExpr)
		return eval, evalErr
	}
	return nil, errors.New(fmt.Sprintf("Unknown Node Type: %T", node))
}

func (ev *Evaluator) EvalExpr(exprNode *Expression) (Literal, error) {
	switch (*exprNode).(type) {
	case Literal:
		litExpr := (*exprNode).(Literal)
		return litExpr, nil
	case *UnaryExpr:
		unaryExpr := (*exprNode).(*UnaryExpr)
		literal, err := evalUnary(&unaryExpr.Op, &unaryExpr.Arg, ev)
		return *literal, err
	case *BinaryExpr:
		binExpr := (*exprNode).(*BinaryExpr)
		literal, err := evalBinary(&binExpr.Op, &binExpr.Left, &binExpr.Right, ev)
		if literal == nil {
			return nil, err
		}
		return *literal, err
	case *TernaryExpr:
		ternaryExpr := (*exprNode).(*TernaryExpr)
		literal, err := evalTernary(&ternaryExpr.Op, &ternaryExpr.Condition, &ternaryExpr.Left, &ternaryExpr.Right, ev)
		if literal == nil {
			return nil, err
		}
		return *literal, err
	case *VariableExpr:
		varExpr := (*exprNode).(*VariableExpr)
		varLit, err := ev.env.GetVariable(varExpr.Name)
		if varLit == nil {
			return nil, err
		}
		return *varLit, err
	case *CallExpr:
		callExpr := (*exprNode).(*CallExpr)
		funcStmt, funcStmtErr := ev.parent.Env.GetFunction(callExpr.Name)
		if funcStmtErr != nil {
			return nil, funcStmtErr
		}
		returnLit, funcErr := ev.parent.ExecuteFunction(funcStmt)
		if returnLit == nil {
			return nil, funcErr
		}
		return *returnLit, funcErr
	}
	if condExpr, ok := (*exprNode).(Condition); ok {
		boolLit, err := ev.EvalCondition(&condExpr)
		return boolLit, err
	}

	return nil, errors.New(fmt.Sprintf("Unknown Expression Type: %T", exprNode))
}

func (ev *Evaluator) EvalCondition(condNode *Condition) (*BoolLiteral, error) {
	switch (*condNode).(type) {
	case *BoolLiteral:
		return (*condNode).(*BoolLiteral), nil
	case *BinaryCondition:
		binCondition := (*condNode).(*BinaryCondition)
		binaryLit, err := evalBinary(&binCondition.Op, &binCondition.Left, &binCondition.Right, ev)
		if binaryLit == nil {
			return nil, err
		}
		return *binaryLit, err
	case *UnaryCondition:
		unaryCondition := (*condNode).(*UnaryCondition)
		unaryLit, err := evalUnary(&unaryCondition.Op, &unaryCondition.Arg, ev)
		if unaryLit == nil {
			return nil, err
		}
		return *unaryLit, err
	}
	return nil, errors.New(fmt.Sprintf("Unknown Condition Type: %T", condNode))
}

func evalUnary[T Literal](op *UnaryOperator[T], arg *Expression, evaluator *Evaluator) (*T, error) {
	argExpr, err := evaluator.EvalExpr(arg)
	if err != nil {
		return nil, err
	}
	evalVal, evalErr := (*op).Apply(&argExpr)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}

func evalBinary[T Literal](op *BinaryOperator[T], left *Expression, right *Expression, evaluator *Evaluator) (*T, error) {
	leftEval, err := evaluator.EvalExpr(left)
	if err != nil {
		return nil, err
	}
	rightEval, err2 := evaluator.EvalExpr(right)
	if err2 != nil {
		return nil, err2
	}
	evalVal, evalErr := (*op).Apply(&leftEval, &rightEval)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}

func evalTernary[T Literal](op *TernaryOperator[T], condition *Condition, left *Expression, right *Expression, evaluator *Evaluator) (*T, error) {
	condEval, err := evaluator.EvalCondition(condition)
	if err != nil {
		return nil, err
	}
	leftEval, err2 := evaluator.EvalExpr(left)
	if err2 != nil {
		return nil, err2
	}
	rightEval, err3 := evaluator.EvalExpr(right)
	if err3 != nil {
		return nil, err3
	}
	evalVal, evalErr := (*op).Apply(condEval, &leftEval, &rightEval)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}
