package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
	"errors"
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
	return false, errors.New("UnaryCondition Check function - Generic UnaryCondition check used")
}
