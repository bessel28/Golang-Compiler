package statements

import (
	"compiler/src/components/ast"
)

type VariableAssignmentStatement struct {
	Name     string
	Value    ast.Expression
	Position int
}

func (v *VariableAssignmentStatement) Node()          {}
func (v *VariableAssignmentStatement) Pos() int       { return v.Position }
func (v *VariableAssignmentStatement) statementNode() {}
