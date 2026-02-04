package statements

import (
	"compiler/src/components/ast"
)

type VariableDeclarationStatement struct {
	Name     string
	Value    ast.Expression
	Position int
}

func (v *VariableDeclarationStatement) Node()          {}
func (v *VariableDeclarationStatement) Pos() int       { return v.Position }
func (v *VariableDeclarationStatement) statementNode() {}
