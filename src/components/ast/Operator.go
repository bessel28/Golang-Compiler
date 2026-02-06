package ast

import (
	. "compiler/src/components"
	. "compiler/src/runtime/value"
)

type TernaryOperator[T RuntimeValue] interface {
	Apply(condition *BoolValue, trueVal, falseVal RuntimeValue) (T, error)
	GetType(trueType Type, falseType Type) (Type, error)
}
type BinaryOperator[T RuntimeValue] interface {
	Apply(left RuntimeValue, right RuntimeValue) (T, error)
	GetType(left Type, right Type) (Type, error)
}

type UnaryOperator[T RuntimeValue] interface {
	Apply(arg RuntimeValue) (T, error)
	GetType(arg Type) (Type, error)
}
