package operators

import (
	. "compiler/components"
	"errors"
)

type MultOperator struct{}

func (op MultOperator) Apply(args ...Expression) (*Literal, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments for Mult Operator")
	}

	left := args[0]
	right := args[1]

	leftLiteral, _ := left.Eval()
	rightLiteral, _ := right.Eval()

	switch {
	case leftLiteral.Kind == IntValue && rightLiteral.Kind == IntValue:
		return &Literal{Kind: IntValue, Int: leftLiteral.Int * rightLiteral.Int}, nil
	default:
		return nil, errors.New("invalid Operations")
	}

}
