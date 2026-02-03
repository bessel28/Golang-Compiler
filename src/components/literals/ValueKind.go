package literals

type ValueKind int

const (
	IntKind ValueKind = iota
	StringKind
	BoolKind
	FunctionRefKind
)
