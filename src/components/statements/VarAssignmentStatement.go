package statements

import (
	"compiler/src/components/ast"
)

type VariableAssignmentStatement struct {
	Name  string
	Value ast.Expression
}
