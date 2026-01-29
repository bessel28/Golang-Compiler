package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
)

type AndCondition struct {
	*BinaryCondition
}

func (e *AndCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *AndCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()
	rightLiteral, err2 := e.Right.Eval()

	if err != nil || err2 != nil {
		return false, err
	}

	leftBool, _ := leftLiteral.AsBool()
	rightBool, _ := rightLiteral.AsBool()

	return leftBool && rightBool, nil
}
