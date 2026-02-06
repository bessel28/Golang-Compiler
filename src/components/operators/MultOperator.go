package operators

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	. "compiler/src/runtime/value"
	"errors"
)

type MultOperator struct{}

func (op MultOperator) Apply(left RuntimeValue, right RuntimeValue) (RuntimeValue, error) {
	if !(left.Kind() == IntKind && right.Kind() == IntKind) {
		return nil, errors.New("invalid Operations")
	}
	return &IntValue{left.(*IntValue).Val * right.(*IntValue).Val}, nil
}

func (op MultOperator) GetType(left Type, right Type) (Type, error) {
	switch {
	case left.Compare(&IntType{}) && right.Compare(&IntType{}):
		return &IntType{}, nil
	}
	return nil, errors.New("Invalid Parameter types for Div Operator")
}
