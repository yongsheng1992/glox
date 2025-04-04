package core

type Interpreter struct {
	expr Expr
}

func NewInterpreter(expr Expr) *Interpreter {
	return &Interpreter{expr: expr}
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return expr.accept(i)
}

func (i *Interpreter) visitBinaryExpr(binary *Binary) interface{} {
	left := i.evaluate(binary.left)
	right := i.evaluate(binary.right)
	operator := binary.operator
	i.checkNumberOperand(operator, left)
	i.checkNumberOperand(operator, right)
	switch operator.TokenType {
	case Star:
		return left.(float64) * right.(float64)
	case Slash:
		return left.(float64) / right.(float64)
	case Minus:
		return left.(float64) - right.(float64)
	case Plus:
		return left.(float64) + right.(float64)
	default:
		return nil
	}
}

func (i *Interpreter) visitLiteralExpr(literal *Literal) interface{} {
	return literal.value
}

func (i *Interpreter) visitUnaryExpr(unary *Unary) interface{} {
	right := i.evaluate(unary.right)
	switch unary.operator.TokenType {
	case Minus:
		i.checkNumberOperand(unary.operator, right)
		return -(right).(float64)
	default:
		return nil
	}
}

func (i *Interpreter) visitLogicalExpr(logical *Logical) interface{} {
	left := i.evaluate(logical.left)
	right := i.evaluate(logical.right)
	switch logical.operator.TokenType {
	case Or:
		return left.(bool) || right.(bool)
	case And:
		return left.(bool) && right.(bool)
	default:
		return nil
	}
}

func (i *Interpreter) visitGrouping(grouping *Grouping) interface{} {
	return i.evaluate(grouping.expr)
}

func (i *Interpreter) checkNumberOperand(operator *Token, operand interface{}) {
	if _, ok := operand.(float64); !ok {
		panic(NewParseError(operator, "Operand must be a number."))
	}
}
