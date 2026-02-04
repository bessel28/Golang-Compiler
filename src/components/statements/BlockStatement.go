package statements

import "compiler/src/components/ast"

type BlockStatement struct {
	Statements []*ast.Statement
}
