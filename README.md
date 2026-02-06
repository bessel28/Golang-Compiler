# Custom GoLang Compiler

A custom GoLang Compiler I wrote as an introduction to GoLang

## Components

### Literals
These are the base value types that are supported by the compiler
- Int
- String
- Bool
 
### Identifiers
- Functions
- Variables

### Conditions

- EqualTo
- GreaterThan
- LessThan
- NotEqualTo

### Operators
- Add
- Subtract
- Multiply
- Divide
- Ternary

### Statements
- Blocks
- Lone Expression
- Function Definition
- Variable Definition
- Variable Assignment

### Functions
- Function Calls

## AST Statement Tree Structure



## AST Expression Tree Structure

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
- ParseCall - Function calls
- ParseTerm - Literals, Parentheses, Identifiers
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
