package statements

import (
	. "compiler/src/components/ast"
	"compiler/src/components/literals"
)

type FunctionStatement struct {
	Name       string
	Params     map[string]*Expression
	Block      *BlockStatement
	ReturnType literals.ValueKind
	Position   int
}

func (f *FunctionStatement) Node()          {}
func (f *FunctionStatement) Pos() int       { return f.Position }
func (f *FunctionStatement) statementNode() {}
