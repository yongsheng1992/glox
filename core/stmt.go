package core

type StmtVisitor interface {
	visitExpressionStmt(expr *Expression) interface{}
	visitPrintStmt(print *Print) interface{}
	visitVarStmt(v *Var) interface{}
}

type Stmt interface {
	accept(visitor StmtVisitor) interface{}
}

type Expression struct {
	expr Expr
}

func (exprStmt *Expression) accept(visitor StmtVisitor) interface{} {
	return visitor.visitExpressionStmt(exprStmt)
}

func NewExpression(expr Expr) *Expression {
	return &Expression{
		expr: expr,
	}
}

type Print struct {
	expr Expr
}

func NewPrint(expr Expr) *Print {
	return &Print{
		expr: expr,
	}
}

func (printStmt *Print) accept(visitor StmtVisitor) interface{} {
	return visitor.visitPrintStmt(printStmt)
}

type Var struct {
	identifier *Token
	expr       Expr
}

func (v *Var) accept(visitor StmtVisitor) interface{} {
	return visitor.visitVarStmt(v)
}
