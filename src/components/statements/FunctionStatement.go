package statements

import (
	. "compiler/src/components"
)

type FunctionStatement struct {
	Name       string
	Params     map[string]Type
	Block      *BlockStatement
	ReturnType Type
	Position   int
}

func (f *FunctionStatement) Node()          {}
func (f *FunctionStatement) Pos() int       { return f.Position }
func (f *FunctionStatement) StatementNode() {}
