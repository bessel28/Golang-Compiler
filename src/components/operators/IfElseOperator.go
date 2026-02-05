package operators

import (
	"compiler/src/components"
	. "compiler/src/components/literals"
	"errors"
)

type IfElseOperator struct{}

func (op IfElseOperator) Apply(condition *BoolLiteral, trueVal *Literal, falseVal *Literal) (Literal, error) {
	if condition.Value {
		return *trueVal, nil
	} else {
		return *falseVal, nil
	}
}

func (op IfElseOperator) GetType(trueType components.Type, falseType components.Type) (components.Type, error) {
	if trueType == falseType {
		return trueType, nil
	}
	return nil, errors.New("True and False Types did not match")
}
