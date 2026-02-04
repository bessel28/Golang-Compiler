package statements

import (
	. "compiler/src/components/ast"
)

type ExpressionStatement struct {
	Expression Expression
	Position   int
}

func (e *ExpressionStatement) Node()          {}
func (e *ExpressionStatement) Pos() int       { return e.Position }
func (e *ExpressionStatement) statementNode() {}
