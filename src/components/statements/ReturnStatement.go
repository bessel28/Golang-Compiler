package statements

import . "compiler/src/components/ast"

type ReturnStatement struct {
	Position   int
	ReturnExpr Expression
}

func (r *ReturnStatement) Node()          {}
func (r *ReturnStatement) Pos() int       { return r.Position }
func (r *ReturnStatement) StatementNode() {}
