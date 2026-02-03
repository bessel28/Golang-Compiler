package literals

type StringLiteral struct {
	Value    string
	Position int
}

func (l StringLiteral) Node()           {}
func (l StringLiteral) Pos() int        { return l.Position }
func (l StringLiteral) ExpressionNode() {}
func (l StringLiteral) ConditionNode()  {}

func (l StringLiteral) AsBool() (bool, error) {
	return l.Value != "", nil
}

func (l StringLiteral) Equals(l2 Literal) bool {
	o, ok := l2.(*StringLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l StringLiteral) String() string {
	return l.Value
}

func (l StringLiteral) Kind() ValueKind {
	return StringKind
}
