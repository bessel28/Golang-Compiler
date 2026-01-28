package components

import "fmt"

type ValueKind int

const (
	IntValue ValueKind = iota
	StringValue
)

type Stringer interface {
	String() string
}
type Literal struct {
	Kind ValueKind
	Int  int
	Str  string
}

type IntLiteral struct {
	Value int
}

type StringLiteral struct {
	Value string
}

func (l *Literal) Eval() (*Literal, error) {
	return l, nil
}
func (l *IntLiteral) Eval() (*Literal, error) {
	return &Literal{Kind: IntValue, Int: l.Value}, nil
}
func (l *StringLiteral) Eval() (*Literal, error) {
	return &Literal{Kind: StringValue, Str: l.Value}, nil
}

func (l *Literal) String() string {
	switch l.Kind {
	case IntValue:
		return fmt.Sprintf("%d", l.Int)
	case StringValue:
		return l.Str
	}
	return ""
}
func (l *IntLiteral) String() string {
	return string(rune(l.Value))
}
func (l *StringLiteral) String() string {
	return l.Value
}

func (l *Literal) Equals(l2 Literal) bool {
	if l.Kind != l2.Kind {
		return false
	}
	if l.Kind == IntValue {
		return l.Int == l2.Int
	}
	return l.Str == l2.Str
}

func NewLiteral(v interface{}) *Literal {
	switch v.(type) {
	case string:
		return &Literal{Kind: StringValue, Str: v.(string)}
	case int:
		return &Literal{Kind: IntValue, Int: v.(int)}
	default:
		return nil
	}
}
