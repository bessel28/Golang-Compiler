package ast

type Expression interface {
	Eval() (Literal, error)
}
