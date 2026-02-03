package operators

import . "compiler/src/components/literals"

type IfElseOperator struct{}

func (op IfElseOperator) Apply(condition *BoolLiteral, trueVal *Literal, falseVal *Literal) (Literal, error) {
	if condition.Value {
		return *trueVal, nil
	} else {
		return *falseVal, nil
	}
}
