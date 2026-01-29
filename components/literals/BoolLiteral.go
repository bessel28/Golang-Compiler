package literals

import (
	"compiler/components/ast"
	"fmt"
)

type BoolLiteral struct {
	Value bool
}

func (l *BoolLiteral) AsBool() (bool, error) {
	return l.Value, nil
}

func (l *BoolLiteral) Kind() ast.ValueKind {
	return ast.BoolKind
}

func (l *BoolLiteral) Equals(l2 ast.Literal) bool {
	o, ok := l2.(*BoolLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l *BoolLiteral) Eval() (ast.Literal, error) {
	return l, nil
}

func (l *BoolLiteral) String() string {
	return fmt.Sprintf("%t", l.Value)
}

func (l *BoolLiteral) Check() (bool, error) {
	return l.Value, nil
}
