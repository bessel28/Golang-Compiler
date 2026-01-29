package literals

import (
	"compiler/components/ast"
	"fmt"
)

type IntLiteral struct {
	Value int
}

func (l *IntLiteral) AsBool() (bool, error) {
	return l.Value != 0, nil
}

func (l *IntLiteral) Equals(l2 ast.Literal) bool {
	o, ok := l2.(*IntLiteral)
	if !ok {
		return false
	}
	return l.Value == o.Value
}

func (l *IntLiteral) Eval() (ast.Literal, error) {
	return l, nil
}

func (l *IntLiteral) String() string {
	return fmt.Sprintf("%d", l.Value)
}

func (l *IntLiteral) Kind() ast.ValueKind {
	return ast.IntKind
}
