package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
)

type OrCondition struct{}

func (op OrCondition) GetType(left Type, right Type) (Type, error) {
	return &BoolType{}, nil
}

func (op OrCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	leftBool, _ := (*left).AsBool()
	rightBool, _ := (*right).AsBool()

	return &BoolLiteral{leftBool || rightBool, (*left).Pos()}, nil
}
