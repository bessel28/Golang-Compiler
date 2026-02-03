package literals

func NewLiteral(v interface{}) Literal {
	switch v.(type) {
	case string:
		return &StringLiteral{v.(string), -1}
	case int:
		return &IntLiteral{v.(int), -1}
	case bool:
		return &BoolLiteral{v.(bool), -1}
	default:
		return nil
	}
}
