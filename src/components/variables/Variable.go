package variables

import (
	. "compiler/src/components"
	. "compiler/src/components/functions"
)

type VariableExpr struct {
	Identifier string
	ReturnType Type
	Position   int
}

func (v *VariableExpr) Node()               {}
func (v *VariableExpr) Pos() int            { return v.Position }
func (v *VariableExpr) ExpressionNode()     {}
func (v *VariableExpr) Type() (Type, error) { return v.ReturnType, nil }
func (v *VariableExpr) Callable()           {}
func (v *VariableExpr) SetFunctionSignature(signature *FunctionSignature) {
	v.ReturnType = signature
}
func (v *VariableExpr) GetFunctionSignature() *FunctionSignature {
	return v.ReturnType.(*FunctionSignature)
}
func (v *VariableExpr) GetIdentifier() string {
	return v.Identifier
}
