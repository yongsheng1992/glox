# 解析expression

目前只处理表达式。

```
expression = term;
term = factor (("-" | "+")factor)*;
factor = unary (("*" | "/") unary)*;
unary = ("!" | "-") unary | primary;
primary = NUMBER | "true" | "false" | "(" expression ")";
```

* **term** 加减运算
* **factor** 乘除运算
* **unary** 单元运算
* **primary** 终结符和分组（grouping）

在定义语法的时候，需要考虑优先级和结合律。在上边的ebnf中，从上到下优先级逐渐增高。因为加减和乘除的优先级不一样，所以分别定义了**term**和**factor**，且后者在前者的下面，表示在解析的时候先解析**factor**。同理**unary**的优先级最高。

## example

### (1 + 2) * (3 + 4)

解析后token：
```
[LeftParen, Literal, Plus, Literal, RightParen, Star, LeftParen, Literal, Plus, Literal, RightParen]
```

对于第一个token **LeftParen**，只能依次匹配产生式规则，最终匹配到primary的。最终的表达式：
```go
	expr := NewBinary(
		NewGrouping(
			NewBinary(
				NewLiteral(1.0),
				&Token{TokenType: Plus, Lexeme: "+", Line: 1},
				NewLiteral(2.0),
			),
		),
		&Token{TokenType: Star, Lexeme: "*", Line: 1},
		NewGrouping(
			NewBinary(
				NewLiteral(3.0),
				&Token{TokenType: Plus, Lexeme: "+", Line: 1},
				NewLiteral(4.0),
			),
		),
	)
```

最终实现的递归下降的解析器需要按照产生式的规则生成表达式。

需要注意，在lox中数字只用number类型，可以是整型，也可以是浮点型，这里和java版本的实现类似，最终只有float64的数字。
