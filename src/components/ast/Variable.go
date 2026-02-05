package ast

import (
	. "compiler/src/components"
)

type VariableExpr struct {
	Name       string
	ReturnType Type
	Position   int
}

func (v *VariableExpr) Node()               {}
func (v *VariableExpr) Pos() int            { return v.Position }
func (v *VariableExpr) ExpressionNode()     {}
func (v *VariableExpr) Type() (Type, error) { return v.ReturnType, nil }
