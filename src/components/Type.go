package components

import "reflect"

type Type interface {
	Type()
	Compare(Type) bool
}

type IntType struct{}
type NilType struct{}
type BoolType struct{}
type StringType struct{}

//type FuncType struct {
//	Params []Type
//	Return Type
//}

func (t *IntType) Type()    {}
func (t *BoolType) Type()   {}
func (t *StringType) Type() {}
func (t *NilType) Type()    {}

//func (t *FuncType) Type()   {}

func (i *IntType) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(i) {
		return false
	}
	return true
}
func (b *BoolType) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(b) {
		return false
	}
	return true
}
func (s *StringType) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(s) {
		return false
	}
	return true
}
func (s *NilType) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(s) {
		return false
	}
	return true
}
