package components

import "compiler/components/literals"

type Expression interface {
	Eval() (literals.Literal, error)
}
