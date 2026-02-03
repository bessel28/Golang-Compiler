package ast

import "compiler/src/components/literals"

type Condition interface {
	ConditionNode()
	Expression
}

type BinaryCondition struct {
	Op    BinaryOperator[*literals.BoolLiteral]
	Left  Expression
	Right Expression
}

func (b *BinaryCondition) Node()           {}
func (b *BinaryCondition) Pos() int        { return b.Left.Pos() }
func (b *BinaryCondition) ExpressionNode() {}
func (b *BinaryCondition) ConditionNode()  {}

type UnaryCondition struct {
	Op  UnaryOperator[*literals.BoolLiteral]
	Arg Expression
}

func (u *UnaryCondition) Node()           {}
func (u *UnaryCondition) Pos() int        { return u.Arg.Pos() }
func (u *UnaryCondition) ExpressionNode() {}
func (u *UnaryCondition) ConditionNode()  {}
