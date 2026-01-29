package operators

import (
	. "compiler/components/ast"
)

type BinaryOperator struct {
	Left  Expression
	Right Expression
	Op    Operator
}

type UnaryOperator struct {
	Left Expression
	Op   Operator
}

type Operator interface {
	Apply(args ...Expression) (Literal, error)
}

func (op UnaryOperator) Eval() (Literal, error) {
	expr, err := op.Op.Apply(op.Left)
	return expr, err
}
func (op BinaryOperator) Eval() (Literal, error) {
	expr, err := op.Op.Apply(op.Left, op.Right)
	return expr, err
}
