package components

import "compiler/components/literals"

type BinaryOperator struct {
	Left  Expression
	Right Expression
	Op    Operator
}

type UnaryOperator struct {
	Left Expression
	Op   Operator
}

type TernaryOperator struct {
	Condition Expression
	Left      Expression
	Right     Expression
	Op        Operator
}

type Operator interface {
	Apply(args ...Expression) (literals.Literal, error)
}

func (op UnaryOperator) Eval() (literals.Literal, error) {
	expr, err := op.Op.Apply(op.Left)
	return expr, err
}
func (op BinaryOperator) Eval() (literals.Literal, error) {
	expr, err := op.Op.Apply(op.Left, op.Right)
	return expr, err
}
func (op TernaryOperator) Eval() (literals.Literal, error) {
	expr, err := op.Op.Apply(op.Condition, op.Left, op.Right)
	return expr, err
}
