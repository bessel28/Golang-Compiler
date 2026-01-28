package conditions

import (
	. "compiler/components"
	. "compiler/components/literals"
	"errors"
)

type EqualCondition struct {
	*BinaryCondition
}

func (e *EqualCondition) Eval() (Literal, error) {
	return e.check(e.Left, e.Right)
}

func (e *EqualCondition) check(args ...Expression) (*BoolLiteral, error) {
	if len(args) != 2 {
		return nil, errors.New("expected 2 arguments")
	}

	leftLiteral, err := args[0].Eval()
	rightLiteral, err2 := args[1].Eval()

	if err != nil || err2 != nil {
		return nil, err
	}

	return leftLiteral.Equals(rightLiteral), nil
}
