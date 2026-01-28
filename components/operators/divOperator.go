package operators

import (
	. "compiler/components"
	"errors"
)

type DivOperator struct{}

func (op DivOperator) Apply(args ...Expression) (*Literal, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments for Div Operator")
	}

	left := args[0]
	right := args[1]

	leftLiteral, _ := left.Eval()
	rightLiteral, _ := right.Eval()

	switch {
	case leftLiteral.Kind == IntValue && rightLiteral.Kind == IntValue:
		if rightLiteral.Int == 0 {
			return nil, errors.New("division by zero")
		}
		return &Literal{Kind: IntValue, Int: leftLiteral.Int / rightLiteral.Int}, nil
	default:
		return nil, errors.New("invalid Operations")
	}

}
