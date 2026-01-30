package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
	"errors"
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
	return false, errors.New("Non-Implemented check function - Generic BinaryCondition")
}
