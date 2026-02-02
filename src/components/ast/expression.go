package ast

type Expression interface {
	Eval() (Literal, error) //TODO Eval() needs to take an Environment param to use variables
}
