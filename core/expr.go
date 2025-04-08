package core

type Expr interface {
	accept(visitor ExprVisitor) interface{}
}

type ExprVisitor interface {
	visitBinaryExpr(binary *Binary) interface{}
	visitLiteralExpr(literal *Literal) interface{}
	visitUnaryExpr(unary *Unary) interface{}
	visitLogicalExpr(logical *Logical) interface{}
	visitGroupingExpr(grouping *Grouping) interface{}
	visitVarExpr(v *VarExpr) interface{}
	visitAssignExpr(assign *Assign) interface{}
}

type Binary struct {
	left     Expr
	operator *Token
	right    Expr
}

func (binary *Binary) accept(visitor ExprVisitor) interface{} {
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

func (unary *Unary) accept(visitor ExprVisitor) interface{} {
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

func (literal *Literal) accept(visitor ExprVisitor) interface{} {
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

func (logical *Logical) accept(visitor ExprVisitor) interface{} {
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

func (grouping *Grouping) accept(visitor ExprVisitor) interface{} {
	return visitor.visitGroupingExpr(grouping)
}
func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		expr: expr,
	}
}

type VarExpr struct {
	token *Token
}

func NewVarExpr(token *Token) *VarExpr {
	return &VarExpr{
		token: token,
	}
}

func (v *VarExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitVarExpr(v)
}

type Assign struct {
	token *Token
	expr  Expr
}

func NewAssign(token *Token, expr Expr) *Assign {
	return &Assign{
		token: token,
		expr:  expr,
	}
}

func (assign *Assign) accept(visitor ExprVisitor) interface{} {
	return visitor.visitAssignExpr(assign)
}
