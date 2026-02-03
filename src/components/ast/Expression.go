package ast

import "compiler/src/components/literals"

type Expression interface {
	Node
	ExpressionNode()
}

type BinaryExpr struct {
	Op    BinaryOperator[literals.Literal]
	Left  Expression
	Right Expression
}

func (u *BinaryExpr) Node()           {}
func (u *BinaryExpr) Pos() int        { return u.Left.Pos() }
func (u *BinaryExpr) ExpressionNode() {}

type UnaryExpr struct {
	Op  UnaryOperator[literals.Literal]
	Arg Expression
}

func (u *UnaryExpr) Node()           {}
func (u *UnaryExpr) Pos() int        { return u.Arg.Pos() }
func (u *UnaryExpr) ExpressionNode() {}

type TernaryExpr struct {
	Op        TernaryOperator[literals.Literal]
	Condition Condition
	Left      Expression
	Right     Expression
}

func (t *TernaryExpr) Node()           {}
func (t *TernaryExpr) ExpressionNode() {}
func (t *TernaryExpr) Pos() int        { return t.Condition.Pos() }
