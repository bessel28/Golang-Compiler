package functions

import (
	"compiler/src/components/ast"
)

type Callable interface {
	ast.Expression
	Callable()
	SetFunctionSignature(*FunctionSignature)
	GetFunctionSignature() *FunctionSignature
}
