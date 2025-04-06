# 处理statement

前面只处理来表达式，现在开始处理语句。先处理表达式，是因为表达式简单，没有任何副作用。

所以需要扩展语法来支持语句，首先语句也分多种类型：
* 声明语句
  * 声明一个变量、方法或者是类
* 赋值语句
  * 给变量赋值
* 控制语句
  * if-else
  * for

目前只考虑处理声明、赋值和print语句。
可以定义`program`是有`statement`组成的，需要使用EOF结束：
```
program = statement* EOF;
statement = exprStmt | printStmt;
exprStmt = expression* ";";
printStmt = "print" expression ";";
```

正如前面提到的，语句还分为多种，只用表达式和print语句是不够的：
```
program = declaration* EOF;
declaration = varDecl | statment;
varDecal = "var " IDENTIFIER ("=" expression )? ";";
statement = exprStmt | printStmt;
exprStmt = expression* ";";
printStmt = "print" expression ";";
```

同时需要扩展`primary`的规则：
```
primary = "true" | "false" | "nil" | NUMBER | STRING | "(" expression ")"
        | IDENTIFIER ;
```

## Environment

语句是有副作用的，比如变量的赋值，需要将一个变量个一个值绑定。所以需要维护这种关系。

```go
type Env struct {
	values values map[string]interface{}
}

func (env *Env) put(name string, value interface{}) {
	env.values[name] = value
}
func (env *Env) get(token Token) interface{} {
	if value, exist := env.values[token.lexeme]; !exist {
	    panic(NewRuntimeError(token, fmt.Scanf("Undefined variable '%s'.", token.lexeme)))	
    }
}
```