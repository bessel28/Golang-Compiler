package ast

import (
	. "compiler/src/components/literals"
)

type TernaryOperator[T Literal] interface {
	Apply(condition *BoolLiteral, trueVal *Literal, falseVal *Literal) (T, error)
}
type BinaryOperator[T Literal] interface {
	Apply(left *Literal, right *Literal) (T, error)
}

type UnaryOperator[T Literal] interface {
	Apply(arg *Literal) (T, error)
}
