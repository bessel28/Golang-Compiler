package ast

import (
	"compiler/src/components/literals"
)

type VariableExpr struct {
	Name     string
	Type     literals.ValueKind
	Position int
}

func (v *VariableExpr) Node()           {}
func (v *VariableExpr) Pos() int        { return v.Position }
func (v *VariableExpr) ExpressionNode() {}
