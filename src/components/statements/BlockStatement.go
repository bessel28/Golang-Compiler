package statements

import "compiler/src/components/ast"

type BlockStatement struct {
	Statements []ast.Statement
	Position   int
}

func (e *BlockStatement) Node()          {}
func (e *BlockStatement) Pos() int       { return e.Position }
func (e *BlockStatement) StatementNode() {}
