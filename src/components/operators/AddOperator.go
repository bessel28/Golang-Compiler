package operators

import (
	. "compiler/src/components/literals"
)

type AddOperator struct{}

func (op AddOperator) Apply(left *Literal, right *Literal) (Literal, error) {
	switch {
	case (*left).Kind() == IntKind && (*right).Kind() == IntKind:
		return &IntLiteral{(*left).(*IntLiteral).Value + (*right).(*IntLiteral).Value, (*left).Pos()}, nil
	default:
		return &StringLiteral{(*left).String() + (*right).String(), (*left).Pos()}, nil
	}
}
