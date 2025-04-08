package core

import "fmt"

type Interpreter struct {
	env *Env
}

func NewInterpreter() *Interpreter {
	return &Interpreter{env: NewEnv()}
}

func (i *Interpreter) evaluate(stmts []Stmt) {
	for _, stmt := range stmts {
		i.evaluateStmt(stmt)
	}
}

func (i *Interpreter) evaluateStmt(stmt Stmt) interface{} {
	return stmt.accept(i)
}

func (i *Interpreter) evaluateExpr(expr Expr) interface{} {
	return expr.accept(i)
}

func (i *Interpreter) visitBinaryExpr(binary *Binary) interface{} {
	left := i.evaluateExpr(binary.left)
	right := i.evaluateExpr(binary.right)
	operator := binary.operator
	i.checkNumberOperand(operator, left)
	i.checkNumberOperand(operator, right)
	switch operator.TokenType {
	case STAR:
		return left.(float64) * right.(float64)
	case SLASH:
		return left.(float64) / right.(float64)
	case MINUS:
		return left.(float64) - right.(float64)
	case PLUS:
		return left.(float64) + right.(float64)
	default:
		return nil
	}
}

func (i *Interpreter) visitLiteralExpr(literal *Literal) interface{} {
	return literal.value
}

func (i *Interpreter) visitUnaryExpr(unary *Unary) interface{} {
	right := i.evaluateExpr(unary.right)
	switch unary.operator.TokenType {
	case MINUS:
		i.checkNumberOperand(unary.operator, right)
		return -(right).(float64)
	default:
		return nil
	}
}

func (i *Interpreter) visitLogicalExpr(logical *Logical) interface{} {
	left := i.evaluateExpr(logical.left)
	right := i.evaluateExpr(logical.right)
	switch logical.operator.TokenType {
	case OR:
		return left.(bool) || right.(bool)
	case AND:
		return left.(bool) && right.(bool)
	default:
		return nil
	}
}

func (i *Interpreter) visitGroupingExpr(grouping *Grouping) interface{} {
	return i.evaluateExpr(grouping.expr)
}

func (i *Interpreter) visitVarExpr(v *VarExpr) interface{} {
	return i.env.get(v.token)
}

func (i *Interpreter) visitAssignExpr(assign *Assign) interface{} {
	i.env.assign(assign.token, assign.expr)
	return nil
}

func (i *Interpreter) checkNumberOperand(operator *Token, operand interface{}) {
	if _, ok := operand.(float64); !ok {
		panic(NewParseError(operator, "Operand must be a number."))
	}
}

func (i *Interpreter) visitExpressionStmt(expression *Expression) interface{} {
	return i.evaluateExpr(expression.expr)
}

func (i *Interpreter) visitPrintStmt(print *Print) interface{} {
	value := i.evaluateExpr(print.expr)
	fmt.Printf("%v\n", value)
	return nil
}

func (i *Interpreter) visitVarStmt(v *VarStmt) interface{} {
	value := i.evaluateExpr(v.expr)
	i.env.define(v.identifier, value)
	return nil
}
