package runtime

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	"errors"
	"fmt"
)

type Evaluator struct{}

func Eval(node Node, env *RuntimeEnvironment) (Literal, error) {
	if _, ok := node.(Expression); ok {
		nodeExpr := node.(Expression)
		eval, evalErr := EvalExpr(&nodeExpr, env)
		return eval, evalErr
	}
	return nil, errors.New(fmt.Sprintf("Unknown Node Type: %T", node))

}

func EvalExpr(exprNode *Expression, env *RuntimeEnvironment) (Literal, error) {
	switch (*exprNode).(type) {
	case Literal:
		litExpr := (*exprNode).(Literal)
		return litExpr, nil
	case *UnaryExpr:
		unaryExpr := (*exprNode).(*UnaryExpr)
		literal, err := evalUnary(&unaryExpr.Op, &unaryExpr.Arg, env)
		return *literal, err
	case *BinaryExpr:
		binExpr := (*exprNode).(*BinaryExpr)
		literal, err := evalBinary(&binExpr.Op, &binExpr.Left, &binExpr.Right, env)
		if literal == nil {
			return nil, err
		}
		return *literal, err
	case *TernaryExpr:
		ternaryExpr := (*exprNode).(*TernaryExpr)
		literal, err := evalTernary(&ternaryExpr.Op, &ternaryExpr.Condition, &ternaryExpr.Left, &ternaryExpr.Right, env)
		if literal == nil {
			return nil, err
		}
		return *literal, err
	}
	if condExpr, ok := (*exprNode).(Condition); ok {
		boolLit, err := EvalCondition(&condExpr, env)
		return boolLit, err
	}

	return nil, errors.New(fmt.Sprintf("Unknown Expression Type: %T", exprNode))
}

func EvalCondition(condNode *Condition, env *RuntimeEnvironment) (*BoolLiteral, error) {
	switch (*condNode).(type) {
	case *BoolLiteral:
		return (*condNode).(*BoolLiteral), nil
	case *BinaryCondition:
		binCondition := (*condNode).(*BinaryCondition)
		binaryLit, err := evalBinary(&binCondition.Op, &binCondition.Left, &binCondition.Right, env)
		if binaryLit == nil {
			return nil, err
		}
		return *binaryLit, err
	case *UnaryCondition:
		unaryCondition := (*condNode).(*UnaryCondition)
		unaryLit, err := evalUnary(&unaryCondition.Op, &unaryCondition.Arg, env)
		if unaryLit == nil {
			return nil, err
		}
		return *unaryLit, err
	}
	return nil, errors.New(fmt.Sprintf("Unknown Condition Type: %T", condNode))
}

func evalUnary[T Literal](op *UnaryOperator[T], arg *Expression, env *RuntimeEnvironment) (*T, error) {
	argExpr, err := EvalExpr(arg, env)
	if err != nil {
		return nil, err
	}
	evalVal, evalErr := (*op).Apply(&argExpr)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}

func evalBinary[T Literal](op *BinaryOperator[T], left *Expression, right *Expression, env *RuntimeEnvironment) (*T, error) {
	leftEval, err := EvalExpr(left, env)
	if err != nil {
		return nil, err
	}
	rightEval, err2 := EvalExpr(right, env)
	if err2 != nil {
		return nil, err2
	}
	evalVal, evalErr := (*op).Apply(&leftEval, &rightEval)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}

func evalTernary[T Literal](op *TernaryOperator[T], condition *Condition, left *Expression, right *Expression, env *RuntimeEnvironment) (*T, error) {
	condEval, err := EvalCondition(condition, env)
	if err != nil {
		return nil, err
	}
	leftEval, err2 := EvalExpr(left, env)
	if err2 != nil {
		return nil, err2
	}
	rightEval, err3 := EvalExpr(right, env)
	if err3 != nil {
		return nil, err3
	}
	evalVal, evalErr := (*op).Apply(condEval, &leftEval, &rightEval)
	if evalErr != nil {
		return nil, evalErr
	}
	return &evalVal, nil
}
