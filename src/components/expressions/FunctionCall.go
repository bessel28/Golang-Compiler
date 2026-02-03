package expressions

import (
	. "compiler/src/components/ast"
)

type CallExpr struct {
	identifier string
	params     []Expression
	Expression
}
