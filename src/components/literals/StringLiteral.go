package literals

import (
	"compiler/src/components/ast"
)

type StringLiteral struct {
	Value string
}

func (l *StringLiteral) AsBool() (bool, error) {
	return l.Value != "", nil
}

func (l *StringLiteral) Equals(l2 ast.Literal) bool {
	o, ok := l2.(*StringLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l *StringLiteral) Eval() (ast.Literal, error) {
	return l, nil
}

func (l *StringLiteral) String() string {
	return l.Value
}

func (l *StringLiteral) Kind() ast.ValueKind {
	return ast.StringKind
}
