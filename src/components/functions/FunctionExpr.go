package functions

type FunctionExpr struct {
	FunctionSignature
	Position int
}

func (f *FunctionExpr) Node()           {}
func (f *FunctionExpr) Pos() int        { return f.Position }
func (f *FunctionExpr) ExpressionNode() {}
