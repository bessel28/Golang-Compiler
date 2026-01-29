# Custom GoLang Compiler

## Structure

Below I'll go into detail of the components and behaviour of the compiler

### Components

#### Literals
These are the base value types that are supported by the compiler
- Int
- String
- Bool

#### Conditions

- EqualTo
- GreaterThan
- LessThan
- NotEqualTo

#### Operators
- Add
- Subtract
- Multiply
- Divide
- Ternary

## AST Tree Structure

The heirarchy of the AST determines the precedence of tokens

- ParseExpr - Any Expression
- ParseTernary - Ternary Operators
- ParseOrExpr - Binary Or operator 
- ParseAndExpr - Binary And operator 
- ParseEqualityExpr - == and !=
- RelationalExpr - > and <
- ParseAdditive - Plus or Minus Binary Operators
- ParseMultiplicative - Times or Divide Binary Operators
- ParseUnary - Unary Operators
- ParseTerm - Literals, Parentheses
```
Expr -> Ternary
Ternary -> OrExpr ( ? Expr : Ternary )?
OrExpr -> AndExpr ( || AndExpr )*
AndExpr -> EqualityExpr ( && EqualityExpr )*
EqualityExpr ->  RelationalExpr ( ( == | != ) RelationalExpr )
RelationalExpr ->  AdditiveExpr  ( ( > | < ) AdditiveExpr )
Additive -> Multiplicative (( + | - ) Multiplicative )*
Multiplicative -> Unary (( * | / ) Unary )*
Unary -> (!)? Term
Term -> Literal | ( Expr )
```
