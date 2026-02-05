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
func (b *BinaryCondition) Type() (Type, error) {
	if b.ReturnType != nil {
		return b.ReturnType, nil
	}
	leftType, leftErr := b.Left.Type()
	if leftErr != nil {
		return nil, leftErr
	}
	rightType, rightErr := b.Right.Type()
	if rightErr != nil {
		return nil, rightErr
	}
	rtnType, err := b.Op.GetType(leftType, rightType)
	if err != nil {
		return nil, err
	}
	b.ReturnType = rtnType
	return rtnType, nil
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
func (u *UnaryCondition) Type() (Type, error) {
	if u.ReturnType != nil {
		return u.ReturnType, nil
	}
	argType, argErr := u.Arg.Type()
	if argErr != nil {
		return nil, argErr
	}

	rtnType, err := u.Op.GetType(argType)
	if err != nil {
		return nil, err
	}
	u.ReturnType = rtnType
	return rtnType, nil
}
