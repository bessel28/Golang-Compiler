package conditions

import (
	"compiler/src/components/ast"
	. "compiler/src/components/literals"
)

type AndCondition struct {
	ast.Condition
}

func (op AndCondition) Apply(left *Literal, right *Literal) (*BoolLiteral, error) {
	leftBool, _ := (*left).AsBool()
	rightBool, _ := (*right).AsBool()

	return &BoolLiteral{leftBool && rightBool, (*left).Pos()}, nil
}
