package value

import (
	. "compiler/src/components/literals"
	"fmt"
)

type RuntimeValue interface {
	Value() any
	Kind() ValueKind
	String() *StringValue
	BoolVal() *BoolValue
}

func NewValue(literal Literal) RuntimeValue {
	if literal == nil {
		return &NilValue{}
	}
	switch literal.(type) {
	case *IntLiteral:
		return &IntValue{literal.(*IntLiteral).Value}
	case *StringLiteral:
		return &StringValue{literal.(*StringLiteral).Value}
	case *BoolLiteral:
		return &BoolValue{literal.(*BoolLiteral).Value}
	}
	return &NilValue{}
}

type IntValue struct {
	Val int
}
type StringValue struct {
	Val string
}
type BoolValue struct {
	Val bool
}
type FuncValue struct {
	Val *FunctionIdentifier
}
type NilValue struct{}

func (v *IntValue) Value() any          { return v.Val }
func (v *IntValue) Kind() ValueKind     { return IntKind }
func (v *IntValue) BoolVal() *BoolValue { return &BoolValue{v.Val != 0} }

func (v *IntValue) String() *StringValue { return &StringValue{fmt.Sprintf("%d", v.Val)} }

func (v *StringValue) Value() any           { return v.Val }
func (v *StringValue) Kind() ValueKind      { return StringKind }
func (v *StringValue) BoolVal() *BoolValue  { return &BoolValue{v.Val != ""} }
func (v *StringValue) String() *StringValue { return v }

func (v *BoolValue) Value() any           { return v.Val }
func (v *BoolValue) Kind() ValueKind      { return BoolKind }
func (v *BoolValue) BoolVal() *BoolValue  { return v }
func (v *BoolValue) String() *StringValue { return &StringValue{fmt.Sprintf("%t", v.Val)} }

func (v *FuncValue) Value() any           { return v.Val }
func (v *FuncValue) Kind() ValueKind      { return FunctionRefKind }
func (v *FuncValue) BoolVal() *BoolValue  { return &BoolValue{true} }
func (v *FuncValue) String() *StringValue { return &StringValue{"<Function>"} }

func (v *NilValue) Value() any           { return nil }
func (v *NilValue) Kind() ValueKind      { return NilKind }
func (v *NilValue) BoolVal() *BoolValue  { return &BoolValue{false} }
func (v *NilValue) String() *StringValue { return &StringValue{"<nil>"} }
