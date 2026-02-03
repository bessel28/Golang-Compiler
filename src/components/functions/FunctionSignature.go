package functions

import (
	. "compiler/src/components/literals"
)

type FunctionSignature struct {
	Name       string
	Params     []ValueKind
	ReturnType *ValueKind
}
