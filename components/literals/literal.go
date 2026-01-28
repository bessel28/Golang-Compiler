package literals

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
	Equals(l Literal) *BoolLiteral
	Eval() (Literal, error)
}

func NewLiteral(v interface{}) Literal {
	switch v.(type) {
	case string:
		return &StringLiteral{v.(string)}
	case int:
		return &IntLiteral{v.(int)}
	case bool:
		return &BoolLiteral{v.(bool)}
	default:
		return nil
	}
}
