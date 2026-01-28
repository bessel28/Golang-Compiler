package literals

import "fmt"

type BoolLiteral struct {
	Value bool
}

func (l *BoolLiteral) Kind() ValueKind {
	return BoolKind
}

func (l *BoolLiteral) Equals(l2 Literal) BoolLiteral {
	o, ok := l2.(*BoolLiteral)
	if !ok {
		return BoolLiteral{false}
	}
	return BoolLiteral{l.Value == o.Value}
}

func (l *BoolLiteral) Eval() (Literal, error) {
	return l, nil
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("%t", l.Value)
}
