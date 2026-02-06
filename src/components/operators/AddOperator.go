package operators

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	. "compiler/src/runtime/value"
)

type AddOperator struct{}

func (op AddOperator) Apply(left RuntimeValue, right RuntimeValue) (RuntimeValue, error) {
	switch {
	case left.Kind() == IntKind && right.Kind() == IntKind:
		return &IntValue{left.(*IntValue).Val + right.(*IntValue).Val}, nil
	default:
		return &StringValue{left.String().Val + right.String().Val}, nil
	}
}

func (op AddOperator) GetType(left Type, right Type) (Type, error) {
	switch {
	case left.Compare(&IntType{}) && right.Compare(&IntType{}):
		return &IntType{}, nil
	}
	return &StringType{}, nil
}
