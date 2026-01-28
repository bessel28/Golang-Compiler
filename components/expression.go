package components

type Expression interface {
	Eval() (*Literal, error)
}
