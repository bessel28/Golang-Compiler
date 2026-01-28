package operators

import (
	. "compiler/components"
	. "compiler/components/literals"
	"errors"
)

type AddOperator struct{}

func (op AddOperator) Apply(args ...Expression) (Literal, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments for Add Operator")
	}

	left := args[0]
	right := args[1]

	leftLiteral, _ := left.Eval()
	rightLiteral, _ := right.Eval()

	switch {
	case leftLiteral.Kind() == IntKind && rightLiteral.Kind() == IntKind:
		return &IntLiteral{leftLiteral.(*IntLiteral).Value + rightLiteral.(*IntLiteral).Value}, nil
	default:
		return &StringLiteral{leftLiteral.String() + rightLiteral.String()}, nil
	}
}
