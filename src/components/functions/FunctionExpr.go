package functions

import . "compiler/src/components"

type FunctionExpr struct {
	FunctionSignature
	Position int
}

func (f *FunctionExpr) Type() Type {
	var params = make([]Type, len(f.Params))
	for _, paramType := range f.Params {
		params = append(params, paramType.ValType)
	}
	return &FuncType{params, f.ReturnType}
}

func (f *FunctionExpr) Node()           {}
func (f *FunctionExpr) Pos() int        { return f.Position }
func (f *FunctionExpr) ExpressionNode() {}
