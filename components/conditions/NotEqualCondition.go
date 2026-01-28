package conditions

import (
	. "compiler/components"
	. "compiler/components/literals"
	"errors"
)

type NotEqualCondition struct {
	*BinaryCondition
}

func (e *NotEqualCondition) Eval() (Literal, error) {
	return e.check(e.Left, e.Right)
}

func (e *NotEqualCondition) check(args ...Expression) (*BoolLiteral, error) {
	if len(args) != 2 {
		return nil, errors.New("expected 2 arguments")
	}

	leftLiteral, err := args[0].Eval()
	rightLiteral, err2 := args[1].Eval()

	if err != nil || err2 != nil {
		return nil, err
	}

	if leftLiteral.Equals(rightLiteral).Value {
		return &BoolLiteral{Value: false}, nil
	}

	return &BoolLiteral{Value: true}, nil
}
