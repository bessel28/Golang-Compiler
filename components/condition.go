package components

import "compiler/components/literals"

type Condition interface {
	Eval() (*literals.BoolLiteral, error)
}
