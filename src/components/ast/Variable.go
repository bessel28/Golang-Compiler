package ast

import "compiler/src/components/literals"

type Variable struct {
	name  string
	value literals.Literal
}
