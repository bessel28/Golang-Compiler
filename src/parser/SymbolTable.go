package parser

import (
	. "compiler/src/components"
	. "compiler/src/components/functions"
)

type SymbolTable[T Type] struct {
	Names  map[string]T
	Parent *SymbolTable[T]
}

func (s *SymbolTable[T]) Declare(name string, varType T) {
	s.Names[name] = varType
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
func (s *SymbolTable[T]) Get(name string) T {
	if _, exists := s.Names[name]; exists {
		typeVal := s.Names[name]
		return typeVal
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	var zero T
	return zero
}

func (s *SymbolTable[T]) Child() *SymbolTable[T] {
	childTable := SymbolTable[T]{}
	childTable.Parent = s
	return &childTable
}

type VarSymbolTable = SymbolTable[Type]
type FuncSymbolTable = SymbolTable[*FunctionSignature]
