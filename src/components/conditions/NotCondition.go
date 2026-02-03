package conditions

import (
	. "compiler/src/components/literals"
)

type NotCondition struct{}

func (e NotCondition) Apply(arg *Literal) (*BoolLiteral, error) {
	argBool, _ := (*arg).AsBool()
	return &BoolLiteral{!argBool, (*arg).Pos()}, nil
}
