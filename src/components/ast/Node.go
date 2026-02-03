package ast

type Node interface {
	Node()
	Pos() int
}
