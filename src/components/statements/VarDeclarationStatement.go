package statements

import (
	"compiler/src/components/ast"
)

type VariableDeclarationStatement struct {
	Name  string
	Value ast.Expression
}
