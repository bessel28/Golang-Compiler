package conditions

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
)

type NotEqualCondition struct {
	*BinaryCondition
}

func (e *NotEqualCondition) Eval() (Literal, error) {
	check, err := e.Check()
	return &BoolLiteral{check}, err
}

func (e *NotEqualCondition) Check() (bool, error) {
	leftLiteral, err := e.Left.Eval()
	rightLiteral, err2 := e.Right.Eval()

	if err != nil || err2 != nil {
		return false, err
	}

	if leftLiteral.Equals(rightLiteral) {
		return false, nil
	}

	return true, nil
}
