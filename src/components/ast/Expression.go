package ast

import (
	. "compiler/src/components"
	"compiler/src/runtime/value"
)

type Expression interface {
	Node
	ExpressionNode()
	Type() (Type, error)
}

type BinaryExpr struct {
	Op         BinaryOperator[value.RuntimeValue]
	Left       Expression
	Right      Expression
	ReturnType Type
}

func (b *BinaryExpr) Node()           {}
func (b *BinaryExpr) Pos() int        { return b.Left.Pos() }
func (b *BinaryExpr) ExpressionNode() {}
func (b *BinaryExpr) Type() (Type, error) {
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

type UnaryExpr struct {
	Op         UnaryOperator[value.RuntimeValue]
	Arg        Expression
	ReturnType Type
}

func (u *UnaryExpr) Node()           {}
func (u *UnaryExpr) Pos() int        { return u.Arg.Pos() }
func (u *UnaryExpr) ExpressionNode() {}
func (u *UnaryExpr) Type() (Type, error) {
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

type TernaryExpr struct {
	Op         TernaryOperator[value.RuntimeValue]
	Condition  Condition
	Left       Expression
	Right      Expression
	ReturnType Type
}

func (t *TernaryExpr) Node()           {}
func (t *TernaryExpr) ExpressionNode() {}
func (t *TernaryExpr) Pos() int        { return t.Condition.Pos() }
func (t *TernaryExpr) Type() (Type, error) {
	if t.ReturnType != nil {
		return t.ReturnType, nil
	}
	leftType, leftErr := t.Left.Type()
	if leftErr != nil {
		return nil, leftErr
	}
	rightType, rightErr := t.Right.Type()
	if rightErr != nil {
		return nil, rightErr
	}
	rtnType, err := t.Op.GetType(leftType, rightType)
	if err != nil {
		return nil, err
	}
	t.ReturnType = rtnType
	return rtnType, nil
}
