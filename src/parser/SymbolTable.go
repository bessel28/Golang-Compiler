package parser

type SymbolTable struct {
	names  map[string]bool
	parent *SymbolTable
}

func (s *SymbolTable) Declare(name string) {
	s.names[name] = true
}

func (s *SymbolTable) IsDeclared(name string) bool {
	if s.names[name] {
		return true
	}
	if s.parent != nil {
		return s.parent.IsDeclared(name)
	}
	return false
}
