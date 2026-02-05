package functions

import (
	. "compiler/src/components"
)

type FunctionSignature struct {
	Name       string
	Params     []*Parameter
	ReturnType Type
}
