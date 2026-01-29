package operators

import (
	. "compiler/components/ast"
	. "compiler/components/literals"
	"errors"
	"fmt"
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
	fmt.Println(op)
	return literal, err
}
