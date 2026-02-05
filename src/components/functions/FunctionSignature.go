package functions

import (
	. "compiler/src/components"
	"reflect"
)

type FunctionSignature struct {
	Params     []*Parameter
	ReturnType Type
}

func (sig *FunctionSignature) Type() {}

func (sig *FunctionSignature) Compare(other Type) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(sig) {
		return false
	}
	if len(sig.Params) != len(other.(*FunctionSignature).Params) {
		return false
	}
	for i, param := range sig.Params {
		if reflect.TypeOf(param) != reflect.TypeOf(other.(*FunctionSignature).Params[i]) {
			return false
		}
		switch param := param.ValType.(type) {
		case *FunctionSignature:
			if !(param.Compare(other.(*FunctionSignature).Params[i].ValType.(*FunctionSignature))) {
				return false
			}
		default:
			if !param.Compare(other.(*FunctionSignature).Params[i].ValType) {
				return false
			}
		}
	}

	if reflect.TypeOf(sig.ReturnType) != reflect.TypeOf(other.(*FunctionSignature).ReturnType) {
		return false
	}
	switch sig.ReturnType.(type) {
	case *FunctionSignature:
		if !(sig.ReturnType).(*FunctionSignature).Compare(other.(*FunctionSignature).ReturnType.(*FunctionSignature)) {
			return false
		}
	default:
		if sig.ReturnType != other.(*FunctionSignature).ReturnType {
			return false
		}
	}

	return true
}
