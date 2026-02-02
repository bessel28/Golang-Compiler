package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
)

type OrCondition struct {
	*BinaryCondition
}

func (e *OrCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *OrCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()
	rightLiteral, err2 := e.Right.Eval()

	if err != nil || err2 != nil {
		return false, err
	}

	leftBool, _ := leftLiteral.AsBool()
	rightBool, _ := rightLiteral.AsBool()

	return leftBool || rightBool, nil
}
