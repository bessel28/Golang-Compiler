package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
)

type EqualCondition struct{}

func (op EqualCondition) GetType(Type, Type) (Type, error) {
	return &BoolType{}, nil
}

func (op EqualCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	return &BoolLiteral{(*left).Equals(*right), (*left).Pos()}, nil
}
