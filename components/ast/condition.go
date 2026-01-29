package ast

type Condition interface {
	Expression
	Check() (bool, error)
}
