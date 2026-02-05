package parser

import (
	"compiler/src/components"
	"compiler/src/components/functions"
)

type SymbolTable[T any] struct {
	Names  map[string]T
	Parent *SymbolTable[T]
}

func (s *SymbolTable[T]) GetType(name string) *T {
	if s.IsDeclared(name) {
		typeVal := s.Names[name]
		return &typeVal
	}
	return nil
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
func (s *SymbolTable[T]) Get(name string) *T {
	if _, exists := s.Names[name]; exists {
		typeVal := s.Names[name]
		return &typeVal
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s *SymbolTable[T]) Child() *SymbolTable[T] {
	childTable := SymbolTable[T]{}
	childTable.Parent = s
	return &childTable
}

type VarSymbolTable = SymbolTable[components.Type]
type FuncSymbolTable = SymbolTable[*functions.FunctionSignature]
