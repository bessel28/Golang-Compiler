package functions

import "compiler/src/components/ast"

type Callable interface {
	ast.Expression
	Callable()
	GetIdentifier() string
	SetFunctionSignature(*FunctionSignature)
	GetFunctionSignature() *FunctionSignature
}
