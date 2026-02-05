package parsing

import (
	. "compiler/src/components"
	. "compiler/src/components/ast"
	. "compiler/src/components/literals"
	. "compiler/src/components/operators"
	"compiler/src/components/statements"
	. "compiler/src/parser"
	. "compiler/src/parser/expression"
	"fmt"
	"testing"
)

func Test_TypeChcker(t *testing.T) {
	var tests = []struct {
		input      string
		expression Expression
	}{
		{
			"1+1",
			&BinaryExpr{AddOperator{}, NewLiteral(1, nil), NewLiteral(1, nil), &IntType{}},
		},
	}

	for index, tt := range tests {
		testName := fmt.Sprintf("TypeChecker%d", index)
		t.Run(testName, func(t *testing.T) {
			parser, parserErr := NewParser(tt.input)
			if parserErr != nil {
				t.Fatal(parserErr)
			}

			expr, exprErr := parser.ParseExpr()
			if exprErr != nil {
				t.Fatal(exprErr)
			}

			stmts := []Statement{&statements.ExpressionStatement{expr, 0}}

			typeChecker := NewTypeChecker(&stmts)

			checkErr := typeChecker.CheckTypes()
			if checkErr != nil {
				t.Fatal(checkErr)
			}

			exprType, _ := expr.Type()
			fmt.Printf("%T", exprType)
		})
	}
}
