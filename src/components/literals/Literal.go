package literals

type Stringer interface {
	String() string
}

type Literal interface {
	Kind() ValueKind
	String() string
	Equals(l Literal) bool
	AsBool() (bool, error)
	ExpressionNode()
	Node()
	Pos() int
}
