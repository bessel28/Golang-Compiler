package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
)

type AndCondition struct{}

func (op AndCondition) GetType(Type, Type) (Type, error) {
	return &BoolType{}, nil
}

func (op AndCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	leftBool, _ := (*left).AsBool()
	rightBool, _ := (*right).AsBool()

	return &BoolLiteral{leftBool && rightBool, (*left).Pos()}, nil
}
