package operators

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	"errors"
)

type SubOperator struct{}

func (op SubOperator) Apply(left *Literal, right *Literal) (Literal, error) {
	if !((*left).Kind() == IntKind && (*right).Kind() == IntKind) {
		return nil, errors.New("invalid Operations")
	}
	return &IntLiteral{(*left).(*IntLiteral).Value - (*right).(*IntLiteral).Value, (*left).Pos()}, nil
}

func (op SubOperator) GetType(left Type, right Type) (Type, error) {
	switch {
	case left.Compare(&IntType{}) && right.Compare(&IntType{}):
		return &IntType{}, nil
	}
	return nil, errors.New("Invalid Parameter types for Div Operator")
}
