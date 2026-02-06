package functions

import (
	. "compiler/src/components"
	"compiler/src/runtime/value"
)

type FunctionExpr struct {
	Identifier *value.FunctionIdentifier
	*FunctionSignature
	Position int
}

func (f *FunctionExpr) Type() (Type, error) {
	return f.FunctionSignature.ReturnType, nil
}

func (f *FunctionExpr) Node()           {}
func (f *FunctionExpr) Pos() int        { return f.Position }
func (f *FunctionExpr) ExpressionNode() {}
func (f *FunctionExpr) Callable()       {}
func (f *FunctionExpr) SetFunctionSignature(signature *FunctionSignature) {
	f.FunctionSignature = signature
}
func (f *FunctionExpr) GetFunctionSignature() *FunctionSignature {
	return f.FunctionSignature
}
func (f *FunctionExpr) GetIdentifier() *value.FunctionIdentifier {
	return f.Identifier
}
