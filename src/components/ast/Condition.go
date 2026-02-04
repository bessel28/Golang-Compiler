package ast

import (
	. "compiler/src/components"
	"compiler/src/components/literals"
)

type Condition interface {
	ConditionNode()
	Expression
}

type BinaryCondition struct {
	Op         BinaryOperator[*literals.BoolLiteral]
	Left       Expression
	Right      Expression
	ReturnType Type
}

func (b *BinaryCondition) Node()           {}
func (b *BinaryCondition) Pos() int        { return b.Left.Pos() }
func (b *BinaryCondition) ExpressionNode() {}
func (b *BinaryCondition) ConditionNode()  {}
func (b *BinaryCondition) Type() Type {
	//TODO implement me
	return b.ReturnType
}

type UnaryCondition struct {
	Op         UnaryOperator[*literals.BoolLiteral]
	Arg        Expression
	ReturnType Type
}

func (u *UnaryCondition) Node()           {}
func (u *UnaryCondition) Pos() int        { return u.Arg.Pos() }
func (u *UnaryCondition) ExpressionNode() {}
func (u *UnaryCondition) ConditionNode()  {}
func (u *UnaryCondition) Type() Type      { return u.ReturnType }
