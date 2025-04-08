package core

type StmtVisitor interface {
	visitExpressionStmt(expr *Expression) interface{}
	visitPrintStmt(print *Print) interface{}
	visitVarStmt(v *VarStmt) interface{}
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

func NewExprStmt(expr Expr) *Expression {
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

type VarStmt struct {
	identifier *Token
	expr       Expr
}

func (v *VarStmt) accept(visitor StmtVisitor) interface{} {
	return visitor.visitVarStmt(v)
}

func NewVarStmt(identifier *Token, expr Expr) *VarStmt {
	return &VarStmt{
		identifier: identifier,
		expr:       expr,
	}
}
