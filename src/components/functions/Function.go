package functions

import (
	"compiler/src/components/ast"
	"compiler/src/components/statements"
)

type Function struct {
	Name   string
	Params []ast.Expression
	Body   *statements.BlockStatement
}
