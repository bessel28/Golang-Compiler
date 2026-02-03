package literals

func NewLiteral(v interface{}, pos *int) Literal {
	if pos == nil {
		tempPos := -1
		pos = &tempPos
	}
	switch v.(type) {
	case string:
		return &StringLiteral{v.(string), *pos}
	case int:
		return &IntLiteral{v.(int), *pos}
	case bool:
		return &BoolLiteral{v.(bool), *pos}
	default:
		return nil
	}
}
