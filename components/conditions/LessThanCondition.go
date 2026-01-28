package conditions

import (
	. "compiler/components"
	. "compiler/components/literals"
	"errors"
)

type LessThanCondition struct {
	*BinaryCondition
}

func (e *LessThanCondition) Eval() (Literal, error) {
	return e.check(e.Left, e.Right)
}

func (e *LessThanCondition) check(args ...Expression) (*BoolLiteral, error) {
	if len(args) != 2 {
		return nil, errors.New("expected 2 arguments")
	}

	leftLiteral, err := args[0].Eval()
	rightLiteral, err2 := args[1].Eval()

	if err != nil || err2 != nil {
		return nil, err
	}

	if leftLiteral.Kind() == BoolKind || rightLiteral.Kind() == BoolKind {
		return nil, errors.New("Invalid comparison")
	}

	if leftLiteral.Kind() != rightLiteral.Kind() {
		return nil, errors.New("Invalid comparison")
	}

	switch leftLiteral.Kind() {
	case IntKind:
		return &BoolLiteral{leftLiteral.(*IntLiteral).Value < rightLiteral.(*IntLiteral).Value}, nil
	case StringKind:
		return &BoolLiteral{leftLiteral.(*StringLiteral).Value < rightLiteral.(*StringLiteral).Value}, nil
	default:
		panic("unhandled default case")
	}
}
