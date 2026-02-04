package components

type Type interface {
	Type()
}
type IntType struct{}
type BoolType struct{}
type StringType struct{}
type FuncType struct {
	Params []Type
	Return Type
}

func (t IntType) Type()    {}
func (t BoolType) Type()   {}
func (t StringType) Type() {}
func (t FuncType) Type()   {}
