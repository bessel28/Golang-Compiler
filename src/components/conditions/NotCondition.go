package conditions

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type NotCondition struct{}

func (e NotCondition) GetType(Type) (Type, error) {
	return &BoolType{}, nil
}

func (e NotCondition) Apply(arg RuntimeValue) (*BoolValue, error) {
	argBool := arg.BoolVal()
	return &BoolValue{!argBool.Val}, nil
}
