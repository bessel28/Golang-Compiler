package conditions

import (
	. "compiler/src/components/literals"
)

type NotEqualCondition struct{}

func (op NotEqualCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	return &BoolLiteral{!(*left).Equals(*right), (*left).Pos()}, nil
}
