package statements

import (
	. "compiler/src/components/ast"
	"compiler/src/components/literals"
)

type FunctionStatement struct {
	Name       string
	Params     []Expression
	Block      BlockStatement
	ReturnType literals.ValueKind
}
