package functions

import (
	. "compiler/src/components"
	"compiler/src/components/ast"
)

type CallExpr struct {
	Callee   ast.Expression
	Position int
	*FunctionSignature
}

func (c *CallExpr) Type() Type {
	return c.ReturnType
}

func (c *CallExpr) Node()           {}
func (c *CallExpr) Pos() int        { return c.Position }
func (c *CallExpr) ExpressionNode() {}
