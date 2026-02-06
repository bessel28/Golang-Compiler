package conditions

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type NotEqualCondition struct{}

func (op NotEqualCondition) GetType(left Type, right Type) (Type, error) {
	return &BoolType{}, nil
}

func (op NotEqualCondition) Apply(left RuntimeValue, right RuntimeValue) (*BoolValue, error) {
	return &BoolValue{left.Value() != right.Value()}, nil
}
