package literals

import (
	"fmt"
)

type BoolLiteral struct {
	Value    bool
	Position int
}

func (l *BoolLiteral) Node()           {}
func (l *BoolLiteral) Pos() int        { return l.Position }
func (l *BoolLiteral) ExpressionNode() {}
func (l *BoolLiteral) ConditionNode()  {}

func (l *BoolLiteral) AsBool() (bool, error) {
	return l.Value, nil
}

func (l *BoolLiteral) Kind() ValueKind {
	return BoolKind
}

func (l *BoolLiteral) Equals(l2 Literal) bool {
	o, ok := l2.(*BoolLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("%t", l.Value)
}
