package functions

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	"errors"
	"fmt"
)

type CallExpr struct {
	Callee   Callable
	Params   []Expression
	Position int
}

func (c *CallExpr) Type() (Type, error) {
	calleeType, calleeErr := c.Callee.Type()
	if calleeErr != nil {
		return nil, calleeErr
	}
	calleeFuncSig, ok := calleeType.(*FunctionSignature)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Callee is not a callable type (%T)", calleeType))
	}
	return calleeFuncSig.ReturnType, nil
}

func (c *CallExpr) Node()           {}
func (c *CallExpr) Pos() int        { return c.Position }
func (c *CallExpr) ExpressionNode() {}
