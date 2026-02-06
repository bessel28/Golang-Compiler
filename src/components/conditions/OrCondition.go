package conditions

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type OrCondition struct{}

func (op OrCondition) GetType(left Type, right Type) (Type, error) {
	return &BoolType{}, nil
}

func (op OrCondition) Apply(left RuntimeValue, right RuntimeValue) (*BoolValue, error) {
	leftBool := left.BoolVal()
	rightBool := right.BoolVal()

	return &BoolValue{leftBool.Val || rightBool.Val}, nil
}
