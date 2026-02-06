package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	. "compiler/src/runtime/value"
	"errors"
)

type GreaterThanCondition struct{}

func (op GreaterThanCondition) GetType(left Type, right Type) (Type, error) {
	if left.Compare(&BoolType{}) || right.Compare(&BoolType{}) {
		return nil, errors.New("Cannot compare bools using > operator")
	}
	if !left.Compare(right) {
		return nil, errors.New("Cannot compare differnt types using > operator")
	}

	return &BoolType{}, nil
}

func (op GreaterThanCondition) Apply(left RuntimeValue, right RuntimeValue) (*BoolValue, error) {
	if left.Kind() == BoolKind || right.Kind() == BoolKind {
		return nil, errors.New("Invalid comparison")
	}

	if left.Kind() != right.Kind() {
		return nil, errors.New("Invalid comparison")
	}

	switch left.Kind() {
	case IntKind:
		return &BoolValue{left.(*IntValue).Value().(int) > right.(*IntValue).Value().(int)}, nil
	case StringKind:
		return &BoolValue{left.(*StringValue).Value().(string) > right.(*StringValue).Value().(string)}, nil
	default:
		return nil, errors.New("unhandled default case")
	}
}
