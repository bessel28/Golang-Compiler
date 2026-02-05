package operators

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	"errors"
)

type DivOperator struct{}

func (op DivOperator) Apply(left *Literal, right *Literal) (Literal, error) {
	if !((*left).Kind() == IntKind && (*right).Kind() == IntKind) {
		return nil, errors.New("invalid Operations")
	}

	if (*right).(*IntLiteral).Value == 0 {
		return nil, errors.New("division by zero")
	}
	return &IntLiteral{(*left).(*IntLiteral).Value / (*right).(*IntLiteral).Value, (*left).Pos()}, nil

}

func (op DivOperator) GetType(left Type, right Type) (Type, error) {
	switch {
	case left.Compare(&IntType{}) && right.Compare(&IntType{}):
		return &IntType{}, nil
	}
	return nil, errors.New("Invalid Parameter types for Div Operator")
}
