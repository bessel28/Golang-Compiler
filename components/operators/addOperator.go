package operators

import (
	. "compiler/components"
	"errors"
	"fmt"
)

type AddOperator struct{}

func (op AddOperator) Apply(args ...Expression) (*Literal, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments for Add Operator")
	}

	left := args[0]
	right := args[1]

	leftLiteral, _ := left.Eval()
	rightLiteral, _ := right.Eval()

	switch {
	case leftLiteral.Kind == IntValue && rightLiteral.Kind == IntValue:
		return &Literal{Kind: IntValue, Int: leftLiteral.Int + rightLiteral.Int}, nil
	case leftLiteral.Kind == StringValue && rightLiteral.Kind == StringValue:
		return &Literal{Kind: StringValue, Str: leftLiteral.Str + rightLiteral.Str}, nil
	case leftLiteral.Kind == StringValue && rightLiteral.Kind == IntValue:
		return &Literal{Kind: StringValue, Str: leftLiteral.Str + fmt.Sprintf("%d", rightLiteral.Int)}, nil
	case leftLiteral.Kind == IntValue && rightLiteral.Kind == StringValue:
		return &Literal{Kind: StringValue, Str: fmt.Sprintf("%d", leftLiteral.Int) + rightLiteral.Str}, nil
	}

	return nil, errors.New("invalid Operations")
}
