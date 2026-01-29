package conditions

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
)

type EqualCondition struct {
	*BinaryCondition
}

func (e *EqualCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *EqualCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()
	rightLiteral, err2 := e.Right.Eval()

	if err != nil || err2 != nil {
		return false, err
	}

	return leftLiteral.Equals(rightLiteral), nil
}
