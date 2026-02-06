package conditions

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type AndCondition struct{}

func (op AndCondition) GetType(Type, Type) (Type, error) {
	return &BoolType{}, nil
}

func (op AndCondition) Apply(left RuntimeValue, right RuntimeValue) (*BoolValue, error) {
	leftBool := left.BoolVal()
	rightBool := right.BoolVal()

	return &BoolValue{leftBool.Val && rightBool.Val}, nil
}
