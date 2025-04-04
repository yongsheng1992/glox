package core

type Expr interface {
	accept(visitor Visitor) interface{}
}

type Visitor interface {
	visitBinaryExpr(binary *Binary) interface{}
	visitLiteralExpr(literal *Literal) interface{}
	visitUnaryExpr(unary *Unary) interface{}
	visitLogicalExpr(logical *Logical) interface{}
	visitGrouping(grouping *Grouping) interface{}
}

type Binary struct {
	left     Expr
	operator *Token
	right    Expr
}

func (binary *Binary) accept(visitor Visitor) interface{} {
	return visitor.visitBinaryExpr(binary)
}

func NewBinary(left Expr, operator *Token, right Expr) *Binary {
	return &Binary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

type Unary struct {
	operator *Token
	right    Expr
}

func (unary *Unary) accept(visitor Visitor) interface{} {
	return visitor.visitUnaryExpr(unary)
}

func NewUnary(operator *Token, right Expr) *Unary {
	return &Unary{
		operator: operator,
		right:    right,
	}
}

type Literal struct {
	value interface{}
}

func (literal *Literal) accept(visitor Visitor) interface{} {
	return visitor.visitLiteralExpr(literal)
}

func NewLiteral(value interface{}) *Literal {
	return &Literal{
		value: value,
	}
}

type Logical struct {
	left     Expr
	operator *Token
	right    Expr
}

func (logical *Logical) accept(visitor Visitor) interface{} {
	return visitor.visitLogicalExpr(logical)
}

func NewLogical(left Expr, operator *Token, right Expr) *Logical {
	return &Logical{
		left:     left,
		operator: operator,
		right:    right,
	}
}

type Grouping struct {
	expr Expr
}

func (grouping *Grouping) accept(visitor Visitor) interface{} {
	return visitor.visitGrouping(grouping)
}
func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		expr: expr,
	}
}
