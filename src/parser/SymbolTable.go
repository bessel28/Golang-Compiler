package parser

import (
	"compiler/src/components/functions"
	. "compiler/src/components/literals"
)

type SymbolTable[T any] struct {
	Names  map[string]*T
	Parent *SymbolTable[T]
}

func (s *SymbolTable[T]) GetType(name string) *T {
	if s.IsDeclared(name) {
		return s.Names[name]
	}
	return nil
}

func (s *SymbolTable[T]) Declare(name string, varType T) {
	s.Names[name] = &varType
}

func (s *SymbolTable[T]) IsDeclared(name string) bool {
	if _, exists := s.Names[name]; exists {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsDeclared(name)
	}
	return false
}
func (s *SymbolTable[T]) Get(name string) *T {
	if _, exists := s.Names[name]; exists {
		return s.Names[name]
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

type VarSymbolTable = SymbolTable[ValueKind]
type FuncSymbolTable = SymbolTable[functions.FunctionSignature]
