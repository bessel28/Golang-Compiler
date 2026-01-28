package literals

import "fmt"

type IntLiteral struct {
	Value int
}

func (l *IntLiteral) Equals(l2 Literal) *BoolLiteral {
	o, ok := l2.(*IntLiteral)
	if !ok {
		return &BoolLiteral{false}
	}
	return &BoolLiteral{l.Value == o.Value}
}

func (l *IntLiteral) Eval() (Literal, error) {
	return l, nil
}

func (l *IntLiteral) String() string {
	return fmt.Sprintf("%d", l.Value)
}

func (l *IntLiteral) Kind() ValueKind {
	return IntKind
}
