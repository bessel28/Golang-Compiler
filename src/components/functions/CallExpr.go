package functions

import (
	. "compiler/src/components"
	"compiler/src/components/ast"
)

type CallExpr struct {
	Callee   ast.Expression
	Position int
	Params   []ast.Expression
	*FunctionSignature
}

func (c *CallExpr) Type() (Type, error) {
	return c.ReturnType, nil
}

func (c *CallExpr) Node()           {}
func (c *CallExpr) Pos() int        { return c.Position }
func (c *CallExpr) ExpressionNode() {}
