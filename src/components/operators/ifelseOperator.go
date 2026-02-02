package operators

import (
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	"errors"
)

type TernaryOperator struct {
	Condition Condition
	Left      Expression
	Right     Expression
}

func (op *TernaryOperator) Eval() (Literal, error) {
	conditionLiteral, err := op.Condition.Eval()
	if err != nil {
		return nil, err
	}

	boolConditionLiteral, ok := conditionLiteral.(*BoolLiteral)
	if !ok {
		return nil, errors.New("non-bool literal evaluated from condition")
	}

	var literal Literal

	if boolConditionLiteral.Value {
		literal, err = op.Left.Eval()
	} else {
		literal, err = op.Right.Eval()
	}
	return literal, err
}
