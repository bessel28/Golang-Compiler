package literals

import (
	. "compiler/src/components"
	"fmt"
)

type IntLiteral struct {
	Value    int
	Position int
}

func (l *IntLiteral) Type() Type {
	//TODO implement me
	panic("implement me")
}

func (l *IntLiteral) Node()           {}
func (l *IntLiteral) Pos() int        { return l.Position }
func (l *IntLiteral) ExpressionNode() {}
func (l *IntLiteral) ConditionNode()  {}

func (l *IntLiteral) AsBool() (bool, error) {
	return l.Value != 0, nil
}

func (l *IntLiteral) Equals(l2 Literal) bool {
	o, ok := l2.(*IntLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l *IntLiteral) String() string {
	return fmt.Sprintf("%d", l.Value)
}

func (l *IntLiteral) Kind() ValueKind {
	return IntKind
}
