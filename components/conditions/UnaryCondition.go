package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
)

type UnaryCondition struct {
	Left Expression
}

func (c *UnaryCondition) Eval() (Literal, error) {
	check, err := c.Check()
	if err != nil {
		return nil, err
	}
	return &BoolLiteral{check}, nil
}

func (c *UnaryCondition) Check() (bool, error) {
	panic("Non-Implemented check function - Generic BinaryCondition")
}
