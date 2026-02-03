package conditions

import (
	. "compiler/src/components/literals"
	"errors"
)

type GreaterThanCondition struct{}

func (op GreaterThanCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	if (*left).Kind() == BoolKind || (*right).Kind() == BoolKind {
		return &BoolLiteral{false, (*left).Pos()}, errors.New("Invalid comparison")
	}

	if (*left).Kind() != (*right).Kind() {
		return &BoolLiteral{false, (*left).Pos()}, errors.New("Invalid comparison")
	}

	switch (*left).Kind() {
	case IntKind:
		return &BoolLiteral{(*left).(*IntLiteral).Value > (*right).(*IntLiteral).Value, (*left).Pos()}, nil
	case StringKind:
		return &BoolLiteral{(*left).(*StringLiteral).Value > (*right).(*StringLiteral).Value, (*left).Pos()}, nil
	default:
		return &BoolLiteral{false, (*left).Pos()}, errors.New("unhandled default case")
	}
}
