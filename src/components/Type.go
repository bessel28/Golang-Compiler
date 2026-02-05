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
type FuncType struct {
	Params []Type
	Return Type
}

func (t *IntType) Type()    {}
func (t *BoolType) Type()   {}
func (t *StringType) Type() {}
func (t *NilType) Type()    {}
func (t *FuncType) Type()   {}

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
func (f *FuncType) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(f) {
		return false
	}
	if len(f.Params) != len(other.(*FuncType).Params) {
		return false
	}
	for i, param := range f.Params {
		if reflect.TypeOf(param) != reflect.TypeOf(other.(*FuncType).Params[i]) {
			return false
		}
		switch param := param.(type) {
		case *FuncType:
			if !(param.Compare(other.(*FuncType).Params[i].(*FuncType))) {
				return false
			}
		default:
			if param != other.(*FuncType).Params[i] {
				return false
			}
		}
	}

	if reflect.TypeOf(f.Return) != reflect.TypeOf(other.(*FuncType).Return) {
		return false
	}
	switch f.Return.(type) {
	case *FuncType:
		if !(f.Return).(*FuncType).Compare(other.(*FuncType).Return.(*FuncType)) {
			return false
		}
	default:
		if f.Return != other.(*FuncType).Return {
			return false
		}
	}

	return true
}
