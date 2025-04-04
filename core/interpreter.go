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
	right := i.evaluate(unary.right)
	switch unary.operator.TokenType {
	case MINUS:
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
	case OR:
		return left.(bool) || right.(bool)
	case AND:
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
