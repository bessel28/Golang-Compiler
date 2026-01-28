package operators

import (
	. "compiler/components"
	. "compiler/components/literals"
	"errors"
)

type DivOperator struct{}

func (op DivOperator) Apply(args ...Expression) (Literal, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments for Div Operator")
	}

	left := args[0]
	right := args[1]

	leftLiteral, _ := left.Eval()
	rightLiteral, _ := right.Eval()

	if !(leftLiteral.Kind() == IntKind && rightLiteral.Kind() == IntKind) {
		return nil, errors.New("invalid Operations")
	}

	if rightLiteral.(*IntLiteral).Value == 0 {
		return nil, errors.New("division by zero")
	}
	return &IntLiteral{leftLiteral.(*IntLiteral).Value / rightLiteral.(*IntLiteral).Value}, nil

}
