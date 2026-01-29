package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
)

type NotCondition struct {
	*UnaryCondition
}

func (e *NotCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *NotCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()

	if err != nil {
		return false, err
	}

	leftBool, _ := leftLiteral.AsBool()

	return !leftBool, nil
}
