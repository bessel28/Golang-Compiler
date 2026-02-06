package operators

import (
	"compiler/src/components"
	. "compiler/src/runtime/value"
	"errors"
)

type IfElseOperator struct{}

func (op IfElseOperator) Apply(condition *BoolValue, trueVal RuntimeValue, falseVal RuntimeValue) (RuntimeValue, error) {
	if condition.Val {
		return trueVal, nil
	} else {
		return falseVal, nil
	}
}

func (op IfElseOperator) GetType(trueType components.Type, falseType components.Type) (components.Type, error) {
	if trueType == falseType {
		return trueType, nil
	}
	return nil, errors.New("True and False Types did not match")
}
