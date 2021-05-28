package ast

import (
	"Q/object"
	"Q/token"
)

// ExpressionStmt : implement Statement
type ExpressionStmt struct {
	Tok  *token.Token
	Expr Expression
}

func (this *ExpressionStmt) statementNode() {}
func (this *ExpressionStmt) TokenLiteral() string {
	return this.Tok.Literal
}
func (this *ExpressionStmt) String() string {
	if this.Expr != nil {
		return this.Expr.String()
	}
	return ""
}
func (this *ExpressionStmt) Eval(env *object.Env, insideLoop bool) (object.Object, error) {
	return this.Expr.Eval(env, insideLoop)
}
