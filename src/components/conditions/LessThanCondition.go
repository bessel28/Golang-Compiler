package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	"errors"
)

type LessThanCondition struct {
	*BinaryCondition
}

func (e *LessThanCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *LessThanCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()
	rightLiteral, err2 := e.Right.Eval()

	if err != nil || err2 != nil {
		return false, err
	}

	if leftLiteral.Kind() == BoolKind || rightLiteral.Kind() == BoolKind {
		return false, errors.New("Invalid comparison")
	}

	if leftLiteral.Kind() != rightLiteral.Kind() {
		return false, errors.New("Invalid comparison")
	}

	switch leftLiteral.Kind() {
	case IntKind:
		return leftLiteral.(*IntLiteral).Value < rightLiteral.(*IntLiteral).Value, nil
	case StringKind:
		return leftLiteral.(*StringLiteral).Value < rightLiteral.(*StringLiteral).Value, nil
	default:
		return false, errors.New("unhandled default case")
	}
}
