package components

import . "compiler/components/literals"

type Condition interface {
	Expression
	check(args ...Expression) (*BoolLiteral, error)
}

type BinaryCondition struct {
	Left, Right Expression
}

func (c *BinaryCondition) Eval() (Literal, error) {
	return c.check(c.Left, c.Right)
}

func (c *BinaryCondition) check(args ...Expression) (*BoolLiteral, error) {
	panic("Non-Implemented check function - Generic BinaryCondition")
}
