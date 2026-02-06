package value

import "fmt"

var id = 0

type FunctionIdentifier struct {
	Name string
	UUID int
}

func NewFunctionIdentifier(name string) *FunctionIdentifier {
	uuid := id
	id++
	return &FunctionIdentifier{name, uuid}
}

func (id FunctionIdentifier) String() string {
	return fmt.Sprintf("%v#%v", id.Name, id.UUID)
}
