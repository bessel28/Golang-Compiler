package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
)

type NotCondition struct{}

func (e NotCondition) GetType(Type) (Type, error) {
	return &BoolType{}, nil
}

func (e NotCondition) Apply(arg *Literal) (*BoolLiteral, error) {
	argBool, _ := (*arg).AsBool()
	return &BoolLiteral{!argBool, (*arg).Pos()}, nil
}
