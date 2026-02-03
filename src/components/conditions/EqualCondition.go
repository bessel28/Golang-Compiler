package conditions

import (
	. "compiler/src/components/literals"
)

type EqualCondition struct{}

func (op EqualCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	return &BoolLiteral{(*left).Equals(*right), (*left).Pos()}, nil
}
