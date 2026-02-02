package ast

type ValueKind int

const (
	IntKind ValueKind = iota
	StringKind
	BoolKind
)

type Stringer interface {
	String() string
}

type Literal interface {
	Kind() ValueKind
	String() string
	Equals(l Literal) bool
	Eval() (Literal, error)
	AsBool() (bool, error)
}
