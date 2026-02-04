package ast

import (
	. "compiler/src/components"
	"compiler/src/components/literals"
)

type Expression interface {
	Node
	ExpressionNode()
	Type() Type
}

type BinaryExpr struct {
	Op         BinaryOperator[literals.Literal]
	Left       Expression
	Right      Expression
	ReturnType Type
}

func (b *BinaryExpr) Node()           {}
func (b *BinaryExpr) Pos() int        { return b.Left.Pos() }
func (b *BinaryExpr) ExpressionNode() {}
func (b *BinaryExpr) Type() Type      { return b.ReturnType }

type UnaryExpr struct {
	Op         UnaryOperator[literals.Literal]
	Arg        Expression
	ReturnType Type
}

func (u *UnaryExpr) Node()           {}
func (u *UnaryExpr) Pos() int        { return u.Arg.Pos() }
func (u *UnaryExpr) ExpressionNode() {}
func (u *UnaryExpr) Type() Type      { return u.ReturnType }

type TernaryExpr struct {
	Op         TernaryOperator[literals.Literal]
	Condition  Condition
	Left       Expression
	Right      Expression
	ReturnType Type
}

func (t *TernaryExpr) Node()           {}
func (t *TernaryExpr) ExpressionNode() {}
func (t *TernaryExpr) Pos() int        { return t.Condition.Pos() }
func (t *TernaryExpr) Type() Type      { return t.ReturnType }
