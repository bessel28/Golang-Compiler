package parser

import (
	. "compiler/src/components"
	. "compiler/src/components/functions"
	. "compiler/src/runtime/value"
)

type SymbolTable[ID any, T Type] interface {
	Declare(name string, varType T)
	IsDeclared(name ID) bool
	Get(name ID) T
	Child() *SymbolTable[ID, T]
}

type VarSymbolTable struct {
	Names  map[string]Type
	Parent *VarSymbolTable
}
type FuncParserSymbolTable struct {
	Names  map[string]*FunctionIdentifier
	Parent *FuncParserSymbolTable
}
type FuncSymbolTable struct {
	Names  map[FunctionIdentifier]*FunctionSignature
	Parent *FuncSymbolTable
}

func (s *VarSymbolTable) Declare(name string, varType Type) {
	s.Names[name] = varType
}

func (s *VarSymbolTable) IsDeclared(name string) bool {
	if _, exists := s.Names[name]; exists {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsDeclared(name)
	}
	return false
}
func (s *VarSymbolTable) Get(name string) Type {
	if _, exists := s.Names[name]; exists {
		typeVal := s.Names[name]
		return typeVal
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s *VarSymbolTable) Child() *VarSymbolTable {
	childTable := VarSymbolTable{}
	childTable.Parent = s
	return &childTable
}

func (s *FuncParserSymbolTable) Declare(name string, varType *FunctionIdentifier) {
	s.Names[name] = varType
}

func (s *FuncParserSymbolTable) IsDeclared(name string) bool {
	if _, exists := s.Names[name]; exists {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsDeclared(name)
	}
	return false
}
func (s *FuncParserSymbolTable) Get(name string) *FunctionIdentifier {
	if _, exists := s.Names[name]; exists {
		typeVal := s.Names[name]
		return typeVal
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s *FuncParserSymbolTable) Child() *FuncParserSymbolTable {
	childTable := FuncParserSymbolTable{}
	childTable.Parent = s
	return &childTable
}

func (s *FuncSymbolTable) Declare(name *FunctionIdentifier, varType *FunctionSignature) {
	s.Names[*name] = varType
}

func (s *FuncSymbolTable) IsDeclared(name *FunctionIdentifier) bool {
	if _, exists := s.Names[*name]; exists {
		return true
	}
	if s.Parent != nil {
		return s.Parent.IsDeclared(name)
	}
	return false
}
func (s *FuncSymbolTable) Get(name *FunctionIdentifier) *FunctionSignature {
	if _, exists := s.Names[*name]; exists {
		typeVal := s.Names[*name]
		return typeVal
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s *FuncSymbolTable) Child() *FuncSymbolTable {
	childTable := FuncSymbolTable{}
	childTable.Parent = s
	return &childTable
}
