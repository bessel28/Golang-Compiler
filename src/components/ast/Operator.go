package ast

import (
	. "compiler/src/components"
	. "compiler/src/components/literals"
)

type TernaryOperator[T Literal] interface {
	Apply(condition *BoolLiteral, trueVal *Literal, falseVal *Literal) (T, error)
	GetType(trueType Type, falseType Type) (Type, error)
}
type BinaryOperator[T Literal] interface {
	Apply(left *Literal, right *Literal) (T, error)
	GetType(left Type, right Type) (Type, error)
}

type UnaryOperator[T Literal] interface {
	Apply(arg *Literal) (T, error)
	GetType(arg Type) (Type, error)
}
