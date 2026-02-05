package conditions

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
	"errors"
)

type LessThanCondition struct{}

func (op LessThanCondition) GetType(left Type, right Type) (Type, error) {
	if left.Compare(&BoolType{}) || right.Compare(&BoolType{}) {
		return nil, errors.New("Cannot compare bools using < operator")
	}
	if !left.Compare(right) {
		return nil, errors.New("Cannot compare differnt types using < operator")
	}
	return &BoolType{}, nil
}

func (op LessThanCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	if (*left).Kind() == BoolKind || (*right).Kind() == BoolKind {
		return nil, errors.New("Invalid comparison")
	}

	if (*left).Kind() != (*right).Kind() {
		return nil, errors.New("Invalid comparison")
	}

	switch (*left).Kind() {
	case IntKind:
		return &BoolLiteral{(*left).(*IntLiteral).Value < (*right).(*IntLiteral).Value, (*left).Pos()}, nil
	case StringKind:
		return &BoolLiteral{(*left).(*StringLiteral).Value < (*right).(*StringLiteral).Value, (*left).Pos()}, nil
	default:
		return nil, errors.New("unhandled default case")
	}
}
