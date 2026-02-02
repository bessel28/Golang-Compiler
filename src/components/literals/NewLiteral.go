package literals

import (
	"compiler/src/components/ast"
)

func NewLiteral(v interface{}) ast.Literal {
	switch v.(type) {
	case string:
		return &StringLiteral{v.(string)}
	case int:
		return &IntLiteral{v.(int)}
	case bool:
		return &BoolLiteral{v.(bool)}
	default:
		return nil
	}
}
