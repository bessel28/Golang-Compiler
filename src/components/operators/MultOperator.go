package operators

import (
	. "compiler/src/components/literals"
	"errors"
)

type MultOperator struct{}

func (op MultOperator) Apply(left *Literal, right *Literal) (Literal, error) {
	if !((*left).Kind() == IntKind && (*right).Kind() == IntKind) {
		return nil, errors.New("invalid Operations")
	}
	return &IntLiteral{(*left).(*IntLiteral).Value * (*right).(*IntLiteral).Value, (*left).Pos()}, nil
}
