package conditions

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type EqualCondition struct{}

func (op EqualCondition) GetType(Type, Type) (Type, error) {
	return &BoolType{}, nil
}

func (op EqualCondition) Apply(left RuntimeValue, right RuntimeValue) (*BoolValue, error) {
	return &BoolValue{left.Value() == right.Value()}, nil
}
