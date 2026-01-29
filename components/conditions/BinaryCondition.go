package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
)

type BinaryCondition struct {
	Left, Right Expression
}

func (c *BinaryCondition) Eval() (Literal, error) {
	check, err := c.Check()
	if err != nil {
		return nil, err
	}
	return &BoolLiteral{check}, nil
}

func (c *BinaryCondition) Check() (bool, error) {
	panic("Non-Implemented check function - Generic BinaryCondition")
}
