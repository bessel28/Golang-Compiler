package literals

type StringLiteral struct {
	Value string
}

func (l *StringLiteral) Equals(l2 Literal) BoolLiteral {
	o, ok := l2.(*StringLiteral)
	if !ok {
		return BoolLiteral{false}
	}
	return BoolLiteral{l.Value == o.Value}
}

func (l *StringLiteral) Eval() (Literal, error) {
	return l, nil
}

func (l *StringLiteral) String() string {
	return l.Value
}

func (l *StringLiteral) Kind() ValueKind {
	return StringKind
}
